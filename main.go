package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"example/ginference-server/utils"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

type user struct {
	UserID    uuid.UUID `json:"userID"`
	UserName  string    `json:"userName"`
	CreatedAt time.Time `json:"createdAt"`
}

func (u user) ErrEmptyList() string {
	return fmt.Sprintln("No users registered!")
}

func (u user) ErrNotFound() string {
	return fmt.Sprintf("No such user found!")
}

type users []user

func (usrs users) FindByName(userName string) (user, string) {
	filteredUsers, err := utils.Filter(usrs, func(usr user) bool {
		return strings.Contains(strings.ToLower(usr.UserName), strings.ToLower(userName))
	})
	if err == "" {
		return utils.First(filteredUsers)
	} else {
		var userErr user
		return userErr, userErr.ErrNotFound()
	}
}

func (usrs users) FindByUUID(uuid string) (user, string) {
	filteredUsers, err := utils.Filter(usrs, func(usr user) bool {
		return strings.Contains(strings.ToLower(usr.UserID.String()), strings.ToLower(uuid))
	})
	if err == "" {
		return utils.First(filteredUsers)
	} else {
		var userErr user
		return userErr, userErr.ErrNotFound()
	}
}

var registeredUsers = users{
	{UserID: uuid.New(), UserName: "Tom", CreatedAt: time.Now()},
	{UserID: uuid.New(), UserName: "Joe", CreatedAt: time.Now()},
	{UserID: uuid.New(), UserName: "Harry", CreatedAt: time.Now()},
}

type model struct {
	ModelID   uuid.UUID `json:"id"`
	ModelName string    `json:"name"`
	CreatedBy user      `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}

var user1, err1 = registeredUsers.FindByName("Tom")
var user2, err2 = registeredUsers.FindByName("Joe")
var user3, err3 = registeredUsers.FindByName("Joe")

var models = []model{
	{ModelID: uuid.New(), ModelName: "pickachu_1", CreatedBy: user1, CreatedAt: time.Now()},
	{ModelID: uuid.New(), ModelName: "bulbasaur_1", CreatedBy: user2, CreatedAt: time.Now()},
	{ModelID: uuid.New(), ModelName: "bulbasaur_2", CreatedBy: user2, CreatedAt: time.Now()},
}

func getAllUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, registeredUsers)
}

func getAllModels(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models)
}

func main() {
	router := gin.Default()
	router.GET("/users", getAllUsers)
	router.GET("/models", getAllModels)
	router.Run("localhost:8080")
}
