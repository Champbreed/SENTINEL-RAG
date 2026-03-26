# Sentinel-RAG 🛡️💡
**A Grounded Knowledge Assistant for Auditable Kernel Security & Neurodiverse Accessibility.**

Sentinel-RAG is a specialized accessibility tool designed to reduce cognitive load for individuals—including those with ADHD, autism, and dyslexia—interacting with complex Linux Kernel telemetry. Built for the **Microsoft Neurodiversity Challenge**, it leverages Python-driven Machine Learning to transform dense, high-anxiety technical data into structured, calm, and actionable insights.

## 🧠 Features for Neurodiversity
- **Cognitive Load Reduction:** Uses Python logic to parse raw eBPF logs into a structured, distraction-free TUI (Terminal User Interface).
- **MLT-Driven Cognitive Compression:** Our Python backend (`rai_audit.py`) uses Machine Learning Text (MLT) logic to dynamically adjust information density.
    - **Level 0 (Technical):** Full breakdown of 7+ security targets (e.g., `capabilities.7`, `bpf.2`).
    - **Level 2 (Summarized):** Intelligent compression into a single "SYSTEM SUMMARY" to prevent information overload.
- **Grounded Remediation Hints:** Provides non-anxiety-inducing, step-by-step tasks to fix non-compliance issues.
- **Visual Focus Zones:** Dedicated spatial areas for Vision, Kernel Monitoring, and AI Coaching to prevent sensory overwhelm.

## 🛠️ Technical Stack & Focus
- **Python (Core Intelligence):** The heart of the system. Uses Python for the Responsible AI (RAI) backend to analyze kernel data and generate grounded responses.
- **Microsoft RAI Toolbox:** Integrated to ensure all AI outputs are safe, ethical, and supportive.
- **Go (Golang):** Provides the high-performance, accessible interface shell using Bubble Tea & Lipgloss.
- **RAG Integration:** Security hints are strictly grounded in authoritative kernel "Rulebooks" (`capabilities.7`, `bpf.2`) to eliminate AI hallucinations.

## 🚀 Getting Started
1. **Activate Python Environment:** `source venv/bin/activate`
2. **Install ML & RAI Deps:** `pip install pandas responsibleai`
3. **Test Compression Levels:** - `python3 scripts/rai_audit.py --level 0` (Detailed)
   - `python3 scripts/rai_audit.py --level 2` (Summarized)
4. **Launch the TUI:** `go run main.go`
