package token

import (
	config "example/ginference-server/config/devconfig"
	"example/ginference-server/utils"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func getUserIDFromClaims(authClaims jwt.MapClaims) (string, error) {
	var (
		ok     bool
		raw    interface{}
		userID string
	)
	raw, ok = authClaims["userID"]
	if !ok {
		return "", fmt.Errorf("missing userID in the JWT token")
	}
	userID, ok = raw.(string)
	if !ok {
		return "", fmt.Errorf("JWT token corrupted. %T type is invalid for userID", userID)
	}
	return userID, nil
}

func GenerateToken(userID string) (string, float64, error) {
	tokenLifespan, err := strconv.Atoi(config.TokenHourLifespan)
	if err != nil {
		return "", 0, err
	}
	apiSecretKey, err := utils.ReadConfig(config.APISecretPath)
	if err != nil {
		return "", 0, err
	}
	tokenMaxAge := time.Hour * time.Duration(tokenLifespan)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":     userID,
		"exp":        time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix(),
		"authorized": true,
	})
	tokenStr, err := token.SignedString([]byte(apiSecretKey))
	if err != nil {
		return "", 0, err
	}
	return tokenStr, tokenMaxAge.Seconds(), nil
}

func ValidateToken(tokenStr string) (string, error) {
	apiSecretKey, err := utils.ReadConfig(config.APISecretPath)
	if err != nil {
		return "", err
	}
	authTokenClaims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(tokenStr, authTokenClaims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(apiSecretKey), nil
	})
	if err != nil {
		return "", err
	}
	if expTime, err := authTokenClaims.GetExpirationTime(); err != nil || expTime.Time.Unix() <= time.Now().Unix() {
		return "", fmt.Errorf("token expired")
	}
	return getUserIDFromClaims(authTokenClaims)
}
