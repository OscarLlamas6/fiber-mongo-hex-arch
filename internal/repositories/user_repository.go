package repositories

import (
	"context"
	domain "fiber-mongo/internal/core/domain"
	"fiber-mongo/internal/core/ports"
	"time"

	mongodb "fiber-mongo/internal/infrastructure/database/mongodb"

	"fiber-mongo/settings"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MongoClientTimeout = 5
)

type UserRepository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

var _ ports.IUserRepository = (*UserRepository)(nil)

func NewUserRepository() *UserRepository {

	client, err := mongodb.ConnectDB()

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &UserRepository{
		client:     client,
		database:   client.Database(settings.AppConfig.DBName),
		collection: client.Database(settings.AppConfig.DBName).Collection(settings.AppConfig.DBCollection),
	}
}

func (r *UserRepository) Login(email string, password string) error {
	//Here your code for login in mongo database
	return nil
}

func (r *UserRepository) Register(email string, password string, name string) (*mongo.InsertOneResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newUser := domain.User{
		Id:       primitive.NewObjectID(),
		Name:     name,
		Email:    email,
		Password: password,
	}

	result, err := r.collection.InsertOne(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return result, nil
}
