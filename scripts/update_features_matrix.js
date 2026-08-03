const fs = require("fs");
const path = require("path");

const htmlPath = "/home/developer/workspace/pranor-platform/features.html";
const html = fs.readFileSync(htmlPath, "utf8");

// Parse existing featuresData
const match = html.match(/const featuresData = (\[[\s\S]*?\]);/);
let existingFeatures = eval(match[1]);
console.log("Base features:", existingFeatures.length);

const existingNames = new Set(existingFeatures.map(f => f.name.toLowerCase()));

// Read roadmap files
const dir = "/home/developer/workspace/pranor-platform";
const files = fs.readdirSync(dir).filter(f => f.startsWith("UNIFIED_ROADMAP")).sort();

let newFeatures = [];

files.forEach(f => {
  const content = fs.readFileSync(path.join(dir, f), "utf8");
  const lines = content.split("\n");
  lines.forEach(l => {
    if (l.startsWith("|") && !l.includes("---") && !l.includes("| # |") && !l.includes("| Item |") && !l.includes("| DC.")) {
      const parts = l.split("|").map(p => p.trim());
      if (parts.length >= 5) {
        let rawName = parts[2].replace(/\*\*/g, "").trim();
        let rawScope = parts[3].replace(/\*\*/g, "").trim();
        let rawDesc = parts[4] ? parts[4].replace(/\*\*/g, "").trim() : "";
        let rawTier = parts[5] ? parts[5].replace(/\*\*/g, "").trim() : "";

        if (rawName.includes("Total Items") || rawName.includes("funcs") || rawName.match(/^\d+$/) || rawName.length < 3) return;

        let name = rawName;
        let desc = rawDesc || rawScope;
        if (rawScope && rawDesc && !rawScope.includes("[x]") && !rawScope.includes("[ ]")) {
          desc = rawScope + ": " + rawDesc;
        }

        let edition = "OSS";
        const fullLine = l.toUpperCase();
        if (fullLine.includes("| EE |") || fullLine.includes("| **EE** |") || name.startsWith("EE.") || rawTier.toUpperCase().includes("EE") || fullLine.includes("ENTERPRISE")) {
          edition = "EE";
        }

        name = name.replace(/^EE\.\d+\.\d+\s*/, "").replace(/^LSP\.\d+\s*/, "").replace(/^VS\.[A-Z0-9]+\s*/, "").trim();
        name = name.replace(/[\r\n\t]/g, " ").replace(/\[x\]|\[ \]|✅|❌/g, "").replace(/"/g, "'").trim();
        desc = desc.replace(/[\r\n\t]/g, " ").replace(/\[x\]|\[ \]|✅|❌/g, "").replace(/"/g, "'").trim();

        if (name && desc && !existingNames.has(name.toLowerCase())) {
          existingNames.add(name.toLowerCase());

          let moduleName = "Pranor Platform";
          const searchStr = (name + " " + desc + " " + rawScope).toLowerCase();
          if (searchStr.includes("gate") || searchStr.includes("proxy") || searchStr.includes("ingress") || searchStr.includes("ebpf") || searchStr.includes("waf")) moduleName = "Pranor Gate";
          else if (searchStr.includes("vault") || searchStr.includes("s3") || searchStr.includes("vector") || searchStr.includes("storage") || searchStr.includes("bucket")) moduleName = "Pranor Vault";
          else if (searchStr.includes("pulse") || searchStr.includes("queue") || searchStr.includes("kafka") || searchStr.includes("stomp") || searchStr.includes("message")) moduleName = "Pranor Pulse";
          else if (searchStr.includes("flow") || searchStr.includes("saga") || searchStr.includes("dag") || searchStr.includes("workflow")) moduleName = "Pranor Flow";
          else if (searchStr.includes("trace") || searchStr.includes("otlp") || searchStr.includes("flamegraph") || searchStr.includes("opentelemetry")) moduleName = "Pranor Trace";
          else if (searchStr.includes("console") || searchStr.includes("dashboard") || searchStr.includes("workbench") || searchStr.includes("soc2")) moduleName = "Pranor Console";
          else if (searchStr.includes("auth") || searchStr.includes("saml") || searchStr.includes("oidc") || searchStr.includes("passkey") || searchStr.includes("mfa")) moduleName = "Pranor Auth";
          else if (searchStr.includes("secret") || searchStr.includes("kms") || searchStr.includes("hsm")) moduleName = "Pranor Secret";
          else if (searchStr.includes("chrono") || searchStr.includes("cron")) moduleName = "Pranor Chrono";
          else if (searchStr.includes("deploy") || searchStr.includes("fleet") || searchStr.includes("canary") || searchStr.includes("blue/green")) moduleName = "Pranor Deploy";
          else if (searchStr.includes("cache") || searchStr.includes("redis") || searchStr.includes("kv")) moduleName = "Pranor Cache";
          else if (searchStr.includes("mesh") || searchStr.includes("mtls") || searchStr.includes("discovery")) moduleName = "Pranor Mesh";
          else if (searchStr.includes("lsp") || searchStr.includes("autocomplete") || searchStr.includes("symbol")) moduleName = "Pranor LSP";
          else if (searchStr.includes("vscode") || searchStr.includes("extension")) moduleName = "Pranor VS Code";

          newFeatures.push({ module: moduleName, edition, name, desc });
        }
      }
    }
  });
});

const merged = existingFeatures.concat(newFeatures);

let jsLines = ["const featuresData = ["];
merged.forEach((item, idx) => {
  const modEsc = item.module.replace(/"/g, '\\"');
  const edEsc = item.edition.replace(/"/g, '\\"');
  const nameEsc = item.name.replace(/"/g, '\\"');
  const descEsc = item.desc.replace(/"/g, '\\"');
  jsLines.push(`  { module: "${modEsc}", edition: "${edEsc}", name: "${nameEsc}", desc: "${descEsc}" }${idx < merged.length - 1 ? "," : ""}`);
});
jsLines.push("];");

const jsCode = jsLines.join("\n");

// Split HTML at const featuresData = [ and renderMatrix
const startIdx = html.indexOf("const featuresData = [");
const endIdx = html.indexOf("let currentTierFilter = 'all';");

if (startIdx === -1 || endIdx === -1) {
  console.error("Could not locate JS injection boundaries!");
  process.exit(1);
}

const prefix = html.substring(0, startIdx);
const suffix = html.substring(endIdx);

const updatedHTML = prefix + jsCode + "\n\n\n" + suffix;
fs.writeFileSync(htmlPath, updatedHTML, "utf8");
console.log("SUCCESS! features.html updated cleanly.");
