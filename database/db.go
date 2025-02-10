package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

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
	db = client.Database("mydatabase")
	return db
}

func CreateUser()
