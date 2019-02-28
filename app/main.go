package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/mohamedgamaleldin/go-alpha/app/server"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	serverHost = ""
	serverPort = "8080"
	mongoHost  = "mongo"
	mongoPort  = "27017"
	database   = "db"
	username   = ""
	password   = ""
	collection = "test"
)

// Item is a dummy struct for testing
type Item struct {
	Name string
	Age  int
}

// MongoSession is a struct to store the instance of a mongo client connection
type MongoSession struct {
	session *mongo.Client
}

var mongoSession = MongoSession{}
var logger *log.Logger

func main() {

	logger = log.New(os.Stdout, "[app] ", (log.Ldate | log.Ltime | log.Lshortfile))

	/* 	// connect to mongo
	   	session := initMongo()
	   	mongoSession.session = session

	   	// add document to test collection
	   	_, err := session.Database(database).Collection(collection).InsertOne(context.TODO(), Item{"Alex", 30})

	   	if err != nil {
	   		logger.Fatal(err)
	   	}

	   	// close the mongo connection when application ends
	   	defer shutDownMongo(mongoSession.session) */

	server.InitServer(logger, getServerURI())

}

func initMongo() (session *mongo.Client) {
	session, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(getMongoURI()))

	if err != nil {
		logger.Fatal(err)
	}

	err = session.Ping(context.TODO(), readpref.Primary())

	if err != nil {
		logger.Fatal(err)
	}

	logger.Println("mongo session created successfully.")

	return session
}

func shutDownMongo(session *mongo.Client) {

	err := session.Disconnect(context.TODO())

	if err != nil {
		logger.Fatal(err)
	}

	logger.Println("mongo connection closed.")

}

func getMongoURI() string {

	var sb strings.Builder

	sb.WriteString("mongodb://")
	sb.WriteString(mongoHost)
	sb.WriteString(":")
	sb.WriteString(mongoPort)

	return sb.String()

}

func getServerURI() string {

	var sb strings.Builder

	sb.WriteString(serverHost)
	sb.WriteString(":")
	sb.WriteString(serverPort)

	return sb.String()
}
