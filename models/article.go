package models

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var artice_coll *mongo.Collection

type Comment struct {
	Username string
	Text     string
}

type Article struct {
	Username string
	Name     string
	Text     string
	Tag      string
	Comments []Comment
}

func LoadArticledbModel() {
	artice_coll = cli.Database("blog").Collection("article")
}

func AddAticle(name, username, text, tag string) (string, error) {
	m := Article{Name: name, Username: username, Text: text, Tag: tag}
	result, err := artice_coll.InsertOne(context.TODO(), m)
	if err != nil {
		return "", err
	}
	id := fmt.Sprintf("%v", result)
	return id, nil
}
func GetAllArticle() ([]Article, error) {
	var articles []Article
	cursor, err := artice_coll.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var elem Article
		if err := cursor.Decode(&elem); err != nil {
			return nil, err
		}

		articles = append(articles, elem)
	}
	return articles, nil

}
func DeleteArticle(name string) error {
	filter := bson.D{primitive.E{Key: "name", Value: name}}
	_, err := artice_coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

// func FindAticles(name string) {

// }

// func ChangeArticle() {

// }
// func AddComment() {

// }
// func DeleteComment() {

// }
