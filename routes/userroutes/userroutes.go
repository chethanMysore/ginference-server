package userroutes

import (
	"net/http"
	"strings"
	"time"

	"example/ginference-server/config/devconfig"
	"example/ginference-server/data"
	"example/ginference-server/models/user"
	"example/ginference-server/utils"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// GET (/users)
func GetAllUsers(c *gin.Context) {
	filter := bson.D{{}}
	findOptions := options.Find()
	var registeredUsers user.Users
	registeredUsers, err := data.Find(registeredUsers, devconfig.DBName, devconfig.UserCollection, filter, findOptions)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	if len(registeredUsers) == 0 {
		c.IndentedJSON(http.StatusNotFound, registeredUsers.ErrEmptyList().Error())
		return
	}
	c.IndentedJSON(http.StatusOK, registeredUsers)
}

// GET (/users/id/:id)
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.IndentedJSON(http.StatusBadRequest, "Please specify the id or use /users for a generic search")
		return
	}
	userID, parseErr := uuid.Parse(id)
	if parseErr != nil {
		c.IndentedJSON(http.StatusBadRequest, "invalid userid")
		return
	}
	filter := bson.D{{Key: "userid", Value: userID}}
	findOptions := options.Find()
	var registeredUsers user.Users
	registeredUsers, err := data.Find(registeredUsers, devconfig.DBName, devconfig.UserCollection, filter, findOptions)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	if len(registeredUsers) == 0 {
		c.IndentedJSON(http.StatusNotFound, registeredUsers.ErrNotFound(id).Error())
		return
	}
	usr, err := utils.First(registeredUsers)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, usr)
}

// GET (/users/name/:name)
func GetUserByName(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.IndentedJSON(http.StatusBadRequest, "Please specify the name or use /users for a generic search")
		return
	}
	name = strings.Join([]string{".*", name, ".*"}, "")
	filter := bson.D{{Key: "fullname", Value: bson.Regex{Pattern: name, Options: "i"}}}
	findOptions := options.Find()
	var registeredUsers user.Users
	registeredUsers, err := data.Find(registeredUsers, devconfig.DBName, devconfig.UserCollection, filter, findOptions)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	if len(registeredUsers) == 0 {
		c.IndentedJSON(http.StatusNotFound, registeredUsers.ErrNotFound(name).Error())
		return
	}
	c.IndentedJSON(http.StatusOK, registeredUsers)
}

func GetUserByUserName(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		c.IndentedJSON(http.StatusBadRequest, "Please specify the username or use /users for a generic search")
		return
	}
	//username = strings.Join([]string{".*", username, ".*"}, "")
	filter := bson.D{{Key: "username", Value: username}}
	findOptions := options.Find()
	var registeredUsers user.Users
	registeredUsers, err := data.Find(registeredUsers, devconfig.DBName, devconfig.UserCollection, filter, findOptions)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	if len(registeredUsers) == 0 {
		c.IndentedJSON(http.StatusNotFound, registeredUsers.ErrNotFound(username).Error())
		return
	}
	usr, err := utils.First(registeredUsers)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, usr)
}

// POST (/users/create)
func CreateNewUser(c *gin.Context) {
	var newUser user.User
	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	newUser.UserID = uuid.New()
	newUser.FullName = strings.Join([]string{newUser.FirstName, newUser.LastName}, " ")
	newUser.CreatedAt = time.Now()
	newUser.ModifiedAt = time.Now()
	if err := data.Create(newUser, devconfig.DBName, devconfig.UserCollection); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusCreated, newUser)
}

// PUT (/users/edit)
func EditUser(c *gin.Context) {
	var usr user.UserUpdate
	if err := c.BindJSON(&usr); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	filter := bson.D{{Key: "userid", Value: usr.UserID}}
	updateOptions := options.UpdateOne().SetUpsert(false)
	usr.ModifiedAt = time.Now()
	if err := data.EditOne(usr, devconfig.DBName, devconfig.UserCollection, filter, updateOptions); err != nil {
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, usr)
}
