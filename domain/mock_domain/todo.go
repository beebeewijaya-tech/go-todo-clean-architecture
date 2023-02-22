// Code generated by MockGen. DO NOT EDIT.
// Source: .\domain\todo.go

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "go-clean-arch-todo.com/domain"
)

// MockTodoRepository is a mock of TodoRepository interface.
type MockTodoRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTodoRepositoryMockRecorder
}

// MockTodoRepositoryMockRecorder is the mock recorder for MockTodoRepository.
type MockTodoRepositoryMockRecorder struct {
	mock *MockTodoRepository
}

// NewMockTodoRepository creates a new mock instance.
func NewMockTodoRepository(ctrl *gomock.Controller) *MockTodoRepository {
	mock := &MockTodoRepository{ctrl: ctrl}
	mock.recorder = &MockTodoRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodoRepository) EXPECT() *MockTodoRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockTodoRepository) Delete(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockTodoRepositoryMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTodoRepository)(nil).Delete), ctx, id)
}

// Fetch mocks base method.
func (m *MockTodoRepository) Fetch(ctx context.Context, page, pageSize, authorId int64) ([]domain.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", ctx, page, pageSize, authorId)
	ret0, _ := ret[0].([]domain.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch.
func (mr *MockTodoRepositoryMockRecorder) Fetch(ctx, page, pageSize, authorId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockTodoRepository)(nil).Fetch), ctx, page, pageSize, authorId)
}

// GetByID mocks base method.
func (m *MockTodoRepository) GetByID(ctx context.Context, id, author int64) (domain.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id, author)
	ret0, _ := ret[0].(domain.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockTodoRepositoryMockRecorder) GetByID(ctx, id, author interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockTodoRepository)(nil).GetByID), ctx, id, author)
}

// Insert mocks base method.
func (m *MockTodoRepository) Insert(ctx context.Context, todo domain.Todo) (domain.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, todo)
	ret0, _ := ret[0].(domain.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockTodoRepositoryMockRecorder) Insert(ctx, todo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockTodoRepository)(nil).Insert), ctx, todo)
}

// Update mocks base method.
func (m *MockTodoRepository) Update(ctx context.Context, todo domain.Todo) (domain.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, todo)
	ret0, _ := ret[0].(domain.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockTodoRepositoryMockRecorder) Update(ctx, todo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTodoRepository)(nil).Update), ctx, todo)
}

// MockTodoUsecase is a mock of TodoUsecase interface.
type MockTodoUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockTodoUsecaseMockRecorder
}

// MockTodoUsecaseMockRecorder is the mock recorder for MockTodoUsecase.
type MockTodoUsecaseMockRecorder struct {
	mock *MockTodoUsecase
}

// NewMockTodoUsecase creates a new mock instance.
func NewMockTodoUsecase(ctrl *gomock.Controller) *MockTodoUsecase {
	mock := &MockTodoUsecase{ctrl: ctrl}
	mock.recorder = &MockTodoUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodoUsecase) EXPECT() *MockTodoUsecaseMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockTodoUsecase) Delete(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockTodoUsecaseMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTodoUsecase)(nil).Delete), ctx, id)
}

// Fetch mocks base method.
func (m *MockTodoUsecase) Fetch(ctx context.Context, page, pageSize, authorId int64) ([]domain.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", ctx, page, pageSize, authorId)
	ret0, _ := ret[0].([]domain.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch.
func (mr *MockTodoUsecaseMockRecorder) Fetch(ctx, page, pageSize, authorId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockTodoUsecase)(nil).Fetch), ctx, page, pageSize, authorId)
}

// GetByID mocks base method.
func (m *MockTodoUsecase) GetByID(ctx context.Context, id, author int64) (domain.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id, author)
	ret0, _ := ret[0].(domain.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockTodoUsecaseMockRecorder) GetByID(ctx, id, author interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockTodoUsecase)(nil).GetByID), ctx, id, author)
}

// Insert mocks base method.
func (m *MockTodoUsecase) Insert(ctx context.Context, todo domain.Todo) (domain.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, todo)
	ret0, _ := ret[0].(domain.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockTodoUsecaseMockRecorder) Insert(ctx, todo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockTodoUsecase)(nil).Insert), ctx, todo)
}

// Update mocks base method.
func (m *MockTodoUsecase) Update(ctx context.Context, todo domain.Todo) (domain.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, todo)
	ret0, _ := ret[0].(domain.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockTodoUsecaseMockRecorder) Update(ctx, todo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTodoUsecase)(nil).Update), ctx, todo)
}