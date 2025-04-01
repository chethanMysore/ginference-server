package authcontroller

import (
	"fmt"
	"html"
	"net/http"
	"strings"
	"time"

	config "example/ginference-server/config/devconfig"
	"example/ginference-server/data"
	"example/ginference-server/models/user"
	"example/ginference-server/utils"
	"example/ginference-server/utils/token"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// @BasePath /api/v1
// PingExample godoc
// @Summary Create new user
// @Schemes
// @Description Register new user for inference
// @Tags Auth
// @Param User body user.UserCreate true "Create User"
// @Accept json
// @Produce json
// @Success 201 {object} user.User
// @Router /auth/register [post]
func Register(c *gin.Context) {
	var newUser user.UserCreate
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	newUser.Username = html.EscapeString(strings.TrimSpace(newUser.Username))
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
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	var auth user.UserAuth
	pwdHash, err := utils.Hasher(newUser.Password)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	auth.UserID = usr.UserID
	auth.PasswordHash = string(pwdHash)
	auth.Username = usr.Username
	auth.Role = config.UserRoles.User
	if err := data.Create(auth, config.DBName, config.AuthCollection); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	token, tokenMaxAge, err := token.GenerateToken(usr.UserID.String())
	domain := strings.Split(c.Request.Host, ":")[0]
	c.SetCookie("access_token", token, int(tokenMaxAge), "/", domain, false, true)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusCreated, struct {
		User user.User `json:"user"`
	}{User: usr})
}

// @BasePath /api/v1
// PingExample godoc
// @Summary User Login
// @Schemes
// @Description performs user authentication and returns JWT Auth token on success
// @Tags Auth
// @Security BasicAuth
// @Accept json
// @Produce json
// @Success 201 {string} success
// @Router /auth/login [get]
func Login(c *gin.Context) {
	req := c.Request
	username, pwd, ok := req.BasicAuth()
	if !ok {
		c.String(http.StatusBadRequest, "Missing Basic Authorization")
		return
	}
	filter := bson.D{{Key: "username", Value: username}}
	findOptions := options.Find()
	var registeredUsers []user.UserAuth
	registeredUsers, err := data.Find(registeredUsers, config.DBName, config.AuthCollection, filter, findOptions)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	if len(registeredUsers) == 0 {
		c.IndentedJSON(http.StatusUnauthorized, fmt.Sprintf("No user found with %s", username))
		return
	}
	usr, err := utils.First(registeredUsers)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := utils.VerifyHash(pwd, usr.PasswordHash); err != nil {
		c.String(http.StatusUnauthorized, "Incorrect Password")
		return
	}
	token, tokenMaxAge, err := token.GenerateToken(usr.UserID.String())
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	domain := strings.Split(c.Request.Host, ":")[0]
	c.SetCookie("access_token", token, int(tokenMaxAge), "/", domain, false, true)
	c.String(http.StatusOK, "Login Successful")
}
