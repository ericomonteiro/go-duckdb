package main

import (
	"database/sql"
	"fmt"
	_ "github.com/marcboeker/go-duckdb"
	"log"
)

const csvQuery = `
SELECT *
FROM read_csv('./products.csv',
    delim = ';',
    header = true,
    columns = {
        'id': 'VARCHAR',
        'name': 'VARCHAR',
        'price': 'VARCHAR',
        'quantity': 'VARCHAR'
    })
ORDER BY id DESC;
`

func main() {
	db, err := sql.Open("duckdb", "")
	if err != nil {
		log.Fatal(err)
	}

	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	var (
		id       int
		name     string
		price    float64
		quantity float64
	)
	rows, err := db.Query(csvQuery)
	if err != nil {
		log.Fatal(err)
	}

	// while rows is not empty
	for rows.Next() {
		err := rows.Scan(&id, &name, &price, &quantity)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id: %d, name: %s, price %f, quantity: %f\n", id, name, price, quantity)
	}

}
