package models

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var cli *mongo.Client

func Init() {
	var err error
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("you need set MONGODB_URI in .env")
	}
	cli, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("connect mogodb error")
	}

}

func Close() {
	if err := cli.Disconnect(context.TODO()); err != nil {
		log.Fatal("disconnect mogodb error")
	}
}
