package handlers

import (
	"fiber-mongo/internal/core/domain"
	"fiber-mongo/internal/core/ports"
	responses "fiber-mongo/internal/server/responses"
	"net/http"

	validator "github.com/go-playground/validator/v10"
	fiber "github.com/gofiber/fiber/v2"
)

type UserHandlers struct {
	userService ports.IUserService
}

var validate = validator.New()
var _ ports.IUserHandlers = (*UserHandlers)(nil)

func NewUserHandlers(userService ports.IUserService) *UserHandlers {
	return &UserHandlers{
		userService: userService,
	}
}

func (h *UserHandlers) Login(c *fiber.Ctx) error {
	var email string
	var password string
	//Extract the body and get the email and password
	err := h.userService.Login(email, password)
	if err != nil {
		return err
	}
	return nil
}

func (h *UserHandlers) Register(c *fiber.Ctx) error {

	var user domain.User

	//validate the request body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	//Extract the body and get the email and password
	result, err := h.userService.Register(user.Email, user.Password, user.Name)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}
