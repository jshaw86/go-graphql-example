package graphqlexample 

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/jshaw86/go-graphql-example/models"
    graphql "github.com/graph-gophers/graphql-go"
)

type todolistResolver struct {
    t *models.TodoList
}

func (r *todolistResolver) ID() graphql.ID {
	gID := graphql.ID{}
	gID.UnmarshalGraphQL(r.t.ID)
    return gID
}

func (r *todolistResolver) Name() string {
    return r.t.Name
}

func (r *todolistResolver) Items() *[]*itemResolver {
    return resolveItems(r.t.Items)

}

type itemResolver struct {
    i *models.Item
}

func (r *itemResolver) ID() int {
    return r.i.ID
}

func (r *itemResolver) Name() string {
    return r.i.Name
}

func (r *itemResolver) DueDate() string {
    return r.i.DueDate

}

type Resolver struct{
     DB *gorm.DB
}

func (r *Resolver) GetTodoList(id struct{ ID graphql.ID }) *models.TodoList {
	var itemArray []models.Item;
    return &models.TodoList{ID:1, Name: "Something", Items:itemArray}

}

func resolveItems(ids []models.Item) *[]*itemResolver {
    var items []*itemResolver
    for _, id := range ids {
        if i := resolveItem(id.ID); i != nil {
            items = append(items, i)
        }
    }
    return &items
}

func resolveItem(id int) *itemResolver {
        return &itemResolver{&models.Item{}}

}
