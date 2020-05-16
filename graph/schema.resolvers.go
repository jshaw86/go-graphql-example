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

func (r *mutationResolver) CreateTodo(ctx context.Context, todo model.NewTodo, items []*model.NewItem) (*model.TodoList, error) {
	todoList := database.CreateTodoList(r.DB, todo.Name)
	var databaseItems []*database.Item
	for _, item := range items {
		databaseItem := database.CreateItem(r.DB, todoList.ID, item.Name, item.DueDate)
		databaseItems = append(databaseItems, databaseItem)
	}

	return transformations.ToGraphQLTodoList(todoList, databaseItems)
}

func (r *mutationResolver) AddItems(ctx context.Context, todoListID int, items []*model.NewItem) ([]*model.Item, error) {
	var itemsToReturn []*model.Item
	for _, item := range items {
		itemToReturn, dbItemErr := transformations.ToGraphQLItem(database.CreateItem(r.DB, todoListID, item.Name, item.DueDate))
		if dbItemErr != nil {
			return nil, dbItemErr
		}
		itemsToReturn = append(itemsToReturn, itemToReturn)
	}
	return itemsToReturn, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.TodoList, error) {
	var todoLists []*database.TodoList
	var todoListsItems []*database.Item
	r.DB.Find(&todoLists)
	var todoListIds []int
	for _, todoList := range todoLists {
		todoListIds = append(todoListIds, todoList.ID)
	}
	r.DB.Where("todo_list_id IN (?)", todoListIds).Find(&todoListsItems)

	return transformations.ToGraphQLTodoLists(todoLists, todoListsItems)
}

func (r *queryResolver) Todo(ctx context.Context, id int) (*model.TodoList, error) {
	var todoList database.TodoList
	var todoListItems []*database.Item
	r.DB.Where("id = ?", id).First(&todoList)
	r.DB.Where("todo_list_id = ?", id).Find(&todoListItems)
	return transformations.ToGraphQLTodoList(&todoList, todoListItems)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
