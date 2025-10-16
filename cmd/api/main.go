package main

import (
	"iot/internal/db"
	"iot/internal/env"
	"iot/internal/store"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	dbconfig := &dbconfig{
		user: env.GetString("CH_USER", "default"),
		pswd: env.GetString("CH_PSWD", ""),
		addr: env.GetString("CH_ADDR", "localhost:8123"),
		db:   env.GetString("CH_DB", "metrics"),
	}

	cfg := &config{
		addr:        env.GetString("ADDR", ":8080"),
		db:          *dbconfig,
		frontendURL: env.GetString("FRONTEND_URL", "http://localhost:3000"),
	}

	ch_conn, err := db.New(dbconfig.user, dbconfig.pswd, dbconfig.addr, dbconfig.db)
	if err != nil {
		panic(err)
	}
	log.Printf("Connected to ClickHouse on %v ", dbconfig.addr)

	store := store.NewStore(ch_conn)

	app := &application{
		config: *cfg,
		store:  store,
	}

	// start data simulation
	SimulateData(store)

	mux := app.Mount()
	log.Fatal(app.run(mux))
}
