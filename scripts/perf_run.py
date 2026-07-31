import sys
import os
import subprocess
import shutil

# Make sure stdout supports UTF-8
if hasattr(sys.stdout, 'reconfigure'):
    sys.stdout.reconfigure(encoding='utf-8')

# Configured paths
COMPONENTS = {
    "pranor-auth": "PranorAuth",
    "servdb": "ServDB",
    "pranor-mesh": "PranorMesh",
    "pranor-gate": "PranorGate/pkg/proxy",
    "pranor-pulse": "PranorPulse/pkg/broker",
    "pranor-vault": "PranorVault/pkg/storage",
    "pranor-cache": "PranorCache",
    "servshared": "PranorCore",
    "pranor-trace": "PranorTrace",
    "pranor-chrono": "PranorChrono",
    "pranor-flow": "PranorFlow",
    "pranor-hub": "PranorHub"
}

ROOT_DIR = "c:\\Mine\\try\\serv"
BASELINES_DIR = os.path.join(ROOT_DIR, "pranor-repo", "tests", "perf", "baselines")

def ensure_dirs():
    os.makedirs(BASELINES_DIR, exist_ok=True)

def run_benchmarks(name, path, out_file):
    abs_path = os.path.join(ROOT_DIR, path)
    print(f"Running benchmarks for {name} in {abs_path}...")
    try:
        cmd = ["go", "test", "-bench=Benchmark", "-run=^$", "-benchmem", "."]
        res = subprocess.run(cmd, cwd=abs_path, capture_output=True, text=True, check=True)
        with open(out_file, "w", encoding="utf-8") as f:
            f.write(res.stdout)
        print(f"  Saved results to {out_file}")
    except subprocess.CalledProcessError as e:
        print(f"  Error running benchmarks for {name}: {e.stderr}")

def save_baseline():
    ensure_dirs()
    print("=== CAPTURING PERFORMANCE BASELINES ===")
    for name, path in COMPONENTS.items():
        out_file = os.path.join(BASELINES_DIR, f"baseline_{name}.txt")
        run_benchmarks(name, path, out_file)
    print("Baselines captured successfully.")

def compare():
    ensure_dirs()
    print("=== RUNNING PERFORMANCE COMPARISONS ===")
    temp_dir = os.path.join(ROOT_DIR, "pranor-repo", "tests", "perf", "temp_run")
    os.makedirs(temp_dir, exist_ok=True)
    
    for name, path in COMPONENTS.items():
        baseline_file = os.path.join(BASELINES_DIR, f"baseline_{name}.txt")
        if not os.path.exists(baseline_file):
            print(f"No baseline file found for {name}. Capture baseline first.")
            continue
            
        current_file = os.path.join(temp_dir, f"new_{name}.txt")
        run_benchmarks(name, path, current_file)
        
        # Compare using benchstat
        print(f"\n--- Comparison for {name} ---")
        try:
            cmd = ["benchstat", baseline_file, current_file]
            res = subprocess.run(cmd, capture_output=True, text=True, check=True)
            print(res.stdout)
        except subprocess.CalledProcessError as e:
            # Fallback if benchstat isn't in PATH but is in GOBIN
            gobin_path = os.path.join(os.environ.get("USERPROFILE", ""), "go", "bin", "benchstat")
            try:
                cmd = [gobin_path, baseline_file, current_file]
                res = subprocess.run(cmd, capture_output=True, text=True, check=True)
                print(res.stdout)
            except Exception as ex:
                print(f"  Error executing benchstat comparison: {e.stderr or ex}")

    shutil.rmtree(temp_dir, ignore_errors=True)

def main():
    if len(sys.argv) < 2:
        print("Usage: python perf_run.py [save-baseline|compare]")
        sys.exit(1)
        
    action = sys.argv[1]
    if action == "save-baseline":
        save_baseline()
    elif action == "compare":
        compare()
    else:
        print(f"Unknown action: {action}")
        sys.exit(1)

if __name__ == "__main__":
    main()
