package interfaces

import "myapp/internal/entities"

type SigninRepository interface {
	Signin(username string, password string) (*entities.User, error)
}
