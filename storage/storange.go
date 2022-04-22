package storage

import (
	"context"
	"database/sql"
	"log"
	"sync"
)

var (
	data *Data
	once sync.Once
)

// Data manages the connection to the database.
type Data struct {
	DB *sql.DB
}

// InitDB is the function that initializes the database
func InitDB() {

	db, err := Connect()

	if err != nil {
		panic(err)
	}

	err = MakeMigration(db)

	if err != nil {
		panic(err)
	}

	data = &Data{
		DB: db,
	}

	log.Println("Migration complete")
}

// NewConnection is the function to get only connection
func NewConnection() *Data {
	once.Do(InitDB)

	return data
}

// CheckConnection that the connection to the database is available
func CheckConnection() bool {

	db, err := data.DB.Conn(context.Background())

	if err != nil {
		panic(err)
	}

	err = db.PingContext(context.Background())

	if err != nil {
		panic(err)
	}
	
	return err == nil
}
