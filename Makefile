.PHONY: build-all test-all clean-all

build-all:
	cd .. && go build -o Pranor/pranor.exe ./Pranor
	cd .. && go build -o PranorGate/pranor-gate.exe ./PranorGate
	cd .. && go build -o PranorPulse/pranor-pulse.exe ./PranorPulse
	cd .. && go build -o PranorVault/pranor-vault.exe ./PranorVault
	cd .. && go build -o PranorConsole/pranor-console.exe ./PranorConsole
	cd .. && go build -o PranorMesh/pranor-mesh.exe ./PranorMesh
	cd .. && go build -o PranorAuth/pranor-auth.exe ./PranorAuth
	cd .. && go build -o ServDB/servdb.exe ./ServDB
	cd .. && go build -o PranorFlow/pranor-flow.exe ./PranorFlow
	cd .. && go build -o PranorNotify/pranor-notify.exe ./PranorNotify
	cd .. && go build -o ServDocs/servdocs.exe ./ServDocs

test-all:
	cd .. && go test ./PranorCore/...
	cd .. && go test ./PranorGate/...
	cd .. && go test ./Pranor/...
	cd .. && go test ./PranorMesh/...
	cd .. && go test ./PranorVault/...
	cd .. && go test ./ServDocs/...
	cd .. && go test ./PranorFlow/...

clean-all:
	cd .. && go clean -cache -testcache
