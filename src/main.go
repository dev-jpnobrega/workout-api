package main

import (
	DB "workoutApi/src/db"
	helper "workoutApi/src/infrastructure/helper"
	router "workoutApi/src/infrastructure/router"
)

func main() {
	server := helper.CreateHTTPServer()

	database := &DB.DB{}
	database.Connect("postgres", "postgres://postgres:postgres@localhost/workout?sslmode=disable")

	router.Build(server)

	server.Logger.Fatal(server.Start(":8082"))
}
