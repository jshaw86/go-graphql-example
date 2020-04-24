package graphqlexample

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/jshaw86/go-graphql-example/models"
    "github.com/samsarahq/thunder/graphql"
    "github.com/samsarahq/thunder/graphql/schemabuilder"
)

type Resolver struct{
     DB *gorm.DB
}

// registerQuery registers the root query type.
func (r *Resolver) registerQuery(schema *schemabuilder.Schema) {
  obj := schema.Query()

  obj.FieldFunc("getTodoList", func() models.TodoList {
    return models.TodoList{}
  })
}

// registerMutation registers the root mutation type.
func (r *Resolver) registerMutation(schema *schemabuilder.Schema) {
  obj := schema.Mutation()
  obj.FieldFunc("createTodoList", func(args struct{ Message string }) string {
    return "created"
  })

  obj.FieldFunc("addItem", func(args struct{ Message string }) string {
    return "added"
  })
}

// registerPost registers the post type.
func (r *Resolver) registerPost(schema *schemabuilder.Schema) {
  schema.Object("TodoList", models.TodoList{})
  schema.Object("Item", models.Item{})

}

// schema builds the graphql schema.
func (r *Resolver) Schema() *graphql.Schema {
  builder := schemabuilder.NewSchema()
  r.registerQuery(builder)
  r.registerMutation(builder)
  r.registerPost(builder)
  return builder.MustBuild()
}
