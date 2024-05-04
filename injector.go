//go:build wireinject
// +build wireinject

package main

import (
	"github.com/abirangers/go-restapi/app"
	"github.com/abirangers/go-restapi/controller"
	"github.com/abirangers/go-restapi/middleware"
	"github.com/abirangers/go-restapi/repository"
	"github.com/abirangers/go-restapi/service"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	service.NewCategoryService,
	controller.NewCategoryController,
)

// NewValidatorOptions adalah penyedia kustom untuk opsi validator
func NewValidatorOptions() []validator.Option {
	// Anda dapat menambahkan opsi validator yang diperlukan di sini
	return nil
}

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		NewValidatorOptions,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}
