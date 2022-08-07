package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/entity"
	"app/graph/generated"
	"app/graph/model"
	"context"
	"strconv"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	var todo entity.Todo

	userId, _ := strconv.Atoi(input.UserID)
	newTodo := &entity.Todo{
		Text:   input.Text,
		UserID: uint64(userId),
	}
	result := r.DB.Create(newTodo)
	if result.Error != nil {
		return nil, result.Error
	}

	result = r.DB.Preload("User").Find(&todo, newTodo.ID)
	if result.Error != nil {
		return nil, result.Error
	}

	return entity.ToModelTodo(&todo), nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	var todos []*entity.Todo

	sql := r.DB.Model(&entity.Todo{})
	result := sql.Preload("User").Find(&todos)
	if result.Error != nil {
		return nil, result.Error
	}

	return entity.ToModelTodos(todos), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
