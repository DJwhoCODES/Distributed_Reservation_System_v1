package main

import (
	"context"
	"log"

	"github.com/djwhocodes/ticket-reservation/configs"
	pgrepo "github.com/djwhocodes/ticket-reservation/internal/repository/postgres"
)

func main() {
	cfg := configs.Load()

	db, err := pgrepo.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := pgrepo.Seed(context.Background(), db); err != nil {
		log.Fatal(err)
	}

	log.Println("seed completed")
}
