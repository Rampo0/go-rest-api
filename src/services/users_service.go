package services

import (
	"multi-lang-microservice/users/src/domain/users"
	"multi-lang-microservice/users/src/utils/crypto_utils"
	"multi-lang-microservice/users/src/utils/data_utils"
	"multi-lang-microservice/users/src/utils/errors"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersServiceInterface interface {
	GetUser(int64) (*users.User, *errors.RestErr)
	CreateUser(users.User) (*users.User, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	Search(string) (users.Users, *errors.RestErr)
}

type usersService struct {
}

func (*usersService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := users.User{ID: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return &result, nil
}

func (*usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.DateCreated = data_utils.GetNowDBFormat()
	user.Status = users.StatusActive

	user.Password = crypto_utils.Hash(user.Password)

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (*usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {

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

func (*usersService) DeleteUser(userId int64) *errors.RestErr {
	user := &users.User{ID: userId}
	return user.Delete()
}

func (*usersService) Search(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
