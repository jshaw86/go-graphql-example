package main

import (
    "log"
    "net/http"
    "github.com/friendsofgo/graphiql"
    graphql "github.com/graph-gophers/graphql-go"
    "github.com/graph-gophers/graphql-go/relay"
)
// TODO: Schema
// TODO: Model
type query struct{}
// TODO: Resolver

func (_ *query) Hello() string {
    return "Hello, world!"

}

func main() {
    s := `
		type Query {
          hello: String!
        }
    `
    schema := graphql.MustParseSchema(s, &query{})
    http.Handle("/graphql", &relay.Handler{Schema: schema})
    // TODO: graphiql
    // First argument must be same as graphql handler path
    graphiqlHandler, err := graphiql.NewGraphiqlHandler("/query")
    if err != nil {
        panic(err)
    }
    http.Handle("/", graphiqlHandler)
    // Run
    log.Println("Server ready at 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
