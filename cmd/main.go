package main

import (
	"fmt"

	"github.com/MdSadiqMd/Phone-Normalisier/internal/db"
	"github.com/MdSadiqMd/Phone-Normalisier/internal/repo"
	"github.com/MdSadiqMd/Phone-Normalisier/pkg"
	_ "github.com/lib/pq"
)

func main() {
	db, err := db.ConnectToDB()
	if err != nil {
		panic(err)
	}

	_, err = repo.InsertData(db, "123456789")
	must(err)
	number, err := repo.GetPhone(db, 1)
	must(err)
	fmt.Println("number:= ", number)
	numbers, err := repo.AllPhones(db)
	must(err)
	for _, p := range numbers {
		fmt.Printf("working on ....%v\n", p)
		number := pkg.Normalize(p.Number)
		if number != p.Number {
			fmt.Println("Updating or Remvoing...", number)
			existing, err := repo.FindNumber(db, number)
			must(err)
			if existing != nil {
				repo.DeleteNumber(db, p.Id)
			} else {
				p.Number = number
				repo.UpdateNumber(db, p)
			}
		} else {
			fmt.Println("No changes required")
		}
	}
	must(db.Ping())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
