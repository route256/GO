package config

import (
	"os"
)

type key string

const (
	NoteGRPCHost                  = key("NOTE_GRPC_HOST")
	NoteHTTPHost                  = key("NOTE_HTTP_HOST")
	PostgresDSN                   = key("POSTGRES_DSN")
	PostgresMaxConnections        = key("POSTGRES_MAX_CONNECTIONS")
	PostgresMaxConnectionIdleTime = key("POSTGRES_MAX_CONNECTION_IDLE_TIME")
)

func GetValue(k key) string {
	return os.Getenv(string(k))
}
