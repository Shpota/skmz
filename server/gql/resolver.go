package gql

import (
	"context"
	"github.com/shpota/skmz/db"
	"github.com/shpota/skmz/model"
)

type Resolver struct {
	DB *db.DB
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Programmers(ctx context.Context) ([]*model.Programmer, error) {
	return r.DB.GetProgrammers()
}
