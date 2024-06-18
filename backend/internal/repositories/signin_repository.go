package repositories

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"myapp/internal/entities"

	"gorm.io/gorm"
)

type SigninRepository struct {
	Conn *gorm.DB
}

type User struct {
	ID   int
	Name string
}

func NewSigninRepository(conn *gorm.DB) *SigninRepository {
	return &SigninRepository{
		Conn: conn,
	}
}

func (r *SigninRepository) Signin(username string, password string) (*entities.User, error) {
	var user User
	hashedPassword := sha256.Sum256([]byte(password))
	// strPassword := string(hashedPassword[:])
	strPassword := hex.EncodeToString(hashedPassword[:])
	fmt.Println(strPassword)
	err := r.Conn.Table("users").Select(`
		id,
		name
	`).Where("BINARY name = ? AND BINARY password = ?", username, strPassword).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &entities.User{
		ID:       user.ID,
		Username: user.Name,
	}, nil
}
