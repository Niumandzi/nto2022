package sqlitedb

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func NewClient(driverName string, filePath string) (*sql.DB, error) {
	db, err := sql.Open(driverName, filePath)
	if err != nil {
		return nil, err
	}

	return db, nil

}

func CreateTables(db *sql.DB) error {

	_, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS contact (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				contact_type TEXT CHECK (contact_type IN ('worker', 'private_client', 'legal_client')),
				name TEXT,
				number TEXT,
				email TEXT
           );

			CREATE TABLE IF NOT EXISTS hotel (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				hotel_name VARCHAR(200),
				location VARCHAR,
				hotel_number VARCHAR,
				worker_id INTEGER,
				hotel_description VARCHAR(500),
				FOREIGN KEY (worker_id) REFERENCES contact (id)
			)`)

	if err != nil {
		return err
	}

	return nil
}
