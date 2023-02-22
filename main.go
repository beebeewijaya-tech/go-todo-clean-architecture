package main

import (
	"fmt"
	"log"

	"go-clean-arch-todo.com/util"

	_deliveryHttp "go-clean-arch-todo.com/delivery/http"
	_deliveryTodo "go-clean-arch-todo.com/delivery/http/todo"
	_pgRepo "go-clean-arch-todo.com/repository/postgres"
	_todoUsecase "go-clean-arch-todo.com/usecase/todo"
)

func main() {
	c, err := util.NewConfig("./env")
	if err != nil {
		log.Fatalf("failed when initialize config: %v", err)
	}

	p := _pgRepo.Postgres{
		HOST:     c.GetString("DB.HOST"),
		USERNAME: c.GetString("DB.USERNAME"),
		PASSWORD: c.GetString("DB.PASSWORD"),
		PORT:     c.GetInt("DB.PORT"),
		DATABASE: c.GetString("DB.DATABASE"),
		DIALECT:  c.GetString("DB.DIALECT"),
	}
	db, err := p.Connect()
	if err != nil {
		log.Fatalf("failed when initialize database: %v", err)
	}

	// Repository Initalization
	todoRepo := _pgRepo.NewPostgresTodoRepository(db)

	// Usecase Initalization
	todoUsecase := _todoUsecase.NewTodoUsecase(todoRepo)

	s := _deliveryHttp.NewServer(c)

	// Registering handler
	_deliveryTodo.NewTodoHandler(s.Router, todoUsecase)

	address := fmt.Sprintf("%s:%d", c.GetString("SERVER.ADDRESS"), c.GetInt("SERVER.PORT"))
	err = s.Start(address)
	if err != nil {
		log.Fatalf("error when starting server: %v", err)
	}
}
