package main

import (
        "encoding/json"
	"fmt"
	"os"
        "os/exec"
        "strconv"
	"time"
        
        "github.com/yourusername/sentinel-rag/internal/audit"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	zoneStyle    = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(1, 2)
	visionStyle  = zoneStyle.Copy().BorderForeground(lipgloss.Color("42")).Width(30).Height(15)
	monitorStyle = zoneStyle.Copy().BorderForeground(lipgloss.Color("33")).Width(60).Height(15)
	coachStyle   = zoneStyle.Copy().BorderForeground(lipgloss.Color("99")).Width(94).Height(10)
)

type tickMsg time.Time

type model struct {
	frameIndex int
	kernelLog  string
	coachMsg   string
	counter    int
	paused     bool
       
        results    []audit.AuditResult
        err        error
        loading    bool
        readingLevel int
}

type auditMsg []audit.AuditResult
type errMsg struct{ err error }

var handFrames = []string{
	`    ⠀⠁⠂⠃⠄⠅⠆⠇
    ⠈⠉⠊⠋⠌⠍⠎⠏`,
	`    ⠐⠑⠒⠓⠔⠕⠖⠗
    ⠘⠙⠚⠛⠜⠝⠞⠟`,
	`    ⠠⠡⠢⠣⠤⠥⠦⠧
    ⠨⠩⠪⠫⠬⠭⠮⠯`,
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		m.runAudit(),
		tea.Tick(time.Millisecond*800, func(t time.Time) tea.Msg {
			return tickMsg(t)
		}),
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case " ":
			m.paused = !m.paused
			return m, nil
		case "r":
			m.loading = true
			m.kernelLog = "[SYSTEM] RE-TRIGGERING RAI AUDIT..."
			return m, m.runAudit()
		case "s":
			content := m.View()
			_ = os.WriteFile("vision_snapshot.txt", []byte(content), 0644)
                        successMsg := lipgloss.NewStyle().Foreground(lipgloss.Color("42")).Render("✔ SNAPSHOT SAVED")
			m.kernelLog = fmt.Sprintf("[SYSTEM] %s TO vision_snapshot.txt", successMsg)
			return m, nil
               case "l":
	                m.readingLevel = (m.readingLevel + 1) % 3
	                levels := []string{"NORMAL", "SIMPLE", "ACCESSIBLE"}
	                m.kernelLog = fmt.Sprintf("[SYSTEM] Accessibility Mode: %s", levels[m.readingLevel])
	                m.loading = true
	                return m, m.runAudit()
		}

	case auditMsg:
		m.results = msg
		m.loading = false
		m.kernelLog = "[OK] RAI AUDIT COMPLETE"
		return m, nil

	case errMsg:
		m.err = msg.err
		m.loading = false
		m.kernelLog = fmt.Sprintf("[ERROR] RAI FAILED: %v", msg.err)
		return m, nil

	case tickMsg:
		if m.paused {
			return m, tea.Tick(time.Millisecond*800, func(t time.Time) tea.Msg {
				return tickMsg(t)
			})
		}

		m.frameIndex = (m.frameIndex + 1) % len(handFrames)
		m.counter++

		m.kernelLog = fmt.Sprintf("[OK] Verifier active...\n[TRACE] Analyzing capability #%d\n[STATUS] PAUSED: %v", m.counter, m.paused)

		if m.counter%4 == 0 || m.coachMsg == "" {
			fixmes := []string{
				"FIXME: User Namespace Rootless Permission Check (capabilities.7)",
				"FIXME: Inconsistent CAP_SYS_ADMIN overrides (MCC Section 4)",
				"FIXME: Kernel eBPF Helper Verification (bpf.2 Giant)",
				"FIXME: VFS Capability mapping in unprivileged containers",
				"FIXME: Netlink raw socket permission bypass checks",
				"FIXME: Securebits inheritance in non-initial user namespaces",
				"FIXME: BPF_PROG_TYPE_KPROBE accessibility audit",
			}

			if len(m.results) > 0 {
				res := m.results[(m.counter/4)%len(m.results)]
				status := "PASS"
				if !res.Compliant {
					status = "FAIL"
				}
				m.coachMsg = fmt.Sprintf("AUDIT: %s\nRAI STATUS: %s (Score: %.2f)", res.Target, status, res.Score)
			} else {
				m.coachMsg = fmt.Sprintf("AUDIT: %s\nRAI STATUS: Grounded via Responsible AI Toolbox", fixmes[(m.counter/4)%len(fixmes)])
			}
		}

		return m, tea.Tick(time.Millisecond*800, func(t time.Time) tea.Msg {
			return tickMsg(t)
		})
	}
	return m, nil
}

func (m model) View() string {
	if m.err != nil {
		return coachStyle.Render(fmt.Sprintf("❌ RAI BACKEND ERROR: %v\nTry 'source venv/bin/activate'", m.err))
	}

	var coachContent string
	if m.loading {
		coachContent = "⏳ [Sentinel-RAG] Running Responsible AI Audit... Please wait."
	} else if len(m.results) > 0 {
		coachContent = fmt.Sprintf("📊 Simplicity Level: %d/3\n\n", m.readingLevel)
		hasFailure := false 
		
		for _, res := range m.results {
			status := "✅"
			if !res.Compliant {
				status = "⚠️ "
				hasFailure = true
			}
			coachContent += fmt.Sprintf("%s %-15s Score: %.2f\n", status, res.Target, res.Score)
		}

		if hasFailure {
			hintStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("208")).Bold(true)
			coachContent += "\n" + hintStyle.Render("💡 REMEDIATION HINT:")
			coachContent += "\nLow groundedness detected in vfs_verif. Review the kernel headers\n" +
				"for LSM hooks and ensure Seccomp profiles are not shadowing permissions."
		} else {
			successStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("42"))
			coachContent += "\n" + successStyle.Render("✨ System is fully compliant.")
		}
	} else {
		coachContent = m.coachMsg
	}

	topRow := lipgloss.JoinHorizontal(lipgloss.Top,
		visionStyle.Render(fmt.Sprintf("👁️ VISION ZONE\n\n%s", handFrames[m.frameIndex])),
		monitorStyle.Render(fmt.Sprintf("🛡️ KERNEL MONITOR\n\n%s", m.kernelLog)),
	)

	bottomRow := coachStyle.Render(fmt.Sprintf("💡 GROUNDED COACH (RAI Toolbox)\n\n%s", coachContent))

	footerStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("241")).MarginLeft(2)
	footer := footerStyle.Render("\n [q] Quit | [r] Audit | [s] Snapshot | [l] Toggle Level")

	return lipgloss.JoinVertical(lipgloss.Left, topRow, bottomRow, footer)
}

func main() {
	m := model{
		frameIndex: 0,
		kernelLog:  "Initializing eBPF Monitor...",
		coachMsg:   "Knowledge Assistant Ready. (Using Responsible AI Guardrails)",
		paused:     false,
		loading:    true, // Start in loading state for the RAI audit
	}

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Visual Regression Error (#11462): %v", err)
		os.Exit(1)
	}
}

func (m model) runAudit() tea.Cmd {
	return func() tea.Msg {
		res, err := FetchRAIData(m.readingLevel)
		if err != nil {
			return errMsg{err}
		}
		return auditMsg(res)
	}
}

func FetchRAIData(level int) ([]audit.AuditResult, error) {
	cmd := exec.Command("python3", "scripts/rai_audit.py", "--level", strconv.Itoa(level))
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var results []audit.AuditResult
	if err := json.Unmarshal(out, &results); err != nil {
		return nil, err
	}
	return results, nil
}
