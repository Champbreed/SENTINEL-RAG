import pandas as pd
import json
import argparse
from responsibleai import RAIInsights

def run_grounded_audit(level):
    data = {
        "Target": ["capabilities.7", "bpf.2", "vfs_verif", "netlink_raw", "rootless_cfg", "seccomp", "sys_admin"],
        "Groundedness": [0.85, 0.92, 0.78, 0.88, 0.95, 0.81, 0.90]
    }

    df = pd.DataFrame(data)

    if level == 1:
        df["Groundedness"] = df["Groundedness"].round(1)
    elif level == 2:
        df = df.sort_values(by="Groundedness").head(1)
        df["Target"] = "SYSTEM SUMMARY"

    results = []
    for _, row in df.iterrows():
        results.append({
            "target": row["Target"],
            "score": float(row["Groundedness"]),
            "compliant": bool(row["Groundedness"] > 0.80)
        })

    print(json.dumps(results))

if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("--level", type=int, default=0)
    args = parser.parse_args()
    
    run_grounded_audit(args.level)
