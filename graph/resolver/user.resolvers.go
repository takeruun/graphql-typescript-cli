package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/entity"
	"app/graph/model"
	"context"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	var user entity.User

	newUser := &entity.User{
		Name: input.Name,
	}
	result := r.DB.Create(newUser)
	if result.Error != nil {
		return nil, result.Error
	}

	result = r.DB.Find(&user, newUser.ID)
	if result.Error != nil {
		return nil, result.Error
	}

	return entity.ToModelUser(&user), nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var users []*entity.User

	sql := r.DB.Model(&model.User{})
	result := sql.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return entity.ToModelUsers(users), nil
}
