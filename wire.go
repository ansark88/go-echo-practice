//+build wireinject

package main

import (
	"go-echo-practice/infra/persistence"
	"go-echo-practice/interfaces/handler"
	"go-echo-practice/usecase"

	"github.com/google/wire"
)

func profileInitialize() handler.ProfileHandler {
	wire.Build(handler.NewProfileHandler, usecase.NewProfileUseCase, persistence.NewProfilePersistence)
	return nil
}
