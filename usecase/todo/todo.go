package todo

import (
	"context"
	"time"

	"go-clean-arch-todo.com/domain"
)

type todoUsecase struct {
	todoRepo domain.TodoRepository
}

func NewTodoUsecase(t domain.TodoRepository) domain.TodoUsecase {
	return &todoUsecase{
		todoRepo: t,
	}
}

func (t *todoUsecase) Fetch(ctx context.Context, page, pageSize, authorId int64) ([]domain.Todo, error) {
	page = (page - 1) * pageSize

	todos, err := t.todoRepo.Fetch(ctx, page, pageSize, authorId)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (t *todoUsecase) Insert(ctx context.Context, todo domain.Todo) (domain.Todo, error) {
	todo, err := t.todoRepo.Insert(ctx, todo)
	if err != nil {
		return domain.Todo{}, err
	}

	return todo, nil
}

func (t *todoUsecase) GetByID(ctx context.Context, id, author int64) (domain.Todo, error) {
	todo, err := t.todoRepo.GetByID(ctx, id, author)
	if err != nil {
		return domain.Todo{}, err
	}

	return todo, nil
}

func (t *todoUsecase) Update(ctx context.Context, todo domain.Todo) (domain.Todo, error) {
	todo.UpdatedAt = time.Now()

	todo, err := t.todoRepo.Update(ctx, todo)
	if err != nil {
		return domain.Todo{}, err
	}

	return todo, nil
}

func (t *todoUsecase) Delete(ctx context.Context, id int64) error {
	err := t.todoRepo.Delete(ctx, id)
	return err
}
