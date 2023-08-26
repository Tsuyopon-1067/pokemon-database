package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

const SQL_NAME = "mysql"
const SQL_LOGIN = "uuu:hogehoge@/Pokemon"

func printAll() {
	//scraping()

    db, _ := sql.Open(SQL_NAME, SQL_LOGIN)
    defer db.Close()

	// Execute the query
    results, err := db.Query("SELECT * FROM Hgss")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    for results.Next() {
		var data Pokemon

        err = results.Scan(&data.Id, &data.Name, &data.H, &data.A, &data.B, &data.C, &data.D, &data.S, &data.Sum)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
		fmt.Printf("%d, %s: [%d, %d, %d, %d, %d, %d]->%d\n", data.Id, data.Name, data.H, data.A, data.B, data.C, data.D, data.S, data.Sum)
    }
}

func insert(pokemon Pokemon) {
	//scraping()

    db, _ := sql.Open(SQL_NAME, SQL_LOGIN)
    defer db.Close()

	// Execute the query
	query := fmt.Sprintf("INSERT INTO Hgss VALUES (%d, '%s', %d, %d, %d, %d, %d, %d, %d)\n", pokemon.Id, pokemon.Name, pokemon.H, pokemon.A, pokemon.B, pokemon.C, pokemon.D, pokemon.S, pokemon.Sum)
    _, err := db.Query(query)

    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
}