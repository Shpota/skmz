package db

import (
	"context"
	"github.com/shpota/skmz/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type DB struct {
	collection *mongo.Collection
}

func New(client *mongo.Client) *DB {
	programmers := client.Database("programmers").Collection("programmers")
	return &DB{
		collection: programmers,
	}
}

func (db DB) GetProgrammers() ([]*model.Programmer, error) {
	res, err := db.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Printf("Error while fetching programmers: %s", err.Error())
		return nil, err
	}
	var p []*model.Programmer
	err = res.All(context.TODO(), &p)
	if err != nil {
		log.Printf("Error while decoding programmers: %s", err.Error())
		return nil, err
	}
	return p, nil
}
