package server

import (
	"fiber-mongo/internal/core/ports"
	"fiber-mongo/settings"
	"fmt"
	"log"

	fiber "github.com/gofiber/fiber/v2"
)

type Server struct {
	userHandlers ports.IUserHandlers
}

func NewServer(uHandlers ports.IUserHandlers) *Server {
	return &Server{
		userHandlers: uHandlers,
	}
}

func (s *Server) Initialize() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("API is working fine :D")
	})

	api := app.Group("/api")

	users := api.Group("/users")

	users.Post("/login", s.userHandlers.Login)
	users.Post("/register", s.userHandlers.Register)

	serverURL := fmt.Sprintf(":%s", settings.AppConfig.APIPort)

	fmt.Printf("Server running on port %s :D!\n", settings.AppConfig.APIPort)

	err := app.Listen(serverURL)
	if err != nil {
		log.Fatal(err)
	}
}
