; Inno Setup Script for Unified PranorVerse Windows Setup
; Generates PranorVerse-setup.exe with component selection of the 16 Pranorverse micropranorices.

[Setup]
AppName=PranorVerse
AppVersion=1.7.0
ArchitecturesAllowed=x64
ArchitecturesInstallIn64BitMode=x64
DefaultDirName={autopf}\PranorVerse
DefaultGroupName=PranorVerse
UninstallDisplayIcon={app}\bin\pranor.exe
Compression=lzma2
SolidCompression=yes
OutputDir=dist
OutputBaseFilename=PranorVerse-windows-setup
ChangesEnvironment=yes

[Types]
Name: "full"; Description: "Full installation (All 16 pranorices & tools)"
Name: "custom"; Description: "Custom installation"; Flags: iscustom

[Components]
Name: "compiler"; Description: "Pranor-lang CLI & Compiler Runtime"; Types: full custom; Flags: fixed
Name: "gateway"; Description: "PranorGate API Gateway proxy"; Types: full custom
Name: "store"; Description: "PranorStore S3 Object Storage engine"; Types: full custom
Name: "queue"; Description: "PranorQueue STOMP Message Broker"; Types: full custom
Name: "cache"; Description: "PranorCache REST caching pranorice"; Types: full custom
Name: "mesh"; Description: "PranorMesh Pranorice Mesh pranorice"; Types: full custom
Name: "console"; Description: "PranorConsole Management Dashboard"; Types: full custom
Name: "trace"; Description: "PranorTrace OpenTelemetry trace collector"; Types: full custom
Name: "auth"; Description: "PranorAuth Identity & Security pranorice"; Types: full custom
Name: "cron"; Description: "PranorCron Distributed Scheduler"; Types: full custom
Name: "cloud"; Description: "PranorCloud Cluster Orchestrator"; Types: full custom
Name: "flow"; Description: "PranorFlow Saga Workflow engine"; Types: full custom
Name: "pool"; Description: "PranorPool Connection Pooler proxy"; Types: full custom
Name: "mail"; Description: "PranorMail Notification API Gateway"; Types: full custom
Name: "tunnel"; Description: "PranorTunnel Secure Localhost Relay Tunnel"; Types: full custom
Name: "registry"; Description: "PranorRegistry Package Module Registry"; Types: full custom
Name: "lock"; Description: "PranorLock Distributed Locking pranorice"; Types: full custom

[Dirs]
Name: "{app}\bin"
Name: "{app}\conf"
Name: "{app}\conf\pranorgate"
Name: "{app}\conf\pranorconsole"
Name: "{app}\logs"

[Files]
; Binaries - sourced from dist/bin/ folder populated by workflow
Source: "dist\bin\pranor.exe"; DestDir: "{app}\bin"; Components: compiler; Flags: ignoreversion
Source: "dist\bin\pranorgate.exe"; DestDir: "{app}\bin"; Components: gateway; Flags: ignoreversion
Source: "dist\bin\pranorstore.exe"; DestDir: "{app}\bin"; Components: store; Flags: ignoreversion
Source: "dist\bin\pranorqueue.exe"; DestDir: "{app}\bin"; Components: queue; Flags: ignoreversion
Source: "dist\bin\pranorconsole.exe"; DestDir: "{app}\bin"; Components: console; Flags: ignoreversion
Source: "dist\bin\pranormesh.exe"; DestDir: "{app}\bin"; Components: mesh; Flags: ignoreversion
Source: "dist\bin\pranorauth.exe"; DestDir: "{app}\bin"; Components: auth; Flags: ignoreversion
Source: "dist\bin\pranorcloud.exe"; DestDir: "{app}\bin"; Components: cloud; Flags: ignoreversion
Source: "dist\bin\pranortrace.exe"; DestDir: "{app}\bin"; Components: trace; Flags: ignoreversion
Source: "dist\bin\pranortunnel.exe"; DestDir: "{app}\bin"; Components: tunnel; Flags: ignoreversion
Source: "dist\bin\pranorflow.exe"; DestDir: "{app}\bin"; Components: flow; Flags: ignoreversion
Source: "dist\bin\pranordb.exe"; DestDir: "{app}\bin"; Components: pool; Flags: ignoreversion
Source: "dist\bin\pranormail.exe"; DestDir: "{app}\bin"; Components: mail; Flags: ignoreversion
Source: "dist\bin\pranorcache.exe"; DestDir: "{app}\bin"; Components: cache; Flags: ignoreversion
Source: "dist\bin\pranorcron.exe"; DestDir: "{app}\bin"; Components: cron; Flags: ignoreversion
Source: "dist\bin\pranorregistry.exe"; DestDir: "{app}\bin"; Components: registry; Flags: ignoreversion
Source: "dist\bin\pranorlock.exe"; DestDir: "{app}\bin"; Components: lock; Flags: ignoreversion

; Configuration Templates - bundled inside the repo under configs/
Source: "configs\pranorgate\config.json"; DestDir: "{app}\conf\pranorgate"; Flags: onlyifdoesntexist
Source: "configs\pranorconsole\pranorices.json"; DestDir: "{app}\conf\pranorconsole"; Flags: onlyifdoesntexist

[Registry]
Root: HKCU; Subkey: "Environment"; ValueType: string; ValueName: "Path"; ValueData: "{olddata};{app}\bin"; Flags: prepranorestringtype

[Icons]
Name: "{group}\PranorConsole Dashboard"; Filename: "{app}\bin\pranorconsole.exe"; Components: console
Name: "{group}\Uninstall PranorVerse"; Filename: "{uninstallexe}"
