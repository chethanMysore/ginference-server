package token

import (
	config "example/ginference-server/config/devconfig"
	"example/ginference-server/utils"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID string) (string, error) {
	tokenLifespan, err := strconv.Atoi(config.TokenHourLifespan)
	if err != nil {
		return "", err
	}
	apiSecretKey, err := utils.ReadConfig(config.APISecretPath)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":     userID,
		"exp":        time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix(),
		"authorized": true,
	})
	tokenStr, err := token.SignedString([]byte(apiSecretKey))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func ValidateToken(tokenStr string) (*jwt.Token, error) {
	apiSecretKey, err := utils.ReadConfig(config.APISecretPath)
	if err != nil {
		return nil, err
	}
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(apiSecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, err
}
