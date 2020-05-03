package graphqlexample

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/jshaw86/go-graphql-example/models"
    "github.com/samsarahq/thunder/graphql"
    "github.com/samsarahq/thunder/graphql/schemabuilder"
)

type Resolver struct{
     DB *gorm.DB
}

type NewItem struct{
    Name     string
    DueDate  string
}

type NewTodoList struct{
    Name string
    Items []NewItem

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
  obj.FieldFunc("createTodoList", func(todoList struct{
      Name string
      Items []*models.Item
  }) *models.TodoList {

      fmt.Println("todolist... %+v", todoList)

      /*
    todoListItems := make([]models.Item, len(NewTodoList.Items))
    for i := range todoList.Items {
        append(todoListItems, models.Item{Name:i.Name, DueDate: i.DueDate})
    }*/

    return models.CreateTodoList(r.DB, todoList.Name,todoList.Items...)
  })

  obj.FieldFunc("addItem", func(args struct{ Message string }) string {
    return "added"
  })
}

func (r *Resolver) registerTodoList(schema *schemabuilder.Schema) {
  object := schema.Object("TodoList", models.TodoList{})

  object.FieldFunc("items", func(args struct{ Message string}) []models.Item {
      first := models.Item{}
      second := models.Item{}
      return []models.Item{first, second}
  })

}

func (r *Resolver) registerItem(schema *schemabuilder.Schema) {
  schema.Object("Item", models.Item{})

}
// schema builds the graphql schema.
func (r *Resolver) Schema() *graphql.Schema {
  builder := schemabuilder.NewSchema()
  r.registerQuery(builder)
  r.registerMutation(builder)
  r.registerTodoList(builder)
  r.registerItem(builder)
  return builder.MustBuild()
}
