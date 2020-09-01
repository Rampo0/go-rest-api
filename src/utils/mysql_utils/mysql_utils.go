package mysql_utils

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/rampo0/multi-lang-microservice/users/src/datasources/mysql/users_db"
	"github.com/rampo0/multi-lang-microservice/users/src/utils/errors"
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
