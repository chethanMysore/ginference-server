package usercontroller

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	config "example/ginference-server/config/devconfig"
	"example/ginference-server/data"
	"example/ginference-server/models/user"
	"example/ginference-server/utils"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary Get all users
// @Schemes
// @Description Find all users registered with the ginference-server
// @Tags Users
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200 {array} user.User
// @Router /users [get]
func GetAllUsers(c *gin.Context) {
	//authUser, ok := c.Get("authUser")
	// if !ok {
	// 	c.String(http.StatusInternalServerError, "authUser not Found in JWT token")
	// 	c.Abort()
	// 	return
	// }
	filter := bson.D{{}}
	findOptions := options.Find()
	var registeredUsers user.Users
	registeredUsers, err := data.Find(registeredUsers, config.DBName, config.UserCollection, filter, findOptions)
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

// @BasePath /api/v1

// PingExample godoc
// @Summary Search user by ID
// @Schemes
// @Description Find the user created with the given ID
// @Tags Users
// @Security ApiKeyAuth
// @Param id path string true "User ID" minlength(36) maxlength(36)
// @Accept json
// @Produce json
// @Success 200 {object} user.User
// @Router /users/id/{id} [get]
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
	registeredUsers, err := data.Find(registeredUsers, config.DBName, config.UserCollection, filter, findOptions)
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

// @BasePath /api/v1

// PingExample godoc
// @Summary Search users by name
// @Schemes
// @Description Find the users created with the given name
// @Tags Users
// @Security ApiKeyAuth
// @Param name path string true "Name" minlength(2) maxlength(18)
// @Accept json
// @Produce json
// @Success 200 {array} user.User
// @Router /users/name/{name} [get]
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
	registeredUsers, err := data.Find(registeredUsers, config.DBName, config.UserCollection, filter, findOptions)
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

// @BasePath /api/v1

// PingExample godoc
// @Summary Search user by username
// @Schemes
// @Description Find the user created with the given username
// @Tags Users
// @Security ApiKeyAuth
// @Param username path string true "Username" minlength(5) maxlength(18)
// @Accept json
// @Produce json
// @Success 200 {object} user.User
// @Router /users/username/{username} [get]
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
	registeredUsers, err := data.Find(registeredUsers, config.DBName, config.UserCollection, filter, findOptions)
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

// @BasePath /api/v1

// PingExample godoc
// @Summary Search user role by userID
// @Schemes
// @Description Find the user role created with the given username
// @Tags Users
// @Security ApiKeyAuth
// @Param id path string true "UserID" minlength(36) maxlength(36)
// @Accept json
// @Produce json
// @Success 200 {string} role
// @Router /users/auth/id/{id} [get]
func GetUserRoleByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.String(http.StatusBadRequest, "Please specify the userID")
		return
	}
	userID, err := uuid.Parse(id)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Invalid uuid - %s", id))
		return
	}
	filter := bson.D{{Key: "userid", Value: userID}}
	findOptions := options.Find()
	var usrs []user.UserAuth
	usrs, err = data.Find(usrs, config.DBName, config.AuthCollection, filter, findOptions)
	if err != nil || len(usrs) == 0 {
		c.String(http.StatusNotFound, fmt.Sprintf("No users found with the userID - %s", id))
		return
	}
	usr, err := utils.First(usrs)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, usr.Role)
}

func CreateNewUser(c *gin.Context) {
	var newUser user.UserCreate
	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	filter := bson.D{{Key: "username", Value: newUser.Username}}
	findOptions := options.Find()
	var registeredUsers user.Users
	registeredUsers, err := data.Find(registeredUsers, config.DBName, config.UserCollection, filter, findOptions)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	if len(registeredUsers) > 0 {
		c.IndentedJSON(http.StatusNotFound, fmt.Sprintf("User with the user name '%s' already exists", newUser.Username))
		return
	}
	var usr user.User
	usr.UserID = uuid.New()
	usr.FirstName = newUser.FirstName
	usr.LastName = newUser.LastName
	usr.Username = newUser.Username
	usr.EmailID = newUser.EmailID
	usr.Phone = newUser.Phone
	usr.CountryCode = newUser.CountryCode
	usr.FullName = strings.Join([]string{newUser.FirstName, newUser.LastName}, " ")
	usr.CreatedAt = time.Now()
	usr.ModifiedAt = time.Now()
	if err := data.Create(usr, config.DBName, config.UserCollection); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusCreated, usr)
}

// @BasePath /api/v1

// PingExample godoc
// @Summary Edit a user
// @Schemes
// @Description Update a registered User's details
// @Tags Users
// @Security ApiKeyAuth
// @Param User body user.UserUpdate true "Update User"
// @Accept json
// @Produce json
// @Success 200 {object} user.User
// @Router /users/edit [put]
func EditUser(c *gin.Context) {
	var usrUpdate user.UserUpdate
	if err := c.BindJSON(&usrUpdate); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	filter := bson.D{{Key: "userid", Value: usrUpdate.UserID}}
	findOptions := options.Find()
	var registeredUsers user.Users
	registeredUsers, err := data.Find(registeredUsers, config.DBName, config.UserCollection, filter, findOptions)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	if len(registeredUsers) == 0 {
		c.IndentedJSON(http.StatusNotFound, registeredUsers.ErrNotFound(usrUpdate.UserID).Error())
		return
	}
	usr, err := utils.First(registeredUsers)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	updateOptions := options.UpdateOne().SetUpsert(false)
	usr.FirstName = usrUpdate.FirstName
	usr.LastName = usrUpdate.LastName
	usr.FullName = strings.Join([]string{usrUpdate.FirstName, usrUpdate.LastName}, " ")
	usr.EmailID = usrUpdate.EmailID
	usr.Phone = usrUpdate.Phone
	usr.CountryCode = usrUpdate.CountryCode
	usr.ModifiedAt = time.Now()
	if err := data.EditOne(usr, config.DBName, config.UserCollection, filter, updateOptions); err != nil {
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, usr)
}
