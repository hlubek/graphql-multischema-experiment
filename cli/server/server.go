package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	backend_graph "github.com/hlubek/graphql-multischema-experiment/backend/graph"
	backend_generated "github.com/hlubek/graphql-multischema-experiment/backend/graph/generated"
	frontend_graph "github.com/hlubek/graphql-multischema-experiment/frontend/graph"
	frontend_generated "github.com/hlubek/graphql-multischema-experiment/frontend/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	frontendSrv := handler.NewDefaultServer(frontend_generated.NewExecutableSchema(frontend_generated.Config{Resolvers: &frontend_graph.Resolver{}}))
	backendSrv := handler.NewDefaultServer(backend_generated.NewExecutableSchema(backend_generated.Config{Resolvers: &backend_graph.Resolver{}}))

	http.Handle("/frontend", playground.Handler("GraphQL playground", "/frontend/query"))
	http.Handle("/frontend/query", frontendSrv)
	http.Handle("/backend", playground.Handler("GraphQL playground", "/backend/query"))
	http.Handle("/backend/query", backendSrv)

	log.Printf("connect to http://localhost:%s/frontend or http://localhost:%s/backend for GraphQL playground", port, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
