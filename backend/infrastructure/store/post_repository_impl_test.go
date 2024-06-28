package store_test

import (
	"context"
	"fmt"
	"myapp/domain/repository"
	"testing"

	"github.com/samber/do"
)

func TestCreate(t *testing.T) {
	ctx := context.Background()
	postRepo := do.MustInvoke[repository.PostRepository](testInjector)

	t.Run("Create a post successfully", func(t *testing.T) {
		title := "title"
		body := "body"
		userID := 1
		id, err := postRepo.Create(ctx, userID, title, body)
		if err != nil {
			t.Errorf("Repository returns error: %v", err.Error())
		}
		if id == 0 {
			t.Error("ID is 0")
		} else {
			fmt.Printf("ID: %d\n", id)
		}
	})
}
