package migrations

import (
	"fmt"
	"log"

	"github.com/rampo0/multi-lang-microservice/users/src/datasources/mysql/users_db"
)

func User() {
	tableName := "users"

	qcheck := "DROP TABLE IF EXISTS %s"
	// exec query
	execQuery(fmt.Sprintf(qcheck, tableName))

	query := "CREATE TABLE %s (%s, %s, %s, %s, %s, %s, %s)"

	values := []interface{}{
		tableName,
		"id bigint PRIMARY KEY NOT NULL AUTO_INCREMENT",
		"first_name varchar(100)",
		"last_name varchar(100)",
		"email varchar(100) UNIQUE",
		"date_created datetime",
		"status varchar(100)",
		"password varchar(100)",
	}

	execQuery(fmt.Sprintf(query, values...))
}

func execQuery(query string) {
	res, err := users_db.Client.Exec(query)

	if err != nil {
		log.Fatalf("Error executing query : %v", err)
	}

	if res != nil {
		fmt.Println("Query executed successfully")
	}
}
