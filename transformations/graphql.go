package transformations

import (
	"github.com/jshaw86/go-graphql-example/graph/model"
	"github.com/jshaw86/go-graphql-example/database"
)

func ToGraphQLTodoList(todoList *database.TodoList) (*model.TodoList, error) {
    todoListModel := model.TodoList{}

    return &todoListModel, nil

}

func ToGraphQLTodoLists(todoLists []*database.TodoList) ([]*model.TodoList, error) {
	var graphQLTodoLists []*model.TodoList
	for _, todoList := range todoLists {
		graphQLTodoList, _ := ToGraphQLTodoList(todoList)
		graphQLTodoLists = append(graphQLTodoLists, graphQLTodoList)

	}
	return graphQLTodoLists, nil
}

func ToGraphQLItem(item *database.Item) (*model.Item, error) {
    itemModel := model.Item{}

    return &itemModel, nil
}
