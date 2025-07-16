package db

import (
	"database/sql"
)

func ConnectToDB() (db *sql.DB, err error) {
	databaseName := "Phone"
	psqlConnectionString := "postgresql://sadiq:sadiq@localhost:5432/phone"
	db, err = Migrate(databaseName, psqlConnectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func createDatabase(db *sql.DB, name string) error {
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS " + name)
	if err != nil {
		return err
	}
	return nil
}

func createPhoneNumberTable(db *sql.DB) error {
	statement := `
		CREATE TABLE IS NOT EXISTS phone_numbers (
			id    SERIAL,
			value VARCHAR(255)
		)
	`
	_, err := db.Exec(statement)
	return err
}

func Migrate(driverName, dataSource string) (*sql.DB, error) {
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return nil, err
	}

	err = createDatabase(db, "phone_numbers")
	if err != nil {
		return nil, err
	}

	err = createPhoneNumberTable(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}
