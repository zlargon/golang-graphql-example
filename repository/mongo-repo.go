package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/zlargon/gograph/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type VideoRepository interface {
	Save(video *model.Video)
	FindAll() []*model.Video
}

const (
	DATABASE   = "graphql"
	COLLECTION = "videos"
)

type database struct {
	client *mongo.Client
}

func New() VideoRepository {
	// mongodb+srv://USERNAME:PASSWORD@HOST:PORT
	MONGODB := os.Getenv("MONGO")
	clientOptions := options.Client().ApplyURI(MONGODB)
	clientOptions = clientOptions.SetMaxPoolSize(50)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	dbClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	defer fmt.Println("Connect to MongoDB!")

	return &database{
		client: dbClient,
	}
}

func (db *database) Save(video *model.Video) {
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	_, err := collection.InsertOne(context.TODO(), video)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *database) FindAll() []*model.Video {
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	var result []*model.Video
	for cursor.Next(context.TODO()) {
		var v *model.Video
		err := cursor.Decode(&v)
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, v)
	}

	return result
}
