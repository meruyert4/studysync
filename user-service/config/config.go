package config

import (
	"time"
	"user-service/pkg/minio"
	"user-service/pkg/postgres"

	env "github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	Postgres postgres.Config
	Minio    minio.Config
	Server   Server

	Version string `env:"VERSION"`
}

type Server struct {
	GRPCServer GRPCServer
}

type GRPCServer struct {
	Port                  int           `env:"GRPC_PORT,notEmpty"`
	MaxRecvMsgSizeMiB     int           `env:"GRPC_MAX_MESSAGE_SIZE_MIB" envDefault:"12"`
	MaxConnectionAge      time.Duration `env:"GRPC_MAX_CONNECTION_AGE" envDefault:"30s"`
	MaxConnectionAgeGrace time.Duration `env:"GRPC_MAX_CONNECTION_AGE_GRACE" envDefault:"10s"`
}

func New() (*Config, error) {
	var cfg Config

	err := godotenv.Load()
	if err != nil {
		return &cfg, err
	}

	err = env.Parse(&cfg)

	return &cfg, err
}
