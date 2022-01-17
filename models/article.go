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
	ID       string `bson:"_id"`
	Username string
	Text     string
}

type Article struct {
	ID       string `bson:"_id"`
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
	id := fmt.Sprintf("%v", result.InsertedID)
	return id, nil
}

func GetArticle(id string) (Article, error) {
	var result Article
	ob_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}
	err = artice_coll.FindOne(context.TODO(), bson.M{"_id": ob_id}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return result, err
	}
	if err != nil {
		return result, err
	}
	return result, nil
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
func DeleteArticle(id string) error {
	ob_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.D{primitive.E{Key: "_id", Value: ob_id}}
	_, err = artice_coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

// func ChangeArticle() {

// }
// func AddComment() {

// }
// func DeleteComment() {

// }
