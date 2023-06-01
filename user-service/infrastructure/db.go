package infrastructure

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var db *sql.DB

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func (app application) ConnectToDB() *sql.DB {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	var err error
	count := 0

	for db == nil {
		db, err = openDB(Cfg.DSN)
		if err != nil {
			log.Println("postgres is not ready yet:", err)
		}
		count++
		if count > 5 {
			log.Fatal("cant connect to postgres:", err)
		}
		<-ticker.C
	}
	return db
}
