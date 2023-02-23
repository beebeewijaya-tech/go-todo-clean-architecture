package todo

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go-clean-arch-todo.com/domain"
	"go-clean-arch-todo.com/domain/mock_domain"
)

func TestFetchTodos(t *testing.T) {
	testCases := []struct {
		name          string
		mockFunc      func(mock *mock_domain.MockTodoRepository)
		checkResponse func(t *testing.T, response []domain.Todo, err error)
	}{
		{
			name: "OK",
			mockFunc: func(mock *mock_domain.MockTodoRepository) {
				mock.EXPECT().
					Fetch(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1).
					Return([]domain.Todo{
						{
							Title:       "Todo 1",
							Description: "Todo 1",
							Author:      1,
							CreatedAt:   time.Now(),
							UpdatedAt:   time.Now(),
						},
						{
							Title:       "Todo 2",
							Description: "Todo 2",
							Author:      1,
							CreatedAt:   time.Now(),
							UpdatedAt:   time.Now(),
						},
						{
							Title:       "Todo 3",
							Description: "Todo 3",
							Author:      1,
							CreatedAt:   time.Now(),
							UpdatedAt:   time.Now(),
						},
					}, nil)
			},
			checkResponse: func(t *testing.T, response []domain.Todo, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, response)
			},
		},
		{
			name: "NotOK",
			mockFunc: func(mock *mock_domain.MockTodoRepository) {
				mock.EXPECT().
					Fetch(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1).
					Return(nil, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, response []domain.Todo, err error) {
				require.Error(t, err)
				require.EqualError(t, err, sql.ErrConnDone.Error())
				require.Empty(t, response)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			todoRepoMock := mock_domain.NewMockTodoRepository(ctrl)
			todoUsecase := NewTodoUsecase(todoRepoMock)

			tc.mockFunc(todoRepoMock)

			page := 1
			pageSize := 2
			author := 1
			res, err := todoUsecase.Fetch(context.Background(), int64(page), int64(pageSize), int64(author))
			tc.checkResponse(t, res, err)
		})
	}
}

func TestInsertTodo(t *testing.T) {
	testCases := []struct {
		name          string
		mockFunc      func(mock *mock_domain.MockTodoRepository)
		checkResponse func(t *testing.T, response domain.Todo, err error)
	}{
		{
			name: "OK",
			mockFunc: func(mock *mock_domain.MockTodoRepository) {
				mock.EXPECT().
					Insert(gomock.Any(), gomock.Any()).
					Times(1).
					Return(domain.Todo{
						Title:       "Todo 1",
						Description: "Todo 1",
						Author:      1,
						CreatedAt:   time.Now(),
						UpdatedAt:   time.Now(),
					}, nil)
			},
			checkResponse: func(t *testing.T, response domain.Todo, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, response)
			},
		},
		{
			name: "NotOK",
			mockFunc: func(mock *mock_domain.MockTodoRepository) {
				mock.EXPECT().
					Insert(gomock.Any(), gomock.Any()).
					Times(1).
					Return(domain.Todo{}, sql.ErrTxDone)
			},
			checkResponse: func(t *testing.T, response domain.Todo, err error) {
				require.Error(t, err)
				require.EqualError(t, err, sql.ErrTxDone.Error())
				require.Empty(t, response)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			todoRepoMock := mock_domain.NewMockTodoRepository(ctrl)
			todoUsecase := NewTodoUsecase(todoRepoMock)

			tc.mockFunc(todoRepoMock)

			todo := domain.Todo{
				Title:       "Todo 1",
				Description: "Todo 1",
				Author:      1,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}
			todo, err := todoUsecase.Insert(context.Background(), todo)
			tc.checkResponse(t, todo, err)
		})
	}
}

