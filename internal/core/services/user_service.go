package services

import (
	"fiber-mongo/internal/core/ports"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	userRepository ports.IUserRepository
}

// This line is for get feedback in case we are not implementing the interface correctly
var _ ports.IUserService = (*UserService)(nil)

func NewUserService(repository ports.IUserRepository) *UserService {
	return &UserService{
		userRepository: repository,
	}
}

func (s *UserService) Login(email string, password string) error {
	err := s.userRepository.Login(email, password)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) Register(email string, password string, name string) (*mongo.InsertOneResult, error) {

	result, err := s.userRepository.Register(email, password, name)
	if err != nil {
		return nil, err
	}
	return result, nil
}
