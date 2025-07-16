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
	_, err = insertData(db, "123456789")
	must(err)
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

func insertData(db *sql.DB, phone string) (int, error) {
	statement := `INSERT INTO phone_numbers(value) VALUES($1) RETURNING id`
	var id int
	err := db.QueryRow(statement, phone).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
