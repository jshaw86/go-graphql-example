package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/jshaw86/go-graphql-example/database"
	"github.com/jshaw86/go-graphql-example/graph/generated"
	"github.com/jshaw86/go-graphql-example/graph/model"
	"github.com/jshaw86/go-graphql-example/transformations"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.TodoList, error) {
	return transformations.ToGraphQLTodoList(database.CreateTodoList(r.DB, input.Name))
}

func (r *mutationResolver) AddItem(ctx context.Context, input model.NewItem) (*model.Item, error) {
	return transformations.ToGraphQLItem(database.CreateItem(r.DB, input.TodoListID, input.Name, input.DueDate))
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.TodoList, error) {
	var todoLists []*database.TodoList
	r.DB.Find(&todoLists)

	return transformations.ToGraphQLTodoLists(todoLists)
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*model.TodoList, error) {
	var todoList database.TodoList
	r.DB.Where("id = ?", id).First(&todoList)
	return transformations.ToGraphQLTodoList(&todoList)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
