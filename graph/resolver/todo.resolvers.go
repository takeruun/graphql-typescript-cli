package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/entity"
	"app/graph/model"
	"app/util"
	"context"
	"math"
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
func (r *queryResolver) Todos(ctx context.Context, input model.PageCondition) (*model.TodoConnection, error) {
	var todos []*entity.Todo
	var totalCount int64 = 0
	var totalPage int64 = 0

	sql := r.DB.Model(&entity.Todo{})
	sql.Count(&totalCount)

	targetCount := 0
	if input.First != nil {
		targetCount = *input.First
	}
	if input.Last != nil {
		targetCount = *input.Last
	}

	totalPage = int64(math.Ceil(float64(totalCount) / float64(targetCount)))

	pageInfo := &model.PageInfo{
		HasPreviousPage: (totalPage > 1 && input.PageNo > 1),
		HasNextPage:     (totalPage > 1 && totalPage != int64(input.PageNo)),
	}

	limit := 5
	if input.First != nil {
		limit = *input.First
	} else {
		limit = *input.Last
	}

	sql = sql.Preload("User")
	if input.After != nil {
		_, afterId, err := util.DecodeCursor(*input.After)
		if err != nil {
			return nil, err
		}
		sql.Where("id > ?", afterId)
	}
	if input.Before != nil {
		_, beforeId, err := util.DecodeCursor(*input.Before)
		if err != nil {
			return nil, err
		}
		if input.First != nil {
			sql.Where("id >= ?", beforeId)
		} else {
			sql.Where("id < ?", beforeId)
		}
	}

	result := sql.Preload("User").
		Limit(limit).
		Find(&todos)
	if result.Error != nil {
		return nil, result.Error
	}

	var edges []*model.TodoEdge
	for idx, todo := range todos {
		cursor := util.CreateCursor("todo", strconv.Itoa(int(todo.ID)))
		edges = append(edges, &model.TodoEdge{
			Cursor: cursor,
			Node:   entity.ToModelTodo(todo),
		})

		if idx == 0 {
			// 今回表示するページの１件目のレコードを特定するカーソルをセット
			// （「前ページ」遷移時に「このカーソルよりも前のレコード」という検索条件に用いる）
			pageInfo.StartCursor = &cursor
		}
		if idx == len(todos)-1 {
			// 今回表示するページの最後のレコードを特定するカーソルをセット
			// （「次ページ」遷移時に「このカーソルよりも後のレコード」という検索条件に用いる）
			pageInfo.EndCursor = &cursor
		}
	}

	return &model.TodoConnection{
		PageInfo:   pageInfo,
		Edges:      edges,
		TotalCount: int(totalCount),
	}, nil
}
