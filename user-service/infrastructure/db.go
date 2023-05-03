package infrastructure

import (
	"database/sql"
	"github.com/spf13/viper"
	"log"
	"time"
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

func ConnectToDB() *sql.DB {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	var err error
	count := 0

	for db == nil {
		db, err = openDB(viper.GetString("dsn"))
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
