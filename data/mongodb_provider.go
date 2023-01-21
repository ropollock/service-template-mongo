package data

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoDBProvider struct {
	mongoContext    context.Context
	mongoClient     *mongo.Client
	mainDB          *mongo.Database
	usersCollection *mongo.Collection
}

type MongoDBProviderInterface interface {
	GetContext() context.Context
	GetClient() *mongo.Client
	GetDB() *mongo.Database
	GetUsersCollection() *mongo.Collection
	Connect(dbURI string)
}

func MongoDBProvider() *mongoDBProvider {
	return &mongoDBProvider{}
}

func (provider *mongoDBProvider) GetContext() context.Context {
	return provider.mongoContext
}

func (provider *mongoDBProvider) GetClient() *mongo.Client {
	return provider.mongoClient
}

func (provider *mongoDBProvider) GetDB() *mongo.Database {
	return provider.mainDB
}

func (provider *mongoDBProvider) GetUsersCollection() *mongo.Collection {
	return provider.usersCollection
}

func (provider *mongoDBProvider) Connect(dbURI string) {
	provider.mongoContext = context.TODO()
	mongoconn := options.Client().ApplyURI(dbURI)
	var dbConnErr error
	provider.mongoClient, dbConnErr = mongo.Connect(provider.mongoContext, mongoconn)

	if dbConnErr != nil {
		panic(dbConnErr)
	}

	if err := provider.mongoClient.Ping(provider.mongoContext, readpref.Primary()); err != nil {
		panic(err)
	}

	provider.mainDB = provider.mongoClient.Database("main")
	provider.usersCollection = provider.mainDB.Collection("users")

	fmt.Println("MongoDB successfully connected.")
}

func StringToObjectID(id string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(id)
}
