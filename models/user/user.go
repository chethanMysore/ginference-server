package user

import (
	"example/ginference-server/utils"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID    uuid.UUID `json:"userID"`
	UserName  string    `json:"userName"`
	CreatedAt time.Time `json:"createdAt"`
}

func (u User) ErrEmptyList() string {
	return "No users registered!"
}

func (u User) ErrNotFound(params ...any) string {
	if len(params) == 0 {
		return "No such user found!"
	} else {
		return fmt.Sprintf("No such user found with %v", params[0])
	}
}

type Users []User

func (usrs Users) FindByName(userName string) (User, string) {
	filteredUsers, err := utils.Filter(usrs, func(usr User) bool {
		return strings.Contains(strings.ToLower(usr.UserName), strings.ToLower(userName))
	})
	if err == "" {
		return utils.First(filteredUsers)
	} else {
		var userErr User
		return userErr, userErr.ErrNotFound(userName)
	}
}

func (usrs Users) FindByUUID(uuid string) (User, string) {
	filteredUsers, err := utils.Filter(usrs, func(usr User) bool {
		return strings.Contains(strings.ToLower(usr.UserID.String()), strings.ToLower(uuid))
	})
	if err == "" {
		return utils.First(filteredUsers)
	} else {
		var userErr User
		return userErr, userErr.ErrNotFound(uuid)
	}
}
