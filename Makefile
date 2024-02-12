# Use the swag formatter to automatically format the docs comments
docfmt:
	@swag fmt -d cmd/api,pkg

# Generate docs, copy into package dir for embedding in the binary
docgen:
	@swag init -g cmd/api/main.go -o cmd/api/dist --ot yaml --v3.1

run:
	go run ./cmd/api

all: docfmt docgen run