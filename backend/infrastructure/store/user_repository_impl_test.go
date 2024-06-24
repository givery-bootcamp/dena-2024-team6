package store_test

import (
	"context"
	"fmt"
	"myapp/domain/repository"
	"testing"

	"github.com/samber/do"
)

func TestGetByUserNameAndPassword(t *testing.T) {
	ctx := context.Background()
	// DIContainerから取得
	userRepo := do.MustInvoke[repository.UserRepository](testInjector)

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
			result, err := userRepo.GetByUserNameAndPassword(ctx, tc.username, tc.password)
			if err != nil {
				t.Errorf("Repository returns error: %v", err.Error())
			}
			if result.IsEmpty() {
				t.Error("Nil")
			} else if result.ID != tc.userID {
				t.Errorf("Wrong value: %+v", result)
			} else {
				fmt.Printf("ID: %d, Username: %s\n", result.ID, result.Name)
			}
		})
	}
	// Not found
	t.Run("username = fr and password = fr should be nil", func(t *testing.T) {
		result, err := userRepo.GetByUserNameAndPassword(ctx, "fr", "fr")
		if err == nil {
			t.Errorf("Repository returns error: %v", nil)
		}
		if !result.IsEmpty() {
			t.Errorf("Not nil %+v", result)
		}
	})
}
