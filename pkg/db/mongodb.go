package mongodb

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// pass:=N2Ud0dzyUK54D7Hq
//username:=git
//string:="mongodb+srv://git:N2Ud0dzyUK54D7Hq@cluster0.xcypdzo.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

type MongoDBConfig struct {
	dbstring string
}

func NewRepository(dbstring string) *MongoDBConfig {
	return &MongoDBConfig{
		dbstring: dbstring,
	}
}

func (m *MongoDBConfig) Connect() error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(m.dbstring))
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.TODO())
	logrus.Info("Connected to MongoDB successfully")
	return nil
}
