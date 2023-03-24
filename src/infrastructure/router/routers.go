package infrastructure

import (
	"net/http"

	factory "workoutApi/src/infrastructure/factory"
	handler "workoutApi/src/infrastructure/http"

	interfaces "github.com/dev-jpnobrega/workout-domain/src/contract/interfaces"

	"github.com/labstack/echo/v4"
)

func adapter(command interfaces.ICommand, h handler.IHandler) func(c echo.Context) error {
	return func(c echo.Context) error {
		return h.Handle(c, command)
	}
}

func Build(server *echo.Echo) {
	server.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "OI")
	})

	server.GET("/v1/client", adapter(
		factory.GetClientFactory(),
		handler.NewHandler(),
	))

	server.POST("/v1/client", adapter(
		factory.CreateClientFactory(),
		handler.NewAuthHandler(),
	))

	server.PUT("/v1/client/:id", adapter(
		factory.PutClientFactory(),
		handler.NewHandler(),
	))
}
