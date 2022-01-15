package models

import (
	"blog/pkg/auth"
	"context"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var user_coll *mongo.Collection

type User struct {
	Username string
	Password []byte
	Email    string
}

func LoadUserdbModel() {
	user_coll = cli.Database("blog").Collection("user")
}

func AddUser(username, password, email string) (string, error) {
	en_aes_pass, err := auth.AesEncrypt([]byte(password))
	if err != nil {
		return "", err
	}
	m := User{username, en_aes_pass, email}

	var temp User
	filter := bson.D{primitive.E{Key: "username", Value: username}}
	err = user_coll.FindOne(context.TODO(), filter).Decode(&temp)
	if err != mongo.ErrNoDocuments {
		return "", errors.New("have already exit")
	}

	result, err := user_coll.InsertOne(context.TODO(), m)
	if err != nil {
		return "", err
	}
	id := fmt.Sprintf("%v", result.InsertedID)
	return id, nil
}

func FindUser(username, password string) (User, error) {
	var result User
	filter := bson.D{primitive.E{Key: "username", Value: username}}
	err := user_coll.FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		log.Fatal("no user")
		return result, err
	}
	if err != nil {
		log.Fatal(err.Error())
		return result, err
	}
	de_aes_pass, err := auth.AesDecrypt([]byte(result.Password))
	if err != nil {
		return result, err
	}
	if string(de_aes_pass) != password {
		return result, errors.New("error password")
	}
	return result, nil
}
