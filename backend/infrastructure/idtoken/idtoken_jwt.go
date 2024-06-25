package idtoken

import (
	"errors"
	"log"
	"myapp/config"
	"myapp/domain/service"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/samber/do"
)

type IDTokenJwtImpl struct {
	key string
}

func NewIDTokenJwtImpl(i *do.Injector) (service.IdtokenService, error) {
	return &IDTokenJwtImpl{
		key: config.JwtKey,
	}, nil
}

// Generate implements service.IdtokenService.
func (i *IDTokenJwtImpl) Generate(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID":        id,
		"ExpiresAt": time.Now().Add(time.Hour * 24).Unix(),
	})

	signedToken, err := token.SignedString([]byte(i.key))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// Verify implements service.IdtokenService.
func (i *IDTokenJwtImpl) VerifyIDToken(token string) (int, error) {
	jwtKey := []byte(i.key)
	claims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	// tokenが不正な場合
	if err != nil {
		return 0, err
	}

	claimsMap, ok := claims.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid credentials")
	}
	userIDraw, ok := claimsMap["ID"].(string)
	if !ok {
		log.Println(userIDraw)
		return 0, errors.New("invalid id credentials")
	}
	userID, err := strconv.Atoi(userIDraw)

	return userID, nil
}
