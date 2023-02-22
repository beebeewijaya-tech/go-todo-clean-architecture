package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"
	"go-clean-arch-todo.com/domain"
)

type postgresTodoRepository struct {
	Store *sqlx.DB
}

func NewPostgresTodoRepository(db *sqlx.DB) domain.TodoRepository {
	return &postgresTodoRepository{
		Store: db,
	}
}

func (p *postgresTodoRepository) Fetch(ctx context.Context, page, pageSize, authorId int64) ([]domain.Todo, error) {
	query := `SELECT * FROM todos WHERE author=$3 LIMIT $2 OFFSET $1`
	rows, err := p.Store.DB.QueryContext(ctx, query, page, pageSize, authorId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var todos []domain.Todo

	for rows.Next() {
		var t domain.Todo

		if err := rows.Scan(
			&t.ID,
			&t.Title,
			&t.Description,
			&t.Author,
			&t.CreatedAt,
			&t.UpdatedAt,
		); err != nil {
			return nil, err
		}

		todos = append(todos, t)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func (p *postgresTodoRepository) Insert(ctx context.Context, todo domain.Todo) (domain.Todo, error) {
	query := `
	INSERT INTO todos (
		title,
		description,
		author
	) VALUES (
		$1,
		$2,
		$3
	) RETURNING *
	`

	row := p.Store.DB.QueryRowContext(ctx, query, todo.Title, todo.Description, todo.Author)

	var t domain.Todo

	err := row.Scan(
		&t.ID,
		&t.Title,
		&t.Description,
		&t.Author,
		&t.CreatedAt,
		&t.UpdatedAt,
	)

	return t, err
}

func (p *postgresTodoRepository) GetByID(ctx context.Context, id, author int64) (domain.Todo, error) {
	query := `SELECT * FROM todos where id = $1 AND author = $2`

	row := p.Store.DB.QueryRowContext(ctx, query, id, author)

	var t domain.Todo

	err := row.Scan(
		&t.ID,
		&t.Title,
		&t.Description,
		&t.Author,
		&t.CreatedAt,
		&t.UpdatedAt,
	)

	return t, err
}

func (p *postgresTodoRepository) Update(ctx context.Context, todo domain.Todo) (domain.Todo, error) {
	query := `
	UPDATE todos SET
	title = $2,
	description = $3,
	updated_at = $4
	WHERE id = $1
	RETURNING *
	`
	row := p.Store.DB.QueryRowContext(ctx, query, todo.ID, todo.Title, todo.Description, todo.UpdatedAt)

	var t domain.Todo

	err := row.Scan(
		&t.ID,
		&t.Title,
		&t.Description,
		&t.Author,
		&t.CreatedAt,
		&t.UpdatedAt,
	)

	return t, err
}

func (p *postgresTodoRepository) Delete(ctx context.Context, id int64) error {
	query := `
	DELETE FROM todos WHERE id = $1
	`

	_, err := p.Store.DB.ExecContext(ctx, query, id)

	return err
}
