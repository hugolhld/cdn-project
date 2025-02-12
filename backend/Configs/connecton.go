package Configs

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbInstance *mongo.Database
var dbOnce sync.Once

func InitDB() *mongo.Database {

	if os.Getenv("GO_ENV") == "test" {
		fmt.Println("🟡 Mode test : MongoDB désactivée")
		return nil
	}

	dbOnce.Do(func() {
		fmt.Println("🔗 URI de connexion à MongoDB :", EnvMongoURI())

		clientOptions := options.Client().ApplyURI(EnvMongoURI())
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			log.Fatal("❌ Erreur de connexion à MongoDB :", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err = client.Ping(ctx, nil)
		if err != nil {
			log.Fatal("❌ Impossible de ping MongoDB :", err)
		}

		fmt.Println("✅ Connexion réussie à MongoDB")
		dbInstance = client.Database("users_db")
	})
	return dbInstance
}

type CollectionInterface interface {
	FindOne(ctx context.Context, filter interface{}) *mongo.SingleResult
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error)
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
}

type MongoCollection struct {
	Collection *mongo.Collection
}

func (m *MongoCollection) FindOne(ctx context.Context, filter interface{}) *mongo.SingleResult {
	return m.Collection.FindOne(ctx, filter)
}

func (m *MongoCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	return m.Collection.Find(ctx, filter, opts...)
}

func (m *MongoCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return m.Collection.InsertOne(ctx, document, opts...)
}

func (m *MongoCollection) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return m.Collection.DeleteOne(ctx, filter, opts...)
}

func (m *MongoCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return m.Collection.UpdateOne(ctx, filter, update, opts...)
}

func GetCollection(collectionName string) CollectionInterface {
	if os.Getenv("GO_ENV") == "test" {
		fmt.Println("🟡 Mode test : MongoDB désactivée")
		return nil
	}

	return &MongoCollection{Collection: InitDB().Collection(collectionName)}
}

var MemberCollection CollectionInterface = GetCollection("users")
