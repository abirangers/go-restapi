// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/abirangers/go-restapi/app"
	"github.com/abirangers/go-restapi/controller"
	"github.com/abirangers/go-restapi/middleware"
	"github.com/abirangers/go-restapi/repository"
	"github.com/abirangers/go-restapi/service"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"net/http"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from injector.go:

func InitializedServer() *http.Server {
	categoryRepository := repository.NewCategoryRepository()
	db := app.NewDB()
	v := NewValidatorOptions()
	validate := validator.New(v...)
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)
	authMiddleware := middleware.NewAuthMiddleware(router)
	server := NewServer(authMiddleware)
	return server
}

// injector.go:

var categorySet = wire.NewSet(repository.NewCategoryRepository, service.NewCategoryService, controller.NewCategoryController)

// NewValidatorOptions adalah penyedia kustom untuk opsi validator
func NewValidatorOptions() []validator.Option {

	return nil
}