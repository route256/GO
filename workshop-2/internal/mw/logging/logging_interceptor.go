package logging

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func Interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	raw, _ := protojson.Marshal((req).(proto.Message)) // для превращения protbuf структур в json используем google.golang.org/protobuf/encoding/protojson пакет а не encoding/json
	log.Printf("method: %v, req: %v\n", info.FullMethod, string(raw))

	resp, err = handler(ctx, req)
	if resp != nil {
		rawResp, _ := protojson.Marshal((resp).(proto.Message))
		log.Printf("method: %v, req: %v\n", info.FullMethod, string(rawResp))
	}

	log.Printf("resp: %v, err: %v\n", resp, err)
	return resp, err
}
