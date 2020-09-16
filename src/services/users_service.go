package services

import (
	"github.com/rampo0/go-utils/crypto_utils"
	"github.com/rampo0/go-utils/data_utils"
	"github.com/rampo0/go-utils/rest_error"
	"github.com/rampo0/multi-lang-microservice/users/src/domain/users"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersServiceInterface interface {
	GetUser(int64) (*users.User, *rest_error.RestErr)
	CreateUser(users.User) (*users.User, *rest_error.RestErr)
	UpdateUser(bool, users.User) (*users.User, *rest_error.RestErr)
	DeleteUser(int64) *rest_error.RestErr
	Search(string) (users.Users, *rest_error.RestErr)
	Login(users.LoginRequest) (*users.User, *rest_error.RestErr)
}

type usersService struct {
}

func (*usersService) Login(request users.LoginRequest) (*users.User, *rest_error.RestErr) {
	dao := &users.User{
		Email:    request.Email,
		Password: crypto_utils.GetMD5(request.Password),
	}

	if err := dao.Login(); err != nil {
		return nil, err
	}

	return dao, nil
}

func (*usersService) GetUser(userId int64) (*users.User, *rest_error.RestErr) {
	result := users.User{ID: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return &result, nil
}

func (*usersService) CreateUser(user users.User) (*users.User, *rest_error.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.DateCreated = data_utils.GetNowDBFormat()
	user.Status = users.StatusActive

	user.Password = crypto_utils.GetMD5(user.Password)

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (*usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *rest_error.RestErr) {

	current, err := UsersService.GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName

		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}

	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func (*usersService) DeleteUser(userId int64) *rest_error.RestErr {
	user := &users.User{ID: userId}
	return user.Delete()
}

func (*usersService) Search(status string) (users.Users, *rest_error.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
