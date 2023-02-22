# Go-Todo <> Clean Architecture

Welcome to the Go Todo implemented using Clean Architecture methodologies!

## What is clean architecture?

It's an architecture that has 5 principles

1. Independent of frameworks
2. Testable without much dependencies
3. Independent of UI
4. Independent of Database
5. Independent of any external agency

## Layers

1. Domain -> Purpose for store any struct and interface that being consumed
2. Repository -> Purpose for handling CRUD to the database only, no business logic applies here
3. Usecase -> Purpose for handling every business logic, like calculation
4. Delivery -> Decide how the data would be presented or how will we interact to the server, could be a REST or GRPC

## Installation

```bash
git clone https://github.com/beebeewijaya-tech/go-todo

cp Makefile.sample Makefile

cp env/sample.config.json env/config.json

make migrate-up

go mod tidy
```

Remember change the Makefile scripts to insert your PostgreSQL credentials there.

Also remember to add the config.json value to the env as well

## Folder Structure

```bash
├───delivery
│   └───http
│       └───todo
├───domain
│   └───mock_domain
├───env
├───migrations
├───pkg
├───repository
│   └───postgres
├───usecase
│   └───todo
└───util
```

`delivery` will represent your route handler or any grpc handler

`domain` will represent the whole struct and interface that being consumed

`env` your environment variables as a JSON

`migrations` migration file to be migrate

`pkg` some external package that you want to implement, such as JWT

`repository` CRUD to the database lives here

`usecase` Business logic / flow lives here

`util` Utils tools to support business logic