func TestGetByID(t *testing.T) {
	testCases := []struct {
		name          string
		mockFunc      func(mock *mock_domain.MockTodoRepository)
		checkResponse func(t *testing.T, response domain.Todo, err error)
	}{
		{
			name: "OK",
			mockFunc: func(mock *mock_domain.MockTodoRepository) {
				mock.EXPECT().
					GetByID(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1).
					Return(domain.Todo{
						Title:       "Todo 1",
						Description: "Todo 1",
						Author:      1,
						CreatedAt:   time.Now(),
						UpdatedAt:   time.Now(),
					}, nil)
			},
			checkResponse: func(t *testing.T, response domain.Todo, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, response)
			},
		},
		{
			name: "NotOK",
			mockFunc: func(mock *mock_domain.MockTodoRepository) {
				mock.EXPECT().
					GetByID(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1).
					Return(domain.Todo{}, sql.ErrTxDone)
			},
			checkResponse: func(t *testing.T, response domain.Todo, err error) {
				require.Error(t, err)
				require.EqualError(t, err, sql.ErrTxDone.Error())
				require.Empty(t, response)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			todoRepoMock := mock_domain.NewMockTodoRepository(ctrl)
			todoUsecase := NewTodoUsecase(todoRepoMock)

			tc.mockFunc(todoRepoMock)

			todo, err := todoUsecase.GetByID(context.Background(), 1, 1)
			tc.checkResponse(t, todo, err)
		})
	}
}

func TestUpdateTodo(t *testing.T) {
	testCases := []struct {
		name          string
		mockFunc      func(mock *mock_domain.MockTodoRepository)
		checkResponse func(t *testing.T, response domain.Todo, err error)
	}{
		{
			name: "OK",
			mockFunc: func(mock *mock_domain.MockTodoRepository) {
				mock.EXPECT().
					Update(gomock.Any(), gomock.Any()).
					Times(1).
					Return(domain.Todo{
						Title:       "Todo 1",
						Description: "Todo 1",
						Author:      1,
						CreatedAt:   time.Now(),
						UpdatedAt:   time.Now(),
					}, nil)
			},
			checkResponse: func(t *testing.T, response domain.Todo, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, response)
			},
		},
		{
			name: "NotOK",
			mockFunc: func(mock *mock_domain.MockTodoRepository) {
				mock.EXPECT().
					Update(gomock.Any(), gomock.Any()).
					Times(1).
					Return(domain.Todo{}, sql.ErrTxDone)
			},
			checkResponse: func(t *testing.T, response domain.Todo, err error) {
				require.Error(t, err)
				require.EqualError(t, err, sql.ErrTxDone.Error())
				require.Empty(t, response)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			todoRepoMock := mock_domain.NewMockTodoRepository(ctrl)
			todoUsecase := NewTodoUsecase(todoRepoMock)

			tc.mockFunc(todoRepoMock)

			todo := domain.Todo{
				ID:          1,
				Title:       "Todo 1",
				Description: "Todo 1",
			}
			todo, err := todoUsecase.Update(context.Background(), todo)
			tc.checkResponse(t, todo, err)
		})
	}
}

func TestDeleteTodo(t *testing.T) {
	testCases := []struct {
		name          string
		mockFunc      func(mock *mock_domain.MockTodoRepository)
		checkResponse func(t *testing.T, err error)
	}{
		{
			name: "OK",
			mockFunc: func(mock *mock_domain.MockTodoRepository) {
				mock.EXPECT().
					Delete(gomock.Any(), gomock.Any()).
					Times(1).
					Return(nil)
			},
			checkResponse: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		{
			name: "NotOK",
			mockFunc: func(mock *mock_domain.MockTodoRepository) {
				mock.EXPECT().
					Delete(gomock.Any(), gomock.Any()).
					Times(1).
					Return(sql.ErrTxDone)
			},
			checkResponse: func(t *testing.T, err error) {
				require.Error(t, err)
				require.EqualError(t, err, sql.ErrTxDone.Error())
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			todoRepoMock := mock_domain.NewMockTodoRepository(ctrl)
			todoUsecase := NewTodoUsecase(todoRepoMock)

			tc.mockFunc(todoRepoMock)

			err := todoUsecase.Delete(context.Background(), 1)
			tc.checkResponse(t, err)
		})
	}
}
