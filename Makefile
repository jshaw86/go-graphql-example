build:
	go build -o graphql cmd/graphql/main.go

run:
	go run cmd/graphql/main.go 

clean:
	rm graphql

.PHONY: clean
