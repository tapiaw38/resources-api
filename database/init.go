package database

import (
	"context"
	"database/sql"
	"log"
	"sync"
)

var (
	data *sql.DB
	once sync.Once
)

// InitDB is the function that initializes the database
func InitDB() *sql.DB {

	db, err := Connect()

	if err != nil {
		panic(err)
	}

	err = MakeMigration(db)

	if err != nil {
		panic(err)
	}

	log.Println("Migration complete")

	return db
}

// NewConnection is the function to get a only connection
func NewConnection() *sql.DB {

	once.Do(func() {
		data = InitDB()
	})

	return data
}

// Data is the function to get the database usage in the packages that need it
func Data() *sql.DB {
	return data
}

// Check that the connection to the database is available
func CheckConnection() bool {

	db, err := data.Conn(context.Background())

	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.PingContext(context.Background())

	return err == nil
}
