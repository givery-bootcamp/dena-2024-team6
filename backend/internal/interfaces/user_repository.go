package interfaces

import (
	"myapp/internal/entities"
)

type UserRepository interface {
	Get() (*entities.User, error)
}
