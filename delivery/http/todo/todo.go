package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-clean-arch-todo.com/domain"
	"go-clean-arch-todo.com/util"
)

const (
	// Dummy author
	author = 1
)

type todoHandler struct {
	todoUsecase domain.TodoUsecase
}

type fetchTodoArgs struct {
	Page     int64 `form:"page" binding:"required,min=1"`
	PageSize int64 `form:"page_size" binding:"required,min=10"`
}

type insertTodoBody struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description,omitempty"`
}

type updateTodoBody struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type getTodoParams struct {
	ID int64 `uri:"id" binding:"required,min=0"`
}

type deleteTodoResponse struct {
	ID      int64  `json:"id"`
	Message string `json:"message"`
}

func NewTodoHandler(r *gin.Engine, t domain.TodoUsecase) {
	h := &todoHandler{
		todoUsecase: t,
	}

	r.GET("/todos", h.fetchTodo)
	r.GET("/todo/:id", h.getTodoByID)
	r.POST("/todo", h.insertTodo)
	r.PUT("/todo/:id", h.updateTodo)
	r.DELETE("/todo/:id", h.deleteTodo)
}

func (t *todoHandler) fetchTodo(ctx *gin.Context) {
	var req fetchTodoArgs
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	todos, err := t.todoUsecase.Fetch(ctx, req.Page, req.PageSize, author)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, todos)
}

func (t *todoHandler) insertTodo(ctx *gin.Context) {
	var req insertTodoBody
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	td := domain.Todo{
		Title:       req.Title,
		Description: req.Description,
		Author:      author,
	}

	todo, err := t.todoUsecase.Insert(ctx, td)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

func (t *todoHandler) getTodoByID(ctx *gin.Context) {
	var req getTodoParams
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	todo, err := t.todoUsecase.GetByID(ctx, req.ID, author)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

func (t *todoHandler) updateTodo(ctx *gin.Context) {
	var p getTodoParams
	if err := ctx.ShouldBindUri(&p); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	var req updateTodoBody
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	td := domain.Todo{
		ID:          p.ID,
		Title:       req.Title,
		Description: req.Description,
		Author:      author,
	}
	todo, err := t.todoUsecase.Update(ctx, td)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

func (t *todoHandler) deleteTodo(ctx *gin.Context) {
	var p getTodoParams
	if err := ctx.ShouldBindUri(&p); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	err := t.todoUsecase.Delete(ctx, p.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	resp := deleteTodoResponse{
		ID:      p.ID,
		Message: "Successfully been deleted",
	}

	ctx.JSON(http.StatusOK, resp)
}
