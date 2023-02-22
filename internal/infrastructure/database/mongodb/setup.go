package configs

import (
	"context"
	"fiber-mongo/settings"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() (*mongo.Client, error) {

	mongoURL := fmt.Sprintf(`mongodb://%v:%v/?authSource=admin&readPreference=primary&directConnection=true&ssl=false`, settings.AppConfig.DBHost, settings.AppConfig.DBPort)
	credential := options.Credential{
		Username: settings.AppConfig.DBUser,
		Password: settings.AppConfig.DBPass,
	}

	ctxMongo, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	clientOptions := options.Client().ApplyURI(mongoURL).SetAuth(credential).SetDirect(true)

	c, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatalf("Error al crear cliente %v", err)
		return nil, err
	}
	err = c.Connect(ctxMongo)
	if err != nil {
		log.Fatalf("Error al realizar conexion %v", err)
		return nil, err
	}

	err = c.Ping(ctxMongo, nil)
	if err != nil {
		log.Fatalf("Error al conectar %v", err)
		return nil, err
	}

	//defer c.Disconnect(ctxMongo)
	return c, nil
}

// getting database collections
func GetCollection(client *mongo.Client, databaseName string, collectionName string) *mongo.Collection {
	collection := client.Database(databaseName).Collection(collectionName)
	return collection
}
