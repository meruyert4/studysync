package minio

import (
	"fmt"

	"github.com/minio/minio-go"
)

type Config struct {
	Endpoint        string `env:"MINIO_ENDPOINT,notEmpty"`
	AccessKeyID     string `env:"MINIO_ACCESS_KEY,notEmpty"`
	SecretAccessKey string `env:"MINIO_SECRET_KEY,notEmpty"`
	UseSSL          bool   `env:"MINIO_USE_SSL" envDefault:"false"`
}

func NewMinio(cfg Config) (*minio.Client, error) {
	client, err := minio.New(cfg.Endpoint, cfg.AccessKeyID, cfg.SecretAccessKey, cfg.UseSSL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to minio: %w", err)
	}

	return client, nil
}
