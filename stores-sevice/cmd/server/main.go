package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"stores/internal/db"
	"stores/internal/graph"
)

func main() {
	postgres := db.NewPostgres()
	defer postgres.Close()

	resolver := &graph.Resolver{
		DB: postgres,
	}

	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{Resolvers: resolver},
		),
	)

	http.Handle("/query", srv)
	http.Handle("/", playground.Handler("Stores GraphQL", "/query"))

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
