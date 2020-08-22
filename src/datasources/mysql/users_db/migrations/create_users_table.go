package migrations

import (
	"fmt"
	"multi-lang-microservice/users/src/utils/mysql_utils"
)

func User() {
	tableName := "users"

	qcheck := "DROP TABLE IF EXISTS %s"
	// exec query
	mysql_utils.ExecQuery(fmt.Sprintf(qcheck, tableName))

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

	mysql_utils.ExecQuery(fmt.Sprintf(query, values...))
}
