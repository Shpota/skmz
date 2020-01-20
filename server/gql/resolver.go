package gql

import (
	"context"
)

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Programmers(ctx context.Context) ([]*Programmer, error) {
	// A stub. Will be updated later.
	skill := Skill{ID: "sk123", Name: "Java", Importance: 10}
	programmer := Programmer{ID: "12345", Name: "John Doe", Company: "Doe LTD", Skills: []*Skill{&skill}}
	return []*Programmer{&programmer}, nil
}
