package auth

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func Interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Print("No metadata")
	}
	values := md.Get("x-authirization")
	if len(values) == 0 {
		log.Print("No x-authirization header")
	}

	// logic

	md.Append("x-header", "foo")

	s := struct{}{}
	ctx = context.WithValue(ctx, "foo", s)

	metadata.NewOutgoingContext(ctx, md)

	resp, err = handler(ctx, req)

	return resp, err
}
