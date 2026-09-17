package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	api "gitlab.ozon.dev/go/classroom-8/students/workshop-2/internal/api/notes"
	"gitlab.ozon.dev/go/classroom-8/students/workshop-2/internal/mw/logging"
	"gitlab.ozon.dev/go/classroom-8/students/workshop-2/internal/mw/panic"
	"gitlab.ozon.dev/go/classroom-8/students/workshop-2/internal/service/notes"
	desc "gitlab.ozon.dev/go/classroom-8/students/workshop-2/pkg/api/notes/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

const (
	grpcPort = 50051
	httpPort = 8081
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort)) // :82
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer( // grpc сервер (aka http.Serever)
		grpc.ChainUnaryInterceptor( // Unary интерсепторы (aka middleware)
			panic.Interceptor,
			logging.Interceptor,
		),
		grpc.ChainStreamInterceptor( // Stream интерсепторы (aka middleware)
		// func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {

		// },
		),
	)

	reflection.Register(grpcServer) // Рефлексия! (Повзоляет получать описание rpc функционала нашего сервиса. Полезно для Postman)

	notesUsecase := notes.NewService() // usecase level

	controller := api.NewNotesServer(notesUsecase) // delivery/controller level

	desc.RegisterNotesServer(grpcServer, controller) // Вешаем наш обработчик (controller) на grpc сервер

	log.Printf("server listening at %v", lis.Addr())

	go func() {
		if err = grpcServer.Serve(lis); err != nil { // запускаем grpc сервер
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Создаем коннект с grpc сервером
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		lis.Addr().String(),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	mux := runtime.NewServeMux() // создаем мультиплекор/роутер (mux == router)

	err = desc.RegisterNotesHandler(context.Background(), mux, conn) // Вешаем обработчик запросов
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{ // Создаем HTTP gateway сервер
		Addr:    fmt.Sprintf(":%d", httpPort),           // :80
		Handler: logging.WithHTTPLoggingMiddleware(mux), // middleware
	}

	log.Printf("Serving gRPC-Gateway on %s\n", gwServer.Addr) // запускаем HTTP сервер
	log.Fatalln(gwServer.ListenAndServe())
}
