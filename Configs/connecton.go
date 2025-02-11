package Configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB connection function
// func ConnectDB() *mongo.Client {
// 	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	err = client.Connect(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	//ping the database
// 	err = client.Ping(ctx, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("MongoDB Connected successfully !!!! ")
// 	return client
// }

func ConnectDB() *mongo.Database {
	mongoURI := "mongodb://admin:password@0.0.0.0:27017/"

	// Options de connexion
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("❌ Erreur de connexion à MongoDB :", err)
	}

	// Vérifier la connexion
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("❌ Impossible de ping MongoDB :", err)
	}

	fmt.Println("✅ Connexion réussie à MongoDB")
	return client.Database("users_db")
}

// // MongoDB Client instance
var DB *mongo.Database = ConnectDB()

// Getting database collection
func GetCollection(client *mongo.Database, collectionName string) *mongo.Collection {
	collection := client.Collection(collectionName)
	return collection
}
