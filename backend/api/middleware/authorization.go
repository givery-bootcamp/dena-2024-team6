package middleware

import (
	"errors"
	"myapp/domain/service"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

type AuthorizationMiddleware struct {
	idTokenService service.IdtokenService
}

func NewAuthorizationMiddleware(i *do.Injector) (*AuthorizationMiddleware, error) {
	idTokenService := do.MustInvoke[service.IdtokenService](i)
	return &AuthorizationMiddleware{
		idTokenService: idTokenService,
	}, nil
}

func (am *AuthorizationMiddleware) Exec() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil {
			_ = c.AbortWithError(401, errors.New("unauthorized"))
			return
		}

		userID, err := am.idTokenService.VerifyIDToken(token)
		if err != nil {
			_ = c.AbortWithError(401, errors.New("unauthorized"))
		}

		// TODO: ユーザの存在チェック

		c.Set("userID", userID)
		c.Next()
	}
}
