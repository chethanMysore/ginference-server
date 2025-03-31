package user

import (
	"example/ginference-server/utils"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID      uuid.UUID `json:"userID"`
	FirstName   string    `json:"firstName" binding:"required,alpha,min=2,max=18"`
	LastName    string    `json:"lastName" binding:"required,alpha,min=2,max=18"`
	FullName    string    `json:"fullName"`
	Username    string    `json:"username" binding:"required,alphanum,min=5,max=18"`
	EmailID     string    `json:"emailID" binding:"required,email"`
	Phone       string    `json:"phone" binding:"required,e164"`
	CountryCode string    `json:"countryCode" binding:"required,iso3166_1_alpha2"`
	CreatedAt   time.Time `json:"createdAt"`
	ModifiedAt  time.Time `json:"modifiedAt"`
}

type UserAuth struct {
	UserID       uuid.UUID `json:"userID" binding:"required"`
	Username     string    `json:"username" binding:"required,alphanum,min=5,max=18"`
	PasswordHash string    `json:"pwdHash" binding:"required"`
	Role         string    `json:"role" binding:"required"`
}

type UserCreate struct {
	FirstName   string `json:"firstName" binding:"required,alpha,min=2,max=18"`
	LastName    string `json:"lastName" binding:"required,alpha,min=2,max=18"`
	Username    string `json:"username" binding:"required,alphanum,min=5,max=18"`
	Password    string `json:"password" binding:"required,min=8,max=18"`
	EmailID     string `json:"emailID" binding:"required,email"`
	Phone       string `json:"phone" binding:"required,e164"`
	CountryCode string `json:"countryCode" binding:"required,iso3166_1_alpha2"`
}

type UserUpdate struct {
	UserID      uuid.UUID `json:"userID" binding:"required"`
	FirstName   string    `json:"firstName" binding:"required,alpha,min=2,max=18"`
	LastName    string    `json:"lastName" binding:"required,alpha,min=2,max=18"`
	EmailID     string    `json:"emailID" binding:"required,email"`
	Phone       string    `json:"phone" binding:"required,e164"`
	CountryCode string    `json:"countryCode" binding:"required,iso3166_1_alpha2"`
}

func (u User) ErrEmptyList() error {
	return fmt.Errorf("no users registered")
}

func (u User) ErrNotFound(params ...any) error {
	if len(params) == 0 {
		return fmt.Errorf("no such user found")
	} else {
		return fmt.Errorf("no such user found with %v", params[0])
	}
}

func (u UserAuth) ErrEmptyList() error {
	return fmt.Errorf("no users registered")
}

func (u UserAuth) ErrNotFound(params ...any) error {
	if len(params) == 0 {
		return fmt.Errorf("no such user found")
	} else {
		return fmt.Errorf("no such user found with %v", params[0])
	}
}

type Users []User

func (u Users) ErrEmptyList() error {
	return fmt.Errorf("no users registered")
}

func (u Users) ErrNotFound(params ...any) error {
	if len(params) == 0 {
		return fmt.Errorf("no such user found")
	} else {
		return fmt.Errorf("no such user found with %v", params[0])
	}
}

func (usrs Users) FindByName(name string) (User, error) {
	filteredUsers, err := utils.Filter(usrs, func(usr User) bool {
		return strings.Contains(strings.ToLower(usr.FullName), strings.ToLower(name))
	})
	if err != nil {
		var userErr User
		return userErr, userErr.ErrNotFound(name)
	}
	return utils.First(filteredUsers)
}

func (usrs Users) FindByUserName(userName string) (User, error) {
	filteredUsers, err := utils.Filter(usrs, func(usr User) bool {
		return strings.Contains(strings.ToLower(usr.Username), strings.ToLower(userName))
	})
	if err != nil {
		var userErr User
		return userErr, userErr.ErrNotFound(userName)
	}
	return utils.First(filteredUsers)
}

func (usrs Users) FindByUUID(uuid string) (User, error) {
	filteredUsers, err := utils.Filter(usrs, func(usr User) bool {
		return strings.Contains(strings.ToLower(usr.UserID.String()), strings.ToLower(uuid))
	})
	if err != nil {
		var userErr User
		return userErr, userErr.ErrNotFound(uuid)
	}
	return utils.First(filteredUsers)
}
