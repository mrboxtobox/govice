BINARY_NAME=achebe

build:
	go build -o ./bin/${BINARY_NAME} achebe.go

run:
	go run achebe.go

clean: ## Remove build related file
	rm -fr ./bin

build_release:
	GOOS=windows GOARCH=amd64 go build -o bin/achebe-amd64-v0.0.0.exe achebe.go
	GOOS=windows GOARCH=386 go build -o bin/achebe-386-v0.0.0.exe achebe.go
	GOOS=darwin GOARCH=amd64 go build -o bin/achebe-amd64-darwin-v0.0.0 achebe.go
	GOOS=linux GOARCH=amd64 go build -o bin/achebe-amd64-linux-v0.0.0 achebe.go
	GOOS=linux GOARCH=386 go build -o bin/achebe-386-linux-v0.0.0 achebe.go

# TODO: Add tasks for building releases and adding semantic versioning. 
