package interfaces

import (
	"myapp/internal/entities"
)

type UserRepository interface {
	Get(id int) (*entities.User, error)
}
