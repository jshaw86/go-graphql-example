package main

import (
    "os"
    "log"
    "net/http"
    "io/ioutil"
    "github.com/friendsofgo/graphiql"
    graphql "github.com/graph-gophers/graphql-go"
    "github.com/graph-gophers/graphql-go/relay"
    "github.com/heptiolabs/healthcheck"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "github.com/jshaw86/go-graphql-example/models"
    "github.com/jshaw86/go-graphql-example"
)

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}


func main() {
    s, err := ioutil.ReadFile("schema.graphql")
    if err != nil {
        panic(err)
    }

    databaseConfig := models.Config{
        DatabaseType: getEnv("DATABASE_TYPE", "mysql"),
        Hostname: getEnv("HOSTNAME","localhost"),
        Username: getEnv("USERNAME","root"),
        Password: getEnv("PASSWORD",""),
        Database: getEnv("DATABASE","test_db"),

    }

    db := models.InitDB(&databaseConfig)
    graphqlResolver := graphqlexample.Resolver{
        DB: db,
    }

    schema := graphql.MustParseSchema(string(s), &graphqlResolver)
    http.Handle("/graphql", &relay.Handler{Schema: schema})
    graphiqlHandler, err := graphiql.NewGraphiqlHandler("/query")
    if err != nil {
        panic(err)
    }
    http.Handle("/graphiql", graphiqlHandler)
    http.Handle("/", healthcheck.NewHandler());
    http.Handle("/metrics", promhttp.Handler())
    log.Println("Server ready at 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
