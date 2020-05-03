package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jshaw86/go-graphql-example/graph"
	"github.com/jshaw86/go-graphql-example/graph/generated"
    "github.com/heptiolabs/healthcheck"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "github.com/jshaw86/go-graphql-example/models"
)

const defaultPort = "8080"

var db *gorm.DB;

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

    databaseConfig := models.Config{
        DatabaseType: getEnv("DATABASE_TYPE", "mysql"),
        Hostname: getEnv("HOSTNAME","localhost"),
        Username: getEnv("USERNAME","root"),
        Password: getEnv("PASSWORD",""),
        Database: getEnv("DATABASE","test_db"),

    }

    db := models.InitDB(&databaseConfig)
    srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

    http.Handle("/", healthcheck.NewHandler());
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
