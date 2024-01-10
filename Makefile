BINARY_NAME=govice

build:
	go build -o ./bin/${BINARY_NAME} govice.go

run:
	go run govice.go

clean: ## Remove build related file
	rm -fr ./bin

build_release:
	GOOS=windows GOARCH=amd64 go build -o bin/govice-amd64-v0.0.0.exe govice.go
	GOOS=windows GOARCH=386 go build -o bin/govice-386-v0.0.0.exe govice.go
	GOOS=darwin GOARCH=amd64 go build -o bin/govice-amd64-darwin-v0.0.0 govice.go
	GOOS=linux GOARCH=amd64 go build -o bin/govice-amd64-linux-v0.0.0 govice.go
	GOOS=linux GOARCH=386 go build -o bin/govice-386-linux-v0.0.0 govice.go

# TODO: Add tasks for building releases and adding semantic versioning. 
