package infrastructure

import (
	interfaces "github.com/dev-jpnobrega/workout-domain/src/contract/interfaces"

	echo "github.com/labstack/echo/v4"
)

// A IHandler represent hanlder interface
type IHandler interface {
	Handle(context echo.Context, command interfaces.ICommand) error
}
