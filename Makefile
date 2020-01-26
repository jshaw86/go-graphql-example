

build-example:
	go build -o graphql-example cmd/graphql-example/main.go

run-example:
	go run cmd/graphql-example/main.go 

clean:
	rm graphql-example

.PHONY: clean
