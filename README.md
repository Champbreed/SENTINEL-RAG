# Sentinel-RAG 🛡️💡
**A Grounded Knowledge Assistant for Auditable Kernel Security & Neurodiverse Accessibility.**

![Sentinel-RAG Demo](tui_final.jpg)

Sentinel-RAG is a specialized accessibility tool designed to reduce cognitive load for individuals—including those with ADHD, autism, and dyslexia—interacting with complex Linux Kernel telemetry. Built for the **Microsoft Neurodiversity Challenge**, it leverages Python-driven Machine Learning to transform dense, high-anxiety technical data into structured, calm, and actionable insights.

## 🧠 Features for Neurodiversity
- **Cognitive Load Reduction:** Uses Python logic to parse raw eBPF logs into a structured, distraction-free TUI (Terminal User Interface).
- **MLT-Driven Cognitive Compression:** Our Python backend (`rai_audit.py`) uses Machine Learning Text (MLT) logic to dynamically adjust information density. This prevents **"Analysis Paralysis"** by allowing users to choose the depth of information that matches their current cognitive energy.
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

Follow these steps to set up the **Sentinel-RAG** environment on your local machine.

### 📋 Prerequisites
* **Python 3.10+**
* **Go 1.21+**
* **Linux/WSL2 Environment** (Required for Kernel telemetry access)

### ⚙️ Installation & Setup

1. **Clone the Repository**
   ```bash
   git clone [https://github.com/champbreed/sentinel-rag.git](https://github.com/champbreed/sentinel-rag.git)
   cd sentinel-rag
Configure Python Environment
Initialize the virtual environment to isolate Machine Learning dependencies.

Bash
python3 -m venv venv
source venv/bin/activate
pip install --upgrade pip
pip install pandas responsibleai
Verify MLT Compression Engine
Ensure the Python backend is correctly simplifying technical data.

Bash
# View detailed technical audit (Level 0)
python3 scripts/rai_audit.py --level 0

# View AI-compressed system summary (Level 2)
python3 scripts/rai_audit.py --level 2
Launch the Accessibility TUI
Start the interactive Go-based interface.

Bash
go run main.go
🛡️ Safety & Ethics
Sentinel-RAG adheres to Microsoft's Responsible AI principles. It employs a Strict-Groundedness policy, ensuring the AI Coach only suggests security remediations verified against official Linux Kernel documentation, preventing harmful or misleading "hallucinations" in critical security contexts.
