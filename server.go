package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/SRsawaguchi/graphql-hanson-server/graph"
	"github.com/SRsawaguchi/graphql-hanson-server/graph/generated"
	"github.com/SRsawaguchi/graphql-hanson-server/internal/auth"
	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v4"
)

const defaultPort = "8080"

func connectDB(ctx context.Context) (*pgx.Conn, error) {
	url := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	// データベースにとのコネクションを確立
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	resolver := &graph.Resolver{}
	conn, err := connectDB(context.Background())
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	resolver.DB = conn

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	router := chi.NewRouter()
	router.Use(auth.Middleware(conn))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
