package users

import (
	"fmt"
	"strings"

	"github.com/rampo0/go-utils/mysql_utils"
	"github.com/rampo0/go-utils/rest_error"

	"github.com/rampo0/multi-lang-microservice/users/src/datasources/mysql/users_db"
	"github.com/rampo0/multi-lang-microservice/users/src/logger"
)

const (
	queryGetUser                   = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id = ?;"
	queryInsertUser                = "INSERT INTO users(first_name, last_name, email, date_created, password, status) VALUES (?,?,?,?,?,?);"
	queryUpdateUser                = "UPDATE users SET first_name=?, last_name=? ,email=? WHERE id=?;"
	queryDeleteUser                = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus          = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
	queryGetUserByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status, password FROM users WHERE email = ? AND password = ? AND status=?;"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) FindByStatus(status string) ([]User, *rest_error.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, rest_error.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	rows, err := stmt.Query(status)

	if err != nil {
		return nil, rest_error.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, rest_error.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}

	return results, nil
}

func (user *User) Get() *rest_error.RestErr {

	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return rest_error.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)

	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		logger.Error("error when trying to get user by id", err)
		return rest_error.NewInternalServerError("database error")
	}

	return nil
}

func (user *User) Save() *rest_error.RestErr {

	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return rest_error.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Password, user.Status)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return rest_error.NewInternalServerError(fmt.Sprintf("Error when trying to get lat user : %s", err.Error()))
	}
	user.ID = userId
	return nil
}

func (user *User) Update() *rest_error.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return rest_error.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, updateErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if updateErr != nil {
		return mysql_utils.ParseError(updateErr)
	}

	return nil
}

func (user *User) Delete() *rest_error.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return rest_error.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, delErr := stmt.Exec(user.ID)
	if delErr != nil {
		return mysql_utils.ParseError(delErr)
	}

	return nil
}

func (user *User) Login() *rest_error.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUserByEmailAndPassword)
	if err != nil {
		return rest_error.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email, user.Password, StatusActive)

	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status, &user.Password); err != nil {
		if strings.Contains(err.Error(), mysql_utils.ErrorNoRows) {
			return rest_error.NewNotFoundError("invalid credentials")
		}
		logger.Error("error when trying to get user by id", err)
		return rest_error.NewInternalServerError("database error")
	}

	return nil

}
