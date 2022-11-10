package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	PostgresUser     = "postgres"
	PostgresDatabase = "crud"
	PostgresPassword = "7"
	PostgresHost     = "localhost"
	PostgresPort     = 5432
)

type Transfer struct{
	Id int64
	Amount float64
	PaymentType string
}

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s database=%s password=%s sslmode=disable",
		PostgresHost,
		PostgresPort,
		PostgresUser,
		PostgresDatabase,
		PostgresPassword,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed connect to database: %v", err)
	}

	query := `
		select
			id,
			amount,
			payment_type
		from transfer
		where payment_type=$1
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatalf("failed to prepare query: %v",err)
	}

	rows, err := stmt.Query("click")
	if err != nil {
		log.Fatalf("failed to prepare query: %v",err)
	}
	defer rows.Close()

	result := make([]Transfer, 0)
	for rows.Next(){
		var tr Transfer

		err := rows.Scan(
			&tr.Id,
			&tr.Amount,
			&tr.PaymentType,
		)
		if err != nil {
			log.Fatalf("failed to scan transfer: %v",err)
		}

		result = append(result, tr)
	}

	fmt.Println(result)
}