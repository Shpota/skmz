package gql

import (
	"context"
	"github.com/shpota/skmz/db"
	"github.com/shpota/skmz/gql/gen"
	"github.com/shpota/skmz/model"
)

type Resolver struct {
	DB db.DB
}

func (r *Resolver) Query() gen.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Programmers(ctx context.Context, skill string) ([]*model.Programmer, error) {
	return r.DB.GetProgrammers(skill)
}
