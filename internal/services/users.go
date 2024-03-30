package services

import (
	"errors"
	"go-finance-api/internal/dtos"
	"go-finance-api/internal/models"
	"go-finance-api/internal/repository"
	"go-finance-api/internal/util"
)

func FindUsersByName(query string, pageNumber, pageSize int) dtos.Page[dtos.UserListDto] {
	return repository.FindUserList(query, pageNumber, pageSize)
}

func Login(dto dtos.UserLoginDto) (string, error) {
	user := repository.FindUserByLogin(dto.Email)
	if user.ID == 0 {
		return "", errors.New("error trying to log in")
	}
	checkPassword := util.CheckPassword(dto.Password, user.Password)
	if !checkPassword {
		return "", errors.New("error trying to log in")
	}
	token, err := util.CreateJwt(&user)
	return token, err
}

func CreateUser(dto dtos.CreateUserDto) (err error) {
	alredyExists := repository.ExistsUserByLogin(dto.Email)
	if alredyExists {
		err = errors.New("user already exists")
		return
	}
	var user models.User
	roles := repository.FindRolesByName([]string{dto.Role, "ROLE_USER"})
	user = buildUser(dto, roles)
	repository.CreateUser(&user)
	if user.ID == 0 {
		err = errors.New("error creating new user")
	}
	return
}

func buildUser(dto dtos.CreateUserDto, roles []models.Role) models.User {
	password, err := util.EncryptPassword(dto.Password)
	if err != nil {
		return models.User{}
	}
	return models.User{
		Roles:    roles,
		Email:    dto.Email,
		Password: password,
		Active:   true,
		Name:     dto.Name,
	}
}
