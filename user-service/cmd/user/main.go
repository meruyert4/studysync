package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"user-service/config"
	"user-service/pkg/minio"
	"user-service/pkg/postgres"
)

func main() {
	// Step 1: Graceful context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Step 2: Handle interrupt
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		log.Println("Graceful shutdown initiated...")
		cancel()
	}()

	// Step 3: Load config
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Step 4: Connect PostgreSQL
	db, err := postgres.NewPostgres(cfg.Postgres)
	if err != nil {
		log.Fatalf("failed to connect to Postgres: %v", err)
	}
	defer db.Close()

	// Step 5: Connect MinIO
	minioClient, err := minio.NewMinio(cfg.Minio)
	if err != nil {
		println(minioClient)
		log.Fatalf("failed to connect to MinIO: %v", err)
	}

	<-ctx.Done()
	log.Println("Shutdown complete.")
}
