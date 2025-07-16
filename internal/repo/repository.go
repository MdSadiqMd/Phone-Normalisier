package repo

import "database/sql"

type phone struct {
	id     int
	number string
}

func InsertData(db *sql.DB, phone string) (int, error) {
	statement := `INSERT INTO phone_numbers(value) VALUES($1) RETURNING id`
	var id int
	err := db.QueryRow(statement, phone).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func GetPhone(db *sql.DB, id int) (string, error) {
	statement := `SELECT value FROM phone_numbers WHERE id=$1`
	var number string
	err := db.QueryRow(statement, id).Scan(&number)
	if err != nil {
		panic(err)
	}
	return number, nil
}

func AllPhones(db *sql.DB) ([]phone, error) {
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

func FindNumber(db *sql.DB, number string) (*phone, error) {
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

func UpdateNumber(db *sql.DB, p phone) error {
	statement := `UPDATE phone_numbers SET value=$2 WHERE id=$1`
	_, err := db.Exec(statement, p.id, p.number)
	return err
}

func DeleteNumber(db *sql.DB, id int) error {
	statement := `DELETE FROM phone_numbers WHERE ID=$1`
	_, err := db.Exec(statement, id)
	return err
}
