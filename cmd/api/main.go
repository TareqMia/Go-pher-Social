package main

import (
	"gopher-social/internal/db"
	"gopher-social/internal/env"
	"gopher-social/internal/env/store"
	"log"
)

func main() {

	config := &config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(config.db.addr, config.db.maxOpenConns, config.db.maxIdleConns, config.db.maxIdleTime)

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	log.Println("DB connection established")

	store := store.NewStorage(db)

	app := &application{
		config: *config,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
