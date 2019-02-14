package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/readpref"
)

const (
	host       = "localhost"
	port       = "27017"
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

func main() {

	logger := log.New(os.Stdout, "[app]", (log.Ldate | log.Ltime | log.Lshortfile))
	
	
	// connect to mongo
	// session := initMongo()
	// mongoSession.session = session

	// close the mongo connection when application ends
	// defer shutDownMongo(mongoSession.session)

}

func initServer(addr string) {
	log.Println("server is starting.")
}

func initMongo() (session *mongo.Client) {
	session, err := mongo.Connect(context.TODO(), getConnectionString())

	if err != nil {
		log.Fatal(err)
	}

	err = session.Ping(context.TODO(), readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	log.Println("mongo session created successfully.")

	return session
}

func shutDownMongo(session *mongo.Client) {

	err := session.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	log.Println("mongo connection closed.")

}

func getConnectionString() string {

	var sb strings.Builder

	sb.WriteString("mongodb://")
	sb.WriteString(host)
	sb.WriteString(":")
	sb.WriteString(port)

	return sb.String()

}
