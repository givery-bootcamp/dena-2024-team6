package repositories

import "gorm.io/gorm"

type SigninRepository struct {
	Conn *gorm.DB
}

func NewSigninRepository(conn *gorm.DB) *SigninRepository {
	return &SigninRepository{
		Conn: conn,
	}
}

func (r *SigninRepository) Signin(username string, password string) (string, error) {
	return "{\"foo\":\"token\"}", nil
}
