//go:build wireinject
// +build wireinject

package di

import (
	"Test/Handlers"
	"Test/repositories"
	"Test/usecase"

	"github.com/google/wire"
)

func InitializeHandlers() *Handlers.TestHandlers {
	wire.Build(
		Handlers.NewTestHandlers,
		usecase.NewUecase,
		repositories.NewRepo,
	)
	return &Handlers.TestHandlers{}
}
