package db

import (
	"database/sql"
)

func ConnectToDB() {
	databaseName := "Phone"
	psqlConnectionString := "postgresql://sadiq:sadiq@localhost:5432/phone"
	db, err := sql.Open("postgres", psqlConnectionString)
	must(err)

	err = createDatabase(db, databaseName)
	must(err)
	db.Close()

	db, err = sql.Open("postgres", psqlConnectionString)
	must(err)
	defer db.Close()

	must(createPhoneNumberTable(db))
	must(db.Ping())
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

func must(err error) {
	if err != nil {
		panic(err)
	}
}
