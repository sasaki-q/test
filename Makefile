test:
	go test -p 1 -cover ./... -coverprofile=coverage/cover-"$(shell date +%Y-%m-%dT%H:%M)".out

debug-test:
	go test -p 1 -v ./...