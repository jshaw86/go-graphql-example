package main

import (
    "os"
    "log"
    "net/http"
    "github.com/heptiolabs/healthcheck"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "github.com/jshaw86/go-graphql-example/models"
    graphqlexample "github.com/jshaw86/go-graphql-example"
    "github.com/samsarahq/thunder/graphql"
    "github.com/samsarahq/thunder/graphql/graphiql"
    "github.com/samsarahq/thunder/graphql/introspection"

)

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}

func main() {
    databaseConfig := models.Config{
        DatabaseType: getEnv("DATABASE_TYPE", "mysql"),
        Hostname: getEnv("HOSTNAME","localhost"),
        Username: getEnv("USERNAME","root"),
        Password: getEnv("PASSWORD",""),
        Database: getEnv("DATABASE","test_db"),

    }

    db := models.InitDB(&databaseConfig)
    resolver := graphqlexample.Resolver{DB: db}
    schema := resolver.Schema()
    introspection.AddIntrospectionToSchema(schema)

    // Run
    http.Handle("/graphql", graphql.Handler(schema))
    http.Handle("/graphiql/", http.StripPrefix("/graphiql/", graphiql.Handler()))
    http.Handle("/", healthcheck.NewHandler());
    http.Handle("/metrics", promhttp.Handler())
    log.Println("Server ready at 3030")
    log.Fatal(http.ListenAndServe(":3030", nil))
}
