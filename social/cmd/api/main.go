package main

import (
	"fmt"
	"log"

	"github.com/ritwanxidig/golang_challenge/social/internal/db"
	"github.com/ritwanxidig/golang_challenge/social/internal/env"
	"github.com/ritwanxidig/golang_challenge/social/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8000"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://postgres:1234@localhost:5432/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 25),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 10),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, error := db.New(cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime)

	if error != nil {
		log.Panic(error)
	}

	defer db.Close()
	fmt.Println("Database connection pool established...")

	store := store.NewPostgresStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
