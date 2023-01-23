package main

import (
	"github.com/gari8/gqlgen-pct/loader"
	"github.com/gari8/gqlgen-pct/repository"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gari8/gqlgen-pct/graph"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		dsn = "test:test@tcp(mysql_db)/test?charset=utf8mb4&parseTime=True&loc=Local"
	}

	db, err := NewDatabase(dsn)
	if err != nil {
		log.Fatalln(err)
	}

	placeRepo := repository.NewPlaceRepository(db)
	programRepo := repository.NewProgramRepository(db)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		PlaceRepo:   placeRepo,
		ProgramRepo: programRepo,
	}}))

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	dl := loader.DataLoaderMiddleware(placeRepo, programRepo)

	r.Get("/", playground.Handler("GraphQL playground", "/query"))
	r.With(dl).Post("/query", srv.ServeHTTP)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
