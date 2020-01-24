package db

import (
	"context"
	"github.com/shpota/skmz/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type DB interface {
	GetProgrammers(skill string) ([]*model.Programmer, error)
}

type MongoDB struct {
	collection *mongo.Collection
}

func New(client *mongo.Client) *MongoDB {
	programmers := client.Database("programmers").Collection("programmers")
	return &MongoDB{
		collection: programmers,
	}
}

func (db MongoDB) GetProgrammers(skill string) ([]*model.Programmer, error) {
	res, err := db.collection.Find(context.TODO(), db.filter(skill))
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

func (db MongoDB) filter(skill string) bson.D {
	return bson.D{{
		"skills.name",
		bson.D{{
			"$regex",
			"^" + skill + ".*$",
		}, {
			"$options",
			"i",
		}},
	}}
}
