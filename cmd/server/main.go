package main

import (
	"fmt"
	"log"
	"os"

	"test_qa-api/internal/server"
	"test_qa-api/pkg/db"
)

func main() {
	dsn := db.DSNFromEnv()
	gdb, err := db.Connect(dsn)
	if err != nil {
		log.Fatalf("db connect failed: %v", err)
	}

	srv := server.NewServer(gdb)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("starting server on :%s", port)
	if err := srv.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
