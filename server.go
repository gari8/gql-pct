package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gari8/gqlgen-pct/graph"
	"github.com/gari8/gqlgen-pct/loader"
	"github.com/gari8/gqlgen-pct/repository"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strconv"
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

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	server := NewGqlServer(r, &graph.Resolver{
		PlaceRepo:   placeRepo,
		ProgramRepo: programRepo,
	})
	server.Run(8080)
}

type GqlServer struct {
	router *gin.Engine
}

func NewGqlServer(
	router *gin.Engine,
	resolver *graph.Resolver,
) *GqlServer {
	config := graph.Config{Resolvers: resolver}

	router.POST("/graphql", func(c *gin.Context) {
		srv := handler.NewDefaultServer(graph.NewExecutableSchema(config))
		h := loader.Middleware(loader.NewLoaders(resolver.PlaceRepo, resolver.ProgramRepo), srv)
		h.ServeHTTP(c.Writer, c.Request)
	})
	router.GET("/playground", playgroundHandler())

	return &GqlServer{router}
}

func playgroundHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := playground.Handler("GraphQL playground", "/graphql")
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func (c *GqlServer) Run(port int) {
	c.router.Run(":" + strconv.Itoa(port))
}
