package infrastructure

import (
	infrastructure "workoutApi/src/db"
	repository "workoutApi/src/infrastructure/repository"

	command "github.com/dev-jpnobrega/workout-domain/src/command"
	interfaces "github.com/dev-jpnobrega/workout-domain/src/contract/interfaces"
)

func CreateClientFactory() interfaces.ICommand {
	return &command.CreateClientComannd{
		Repository: &repository.ClientRepository{
			Database: &infrastructure.DB{},
		},
	}
}
