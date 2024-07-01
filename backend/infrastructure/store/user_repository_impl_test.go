package store_test

import (
	"context"
	"fmt"
	"myapp/domain/repository"
	"testing"

	"github.com/samber/do"
)

func TestGetUserByID(t *testing.T) {
	ctx := context.Background()
	// DIContainerから取得
	userRepo := do.MustInvoke[repository.UserRepository](testInjector)

	testcases := []int{1, 2}
	for _, userID := range testcases {
		t.Run(fmt.Sprintf("userID = %d", userID), func(t *testing.T) {
			result, err := userRepo.GetByID(ctx, userID)
			if err != nil {
				t.Errorf("Repository returns error: %v", err.Error())
			}
			if result.IsEmpty() {
				t.Error("Nil")
			} else {
				fmt.Printf("ID: %d, Username: %s\n", result.ID, result.Name)
			}
		})
	}
	// Not found
	t.Run("userID = 3 should be nil", func(t *testing.T) {
		result, err := userRepo.GetByID(ctx, 3)
		if err != nil {
			t.Errorf("Repository returns error: %v", nil)
		}
		if !result.IsEmpty() {
			t.Errorf("Not nil %+v", result)
		}
	})
}

func TestGetByUserNameAndPassword(t *testing.T) {
	ctx := context.Background()
	// DIContainerから取得
	userRepo := do.MustInvoke[repository.UserRepository](testInjector)

	testcases := []struct {
		username string
		password string
		userID   int
		iconURL  string
	}{
		{"taro", "password", 1, "https://example.com/icon/taro.png"},
		{"hanako", "PASSWORD", 2, "https://example.com/icon/hanako.png"},
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
		if err != nil {
			t.Errorf("Repository returns error: %v", nil)
		}
		if !result.IsEmpty() {
			t.Errorf("Not nil %+v", result)
		}
	})
}
