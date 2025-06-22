package main

import (
	"log"
	"net"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/lucianocasa/api/internal/api/graphql"
	"github.com/lucianocasa/api/internal/api/graphql/generated"
	"github.com/lucianocasa/api/internal/api/grpcapi"
	"github.com/lucianocasa/api/internal/api/rest"
	"github.com/lucianocasa/api/internal/order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := order.InitMySQL("user:pass@tcp(db:3306)/ordersdb")
	if err != nil {
		log.Fatalf("DB error: %v", err)
	}

	repo := order.NewRepository(db)
	service := order.NewService(repo)

	// REST
	go func() {
		router := gin.Default()
		rest.RegisterRoutes(router, service)
		log.Println("REST server at :8080")
		router.Run(":8080")
	}()

	// gRPC
	log.Println("Starting gRPC server...")
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}

		grpcServer := grpc.NewServer()
		grpcapi.Register(grpcServer, service)
		reflection.Register(grpcServer)
		log.Fatal(grpcServer.Serve(lis))
	}()

	// GraphQL
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graphql.Resolver{Service: service},
			},
		),
	)
	http.Handle("/graphql", srv)
	http.Handle("/", playground.Handler("GraphQL Playground", "/graphql"))
	log.Println("GraphQL server at :8081")
	http.ListenAndServe(":8081", nil)
}
