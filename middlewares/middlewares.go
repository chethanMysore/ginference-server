package middlewares

import (
	config "example/ginference-server/config/devconfig"
	"example/ginference-server/data"
	"example/ginference-server/models/user"
	"example/ginference-server/utils"
	"example/ginference-server/utils/token"
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// bearerToken := c.GetHeader("Authorization")
		// tokenStr := ""
		// if len(strings.Split(bearerToken, " ")) == 2 {
		// 	tokenStr = strings.Split(bearerToken, " ")[1]
		// } else {
		// 	c.String(http.StatusUnauthorized, "Unauthorized - Please provide valid bearer token")
		// 	c.Abort()
		// 	return
		// }
		tokenStr, err := c.Cookie(config.JWTCookieName)
		if err != nil {
			c.String(http.StatusUnauthorized, fmt.Sprintf("%s cookie not found", config.JWTCookieName))
			c.Abort()
			return
		}
		id, err := token.ValidateToken(tokenStr)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized - Invalid / Corrupted JWT token in access_token cookie")
			c.Abort()
			return
		}
		userID, err := uuid.Parse(id)
		if err != nil {
			c.String(http.StatusUnauthorized, fmt.Sprintf("Corrupted JWT token in access_token cookie: Invalid userID %s", id))
			return
		}
		filter := bson.D{{Key: "userid", Value: userID}}
		findOptions := options.Find()
		var registeredUsers user.Users
		registeredUsers, err = data.Find(registeredUsers, config.DBName, config.UserCollection, filter, findOptions)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			c.Abort()
			return
		}
		if len(registeredUsers) == 0 {
			c.String(http.StatusUnauthorized, fmt.Sprintf("Corrupted JWT token in access_token cookie: %s", registeredUsers.ErrNotFound(userID).Error()))
			c.Abort()
			return
		}
		authUser, err := utils.First(registeredUsers)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			c.Abort()
			return
		}
		c.Set("authUser", authUser)
		c.Next()
	}
}
