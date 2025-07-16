package main

import (
	"fmt"

	"github.com/MdSadiqMd/Phone-Normalisier/pkg"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println(pkg.Normalize("(123) 456-7890"))
}
