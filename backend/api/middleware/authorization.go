package middleware

import (
	"errors"
	"myapp/application"
	"myapp/domain/model"
	"myapp/domain/service"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

type AuthorizationMiddleware struct {
	idTokenService service.IdtokenService
	userUsecase    application.GetUserUsecase
}

func NewAuthorizationMiddleware(i *do.Injector) (*AuthorizationMiddleware, error) {
	idTokenService := do.MustInvoke[service.IdtokenService](i)
	userUsecase := do.MustInvoke[application.GetUserUsecase](i)
	return &AuthorizationMiddleware{
		idTokenService: idTokenService,
		userUsecase:    userUsecase,
	}, nil
}

func (am *AuthorizationMiddleware) Exec() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil {
			_ = c.AbortWithError(401, errors.New("unauthorized1"))
			return
		}

		userID, err := am.idTokenService.VerifyIDToken(token)
		if err != nil {
			_ = c.AbortWithError(401, errors.New("unauthorized2"))
			return
		}

		result, err := am.userUsecase.Execute(c, application.GetUserUsecaseInput{
			ID: userID,
		})
		if err != nil {
			_ = c.AbortWithError(401, errors.New("unauthorized3"))
			return
		}

		SetUserAuthContext(c, result.User)
		c.Next()
	}
}

func SetUserAuthContext(c *gin.Context, user model.User) {
	c.Set("user", user)
}

func GetUserAuthContext(c *gin.Context) (model.User, bool) {
	userAny, ok := c.Get("user")
	if !ok {
		return model.User{}, false
	}
	user, ok := userAny.(model.User)
	if !ok {
		return model.User{}, false
	}
	return user, true
}
