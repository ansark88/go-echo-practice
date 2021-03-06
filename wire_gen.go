// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"go-echo-practice/infra/persistence"
	"go-echo-practice/interfaces/handler"
	"go-echo-practice/usecase"
)

// Injectors from wire.go:

func profileInitialize() handler.ProfileHandler {
	profileRepository := persistence.NewProfilePersistence()
	profileUseCase := usecase.NewProfileUseCase(profileRepository)
	profileHandler := handler.NewProfileHandler(profileUseCase)
	return profileHandler
}
