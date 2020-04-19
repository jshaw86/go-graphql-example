package main

import (
    "log"
    "fmt"
    "net/http"
    "github.com/friendsofgo/graphiql"
    graphql "github.com/graph-gophers/graphql-go"
    "github.com/graph-gophers/graphql-go/relay"
    "github.com/heptiolabs/healthcheck"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/jshaw86/go-graphql-example/models"
)
// TODO: Schema
// TODO: Model
type query struct{}
// TODO: Resolver

func (_ *query) Hello() string {
    return "Hello, world!"

}

var db *gorm.DB;

func initDB() {
    var err error
    dataSourceName := "root:@tcp(localhost:3306)/?parseTime=True"
    db, err = gorm.Open("mysql", dataSourceName)

    if err != nil {
        fmt.Println(err)
        panic("failed to connect database")
    }

    db.LogMode(true)

    // Create the database. This is a one-time step.
    // Comment out if running multiple times - You may see an error otherwise
    db.Exec("CREATE DATABASE test_db")
    db.Exec("USE test_db")

    // Migration to create tables for Order and Item schema
    db.AutoMigrate(&models.List{}, &models.Item{})
}

func main() {
    s := `
      type Query {
        hello: String!
      }
    `

    initDB()
    schema := graphql.MustParseSchema(s, &query{})
    http.Handle("/graphql", &relay.Handler{Schema: schema})
    // TODO: graphiql
    // First argument must be same as graphql handler path
    graphiqlHandler, err := graphiql.NewGraphiqlHandler("/query")
    if err != nil {
        panic(err)
    }
    http.Handle("/graphiql", graphiqlHandler)
    // Run
    health := healthcheck.NewHandler()
    http.Handle("/", health);
    http.Handle("/metrics", promhttp.Handler())
    log.Println("Server ready at 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
