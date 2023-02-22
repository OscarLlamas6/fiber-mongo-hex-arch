package main

import (
	"fiber-mongo/internal/core/services"
	"fiber-mongo/internal/handlers"
	"fiber-mongo/internal/repositories"
	"fiber-mongo/internal/server"
	"fiber-mongo/settings"
)

func main() {

	settings.SetConfig()

	//repositories
	userRepository := repositories.NewUserRepository()
	//services
	userService := services.NewUserService(userRepository)
	//handlers
	userHandlers := handlers.NewUserHandlers(userService)
	//server
	myServer := server.NewServer(userHandlers)
	myServer.Initialize()
}
