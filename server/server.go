package main

import (
	"context"
	"github.com/99designs/gqlgen/handler"
	"github.com/shpota/skmz/cors"
	"github.com/shpota/skmz/db"
	"github.com/shpota/skmz/gql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
)

func main() {
	profile := os.Getenv("profile")
	opts := options.Client().ApplyURI(
		"mongodb://" + dbHost(profile) + ":27017",
	)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())
	resolver := &gql.Resolver{
		DB: db.New(client),
	}
	hndlr := handler.GraphQL(gql.NewExecutableSchema(gql.Config{
		Resolvers: resolver,
	}))

	if profile != "prod" {
		hndlr = cors.Disable(hndlr)
		http.Handle("/playground",
			handler.Playground("GraphQL playground", "/query"),
		)
	}
	http.Handle("/query", hndlr)
	http.Handle("/", http.FileServer(http.Dir("/webapp")))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func dbHost(profile string) string {
	if profile != "prod" {
		return "localhost"
	}
	return "db"
}
