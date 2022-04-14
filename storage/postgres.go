package storage

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Connect to the database
func Connect() (*sql.DB, error) {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	client := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" sslmode=disable"

	return sql.Open("postgres", client)

}

// MakeMigration creates the tables
func MakeMigration(db *sql.DB) error {
	b, err := ioutil.ReadFile("database/models.sql")

	if err != nil {
		return err
	}

	rows, err := db.Query(string(b))

	if err != nil {
		return err
	}

	return rows.Close()
}
