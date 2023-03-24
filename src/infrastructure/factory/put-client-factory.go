package infrastructure

import (
	infrastructure "workoutApi/src/db"
	repository "workoutApi/src/infrastructure/repository"

	command "github.com/dev-jpnobrega/workout-domain/src/command"
	interfaces "github.com/dev-jpnobrega/workout-domain/src/contract/interfaces"
)

func PutClientFactory() interfaces.ICommand {
	return &command.PutClientComannd{
		Repository: &repository.ClientRepository{
			Database: &infrastructure.DB{},
		},
	}
}
