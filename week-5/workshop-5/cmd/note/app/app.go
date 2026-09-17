package app

import (
	"context"
	"golang.org/x/sync/errgroup"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"route256/ws5/internal/app"
	"route256/ws5/internal/client/postgres"
	"route256/ws5/internal/config"
	"route256/ws5/internal/core/note"
	notestorage "route256/ws5/internal/storage/note"
	pb "route256/ws5/pkg"
)

type App struct {
	impl *app.Implementation
}

func New(ctx context.Context) (*App, error) {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	maxConns, err := strconv.ParseInt(config.GetValue(config.PostgresMaxConnections), 10, 32)
	if err != nil {
		return nil, err
	}

	maxConnIdleTime, err := time.ParseDuration(config.GetValue(config.PostgresMaxConnectionIdleTime))
	if err != nil {
		return nil, err
	}

	db, err := postgres.New(ctx, postgres.Config{
		DSN:                   config.GetValue(config.PostgresDSN),
		MaxConnections:        int32(maxConns),
		MaxConnectionIdleTime: maxConnIdleTime,
	})
	if err != nil {
		return nil, err
	}

	noteStorage := notestorage.New(db)

	noteCore := note.New(noteStorage)

	return &App{
		impl: app.New(noteCore),
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	grpcHost := config.GetValue(config.NoteGRPCHost)

	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		mux := runtime.NewServeMux()
		if err := pb.RegisterNoteHandlerFromEndpoint(gCtx, mux, grpcHost, []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		}); err != nil {
			return errors.Wrap(err, "cannot register http server")
		}

		httpHost := config.GetValue(config.NoteHTTPHost)
		log.Println("HTTP server started on: ", httpHost)

		if err := http.ListenAndServe(httpHost, mux); err != nil {
			return err
		}

		return nil
	})

	g.Go(func() error {
		server := grpc.NewServer()
		pb.RegisterNoteServer(server, a.impl)

		list, err := net.Listen("tcp", grpcHost)
		if err != nil {
			return err
		}

		log.Println("GRPC server started on: ", grpcHost)

		if err = server.Serve(list); err != nil {
			log.Fatalf("failed to process gRPC server: %s", err.Error())
		}

		return nil
	})

	return g.Wait()
}
