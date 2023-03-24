package infrastructure

import (
	"encoding/json"
	"net/http"

	authService "workoutApi/src/infrastructure/helper"

	interfaces "github.com/dev-jpnobrega/workout-domain/src/contract/interfaces"
	values "github.com/dev-jpnobrega/workout-domain/src/contract/valueObject"

	echo "github.com/labstack/echo/v4"
)

func onUnauthorized(context echo.Context, err error) error {
	r := values.ResponseError{}

	r.StatusCode = http.StatusUnauthorized

	return context.JSONPretty(http.StatusUnprocessableEntity, r, " ")
}

func buildAuthParameters(context echo.Context, command interfaces.ICommand) (*values.RequestData, error) {
	model := new(values.RequestData)
	model.Args = command.GetModelValidate().Modal
	model.Authorization = context.Request().Header.Get("Authorization")

	user, err := authService.VerifyAndDecode(model.Authorization)

	if err != nil {
		return nil, err
	}

	jsonBody, err := json.Marshal(user)

	json.Unmarshal(jsonBody, &model.UserInfo)

	return model, err
}

type AuthenticationHandler struct{}

func (h *AuthenticationHandler) Handle(context echo.Context, command interfaces.ICommand) error {
	model, err := buildAuthParameters(context, command)

	if err != nil {
		return onUnauthorized(context, err)
	}

	if err := context.Bind(&model.Args); err != nil {
		return onUnprocessableEntity(context, err)
	}

	if err := context.Validate(model.Args); err != nil {
		return onUnprocessableEntity(context, err)
	}

	rs, errC := command.Execute(*model)

	if errC != nil {
		return onFaliure(context, *errC)
	}

	return onSuccess(context, rs)
}

func NewAuthHandler() IHandler {
	return &AuthenticationHandler{}
}
