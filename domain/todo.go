package domain

import (
	"context"
	"time"
)

type Todo struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Author      int64     `json:"author"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TodoRepository interface {
	Fetch(ctx context.Context, page, pageSize, authorId int64) ([]Todo, error)
	Insert(ctx context.Context, todo Todo) (Todo, error)
	GetByID(ctx context.Context, id, author int64) (Todo, error)
	Update(ctx context.Context, todo Todo) (Todo, error)
	Delete(ctx context.Context, id int64) error
}

type TodoUsecase interface {
	Fetch(ctx context.Context, page, pageSize, authorId int64) (res []Todo, err error)
	Insert(ctx context.Context, todo Todo) (Todo, error)
	GetByID(ctx context.Context, id, author int64) (Todo, error)
	Update(ctx context.Context, todo Todo) (Todo, error)
	Delete(ctx context.Context, id int64) error
}
