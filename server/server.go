package main

import (
	"context"
	"github.com/99designs/gqlgen/handler"
	"github.com/shpota/skmz/cors"
	"github.com/shpota/skmz/db"
	"github.com/shpota/skmz/gql"
	"github.com/shpota/skmz/gql/gen"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
)

func main() {
	client, err := mongo.Connect(context.TODO(), clientOptions())
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())
	http.Handle("/query", gqlHandler(db.New(client)))
	http.Handle("/playground",
		handler.Playground("GraphQL playground", "/query"),
	)
	http.Handle("/", http.FileServer(http.Dir("/webapp")))
	err = http.ListenAndServe(":8080", nil)
	log.Println(err)
}

func gqlHandler(db db.DB) http.HandlerFunc {
	config := gen.Config{
		Resolvers: &gql.Resolver{DB: db},
	}
	gh := handler.GraphQL(gen.NewExecutableSchema(config))
	if os.Getenv("profile") != "prod" {
		gh = cors.Disable(gh)
	}
	return gh
}

func clientOptions() *options.ClientOptions {
	host := "db"
	if os.Getenv("profile") != "prod" {
		host = "localhost"
	}
	return options.Client().ApplyURI(
		"mongodb://" + host + ":27017",
	)
}
