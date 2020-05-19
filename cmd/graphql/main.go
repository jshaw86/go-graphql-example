package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/heptiolabs/healthcheck"
	"github.com/jshaw86/go-graphql-example/database"
	"github.com/jshaw86/go-graphql-example/graph"
	"github.com/jshaw86/go-graphql-example/graph/generated"
)

const defaultPort = "8080"

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	databaseConfig := database.Config{
		DatabaseType: getEnv("DATABASE_TYPE", "mysql"),
		Hostname:     getEnv("HOSTNAME", "localhost"),
		Username:     getEnv("USERNAME", "root"),
		Password:     getEnv("PASSWORD", ""),
		Database:     getEnv("DATABASE", "test_db"),
	}

	db := database.InitDB(&databaseConfig)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	http.Handle("/", healthcheck.NewHandler())
	http.Handle("/v1/graphiql", playground.Handler("GraphQL playground", "/v1/query"))
	http.Handle("/v1/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
