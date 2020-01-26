package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	mongoDB := MongoDB{}

	got := mongoDB.filter("test")

	want := bson.D{{
		"skills.name",
		bson.D{{
			"$regex",
			"^test.*$",
		}, {
			"$options",
			"i",
		}},
	}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("filter() got = %v, want %v", got, want)
	}
}
