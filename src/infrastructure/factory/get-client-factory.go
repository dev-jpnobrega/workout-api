package infrastructure

import (
	infrastructure "workoutApi/src/db"
	repository "workoutApi/src/infrastructure/repository"

	command "github.com/dev-jpnobrega/workout-domain/src/command"
	interfaces "github.com/dev-jpnobrega/workout-domain/src/contract/interfaces"
)

func GetClientFactory() interfaces.ICommand {
	return &command.GetClientCommand{
		Repository: &repository.ClientRepository{
			Database: &infrastructure.DB{},
		},
	}
}
