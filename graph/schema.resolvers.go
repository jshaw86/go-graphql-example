package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/jshaw86/go-graphql-example/database"
	"github.com/jshaw86/go-graphql-example/graph/generated"
	"github.com/jshaw86/go-graphql-example/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.TodoList, error) {
	return database.CreateTodoList(r.DB, input.Name), nil
}

func (r *mutationResolver) AddItem(ctx context.Context, input model.NewItem) (*model.Item, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.TodoList, error) {
	var todoLists []*database.TodoList
	r.DB.Find(&todoLists)
    return todoLists, nil
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*model.TodoList, error) {
	var todoList database.TodoList
	r.DB.Where("id = ?", id).First(&todoList)
    return todoList, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
