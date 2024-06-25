package repositories

import (
	"fmt"
	"myapp/internal/external"
	"myapp/internal/interfaces"
	"myapp/internal/repositories"
	"testing"
)

func setupSignin() (interfaces.SigninRepository, func()) {
	db := external.DB.Begin()
	repo := repositories.NewSigninRepository(db)
	teardown := func() {
		db.Rollback()
	}
	return repo, teardown
}

func TestSignin(t *testing.T) {
	repo, teardown := setupSignin()
	defer teardown()

	// Valid testcases
	testcases := []struct {
		username string
		password string
		userID   int
	}{
		{"taro", "password", 1},
		{"hanako", "PASSWORD", 2},
	}
	for _, tc := range testcases {
		t.Run(fmt.Sprintf("username = %s and password = %s", tc.username, tc.password), func(t *testing.T) {
			result, err := repo.Signin(tc.username, tc.password)
			if err != nil {
				t.Errorf("Repository returns error: %v", err.Error())
			}
			if result == nil {
				t.Error("Nil")
			} else if result.ID != tc.userID {
				t.Errorf("Wrong value: %+v", result)
			} else {
				fmt.Printf("ID: %d, Username: %s\n", result.ID, result.Username)
			}
		})
	}
	// Not found
	t.Run("username = fr and password = fr should be nil", func(t *testing.T) {
		result, err := repo.Signin("fr", "fr")
		if err == nil {
			t.Errorf("Repository returns error: %v", nil)
		}
		if result != nil {
			t.Errorf("Not nil %+v", result)
		}
	})
}
