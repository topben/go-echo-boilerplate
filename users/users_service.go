package users

import (
	"sync"

	"github.com/topben/go-echo-boilerplate/common"
	"github.com/topben/go-echo-boilerplate/database"
	"github.com/topben/go-echo-boilerplate/users/models"
)

type usersService struct{}

var singleton UsersService
var once sync.Once

func GetUsersService() UsersService {
	if singleton != nil {
		return singleton
	}
	once.Do(func() {
		singleton = &usersService{}
	})
	return singleton
}

func SetUsersService(service UsersService) UsersService {
	original := singleton
	singleton = service
	return original
}

type UsersService interface {
	FindUserByEmail(email string) *models.User
	AddUser(name string, email string, password string) *models.User
}

func (u *usersService) FindUserByEmail(email string) *models.User {
	db := database.GetInstance()
	var user models.User
	err := db.First(&user, "email = ?", email).Error
	if err == nil {
		return &user
	}
	return nil
}

func (u *usersService) AddUser(name string, email string, password string) *models.User {
	user := models.User{
		Name:     name,
		Role:     common.Admin,
		Email:    email,
		Password: password,
	}
	db := database.GetInstance()
	db.Create(&user)
	return &user
}
