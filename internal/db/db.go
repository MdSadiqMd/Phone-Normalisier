package db

import (
	"database/sql"
	"fmt"

	"github.com/MdSadiqMd/Phone-Normalisier/pkg"
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
	number, err := getPhone(db, 1)
	must(err)
	fmt.Println("number:= ", number)
	numbers, err := allPhones(db)
	must(err)
	for _, p := range numbers {
		fmt.Printf("working on ....%v\n", p)
		number := pkg.Normalize(p.number)
		if number != p.number {
			fmt.Println("Updating or Remvoing...", number)
			existing, err := findNumber(db, number)
			must(err)
			if existing != nil {
				deleteNumber(db, p.id)
			} else {
				p.number = number
				updateNumber(db, p)
			}
		} else {
			fmt.Println("No changes required")
		}
	}
	must(db.Ping())
}

type phone struct {
	id     int
	number string
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

func getPhone(db *sql.DB, id int) (string, error) {
	statement := `SELECT value FROM phone_numbers WHERE id=$1`
	var number string
	err := db.QueryRow(statement, id).Scan(&number)
	if err != nil {
		panic(err)
	}
	return number, nil
}

func allPhones(db *sql.DB) ([]phone, error) {
	statement := `SELECT id, value FROM phone_numbers`
	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret []phone
	for rows.Next() {
		var p phone
		err := rows.Scan(&p.id, &p.number)
		if err != nil {
			return nil, err
		}
		ret = append(ret, p)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func findNumber(db *sql.DB, number string) (*phone, error) {
	var p phone
	statement := `SELECT * FROM phone_numbers WHERE value=$1`
	row := db.QueryRow(statement, number)
	err := row.Scan(&p.id, &p.number)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &p, nil
}

func updateNumber(db *sql.DB, p phone) error {
	statement := `UPDATE phone_numbers SET value=$2 WHERE id=$1`
	_, err := db.Exec(statement, p.id, p.number)
	return err
}

func deleteNumber(db *sql.DB, id int) error {
	statement := `DELETE FROM phone_numbers WHERE ID=$1`
	_, err := db.Exec(statement, id)
	return err
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
