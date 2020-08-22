package mysql_utils

import (
	"fmt"
	"log"
	"multi-lang-microservice/users/src/datasources/mysql/users_db"
	"multi-lang-microservice/users/src/utils/errors"
	"strings"

	"github.com/go-sql-driver/mysql"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("no record match given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("duplicated records")
	}
	return errors.NewInternalServerError("error processing request")
}

func ExecQuery(query string) {
	res, err := users_db.Client.Exec(query)

	if err != nil {
		log.Fatalf("Error executing query : %v", err)
	}

	if res != nil {
		fmt.Println("Query executed successfully")
	}
}
