package mongodb

import (
	"context"
	"fmt"
	"jit/factory"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// pass:=N2Ud0dzyUK54D7Hq
//username:=git
//string:="mongodb+srv://git:N2Ud0dzyUK54D7Hq@cluster0.xcypdzo.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

type MongoDBConfig struct {
	dbstring string
	client   *mongo.Client
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
	m.client = client // Store the connection!
	logrus.Info("Connected to MongoDB successfully")
	return nil
}

func (m *MongoDBConfig) Create(model *factory.Model) error {
	collection := m.client.Database("git").Collection("test")
	_, err := collection.InsertOne(context.TODO(), model)
	if err != nil {
		return fmt.Errorf("failed to create document: %v", err)
	}
	return nil
}

func (m *MongoDBConfig) GetByRequestID(requestID string) (*factory.Model, error) {
	collection := m.client.Database("git").Collection("test")
	filter := map[string]string{"request_id": requestID}
	var model factory.Model
	err := collection.FindOne(context.TODO(), filter).Decode(&model)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no document found with request_id: %s", requestID)
		}
		return nil, fmt.Errorf("failed to get document by request_id: %v", err)
	}
	return &model, nil
}
