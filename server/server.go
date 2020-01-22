package main

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/shpota/skmz/cors"
	"github.com/shpota/skmz/gql"
	"log"
	"net/http"
	"os"
)

func main() {
	hndlr := handler.GraphQL(gql.NewExecutableSchema(gql.Config{
		Resolvers: &gql.Resolver{},
	}))
	static := "/webapp"

	profile := os.Getenv("profile")
	if profile != "prod" {
		hndlr = cors.Disable(hndlr)
		static = "../webapp/build"
		http.Handle("/playground",
			handler.Playground("GraphQL playground", "/query"),
		)
	}

	http.Handle("/query", hndlr)
	http.Handle("/", http.FileServer(http.Dir(static)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
