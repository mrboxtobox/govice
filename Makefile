BINARY_NAME=achebe

build:
	go build -o ./bin/${BINARY_NAME} achebe.go

run:
	go run achebe.go

# TODO: Add tasks for building releases and adding semantic versioning. 
