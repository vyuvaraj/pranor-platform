.PHONY: build-all test-all clean-all

build-all:
	cd .. && go build -o Serv-lang/pranor.exe ./Serv-lang
	cd .. && go build -o ServGate/servgate.exe ./ServGate
	cd .. && go build -o ServQueue/servqueue.exe ./ServQueue
	cd .. && go build -o ServStore/servstore.exe ./ServStore
	cd .. && go build -o ServConsole/servconsole.exe ./ServConsole
	cd .. && go build -o ServMesh/servmesh.exe ./ServMesh
	cd .. && go build -o ServAuth/servauth.exe ./ServAuth
	cd .. && go build -o ServDB/servdb.exe ./ServDB
	cd .. && go build -o ServFlow/servflow.exe ./ServFlow
	cd .. && go build -o ServMail/servmail.exe ./ServMail
	cd .. && go build -o ServDocs/servdocs.exe ./ServDocs

test-all:
	cd .. && go test ./ServShared/...
	cd .. && go test ./ServGate/...
	cd .. && go test ./Serv-lang/...
	cd .. && go test ./ServMesh/...
	cd .. && go test ./ServStore/...
	cd .. && go test ./ServDocs/...
	cd .. && go test ./ServFlow/...

clean-all:
	cd .. && go clean -cache -testcache
