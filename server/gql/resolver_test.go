package gql

import (
	"context"
	"errors"
	"github.com/shpota/skmz/model"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

type MockDB struct {
	collection *mongo.Collection
}

func (mockDB MockDB) GetProgrammers(string) ([]*model.Programmer, error) {
	return []*model.Programmer{{ID: "test-id"}}, errors.New("test-error")
}

func TestProgrammers(t *testing.T) {
	r := &queryResolver{
		Resolver: &Resolver{&MockDB{}},
	}

	programmers, err := r.Programmers(context.TODO(), "test")

	if programmers[0].ID != "test-id" {
		t.Errorf("GetProgrammers() got = %v, want test-id", programmers[0].ID)
	}
	if err.Error() != "test-error" {
		t.Errorf("GetProgrammers() got = %v, want test-error", err.Error())
	}
}
