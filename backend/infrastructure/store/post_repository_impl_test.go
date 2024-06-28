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

func TestGetDetail(t *testing.T) {
	ctx := context.Background()
	postRepo := do.MustInvoke[repository.PostRepository](testInjector)

	t.Run("Get a post detail successfully", func(t *testing.T) {
		id := 1
		result, err := postRepo.GetDetail(ctx, id)
		if err != nil {
			t.Errorf("Repository returns error: %v", err.Error())
		}
		if result.IsEmpty() {
			t.Error("Nil")
		} else {
			fmt.Printf("ID: %d, Title: %s, Body: %s\n", result.ID, result.Title, result.Body)
		}
	})

	t.Run("Getting a post detail that does not exist", func(t *testing.T) {
		id := 999
		result, err := postRepo.GetDetail(ctx, id)
		if err != nil {
			t.Errorf("Repository returns error: %v", err.Error())
		}
		if !result.IsEmpty() {
			t.Errorf("Not nil %+v", result)
		}
	})

	t.Run("Getting a post detail with a negative ID", func(t *testing.T) {
		id := -1
		result, err := postRepo.GetDetail(ctx, id)
		if err != nil {
			t.Errorf("Repository returns error: %v", err.Error())
		}
		if !result.IsEmpty() {
			t.Errorf("Not nil %+v", result)
		}
	})
}

func TestList(t *testing.T) {
	ctx := context.Background()
	postRepo := do.MustInvoke[repository.PostRepository](testInjector)

	t.Run("List posts successfully", func(t *testing.T) {
		result, err := postRepo.List(ctx)
		if err != nil {
			t.Errorf("Repository returns error: %v", err.Error())
		}
		if len(result) == 0 {
			t.Error("Empty")
		} else {
			for _, r := range result {
				fmt.Printf("ID: %d, Title: %s\n", r.ID, r.Title)
			}
		}
	})
}

func TestUpdate(t *testing.T) {
	ctx := context.Background()
	postRepo := do.MustInvoke[repository.PostRepository](testInjector)

	t.Run("Update a post successfully", func(t *testing.T) {
		id := 1
		title := "new title"
		body := "new body"
		err := postRepo.Update(ctx, id, title, body)
		if err != nil {
			t.Errorf("Repository returns error: %v", err.Error())
		}

		result, err := postRepo.GetDetail(ctx, id)
		if err != nil {
			t.Errorf("Repository returns error: %v", err.Error())
		}
		if result.IsEmpty() {
			t.Error("Nil")
		} else {
			if result.Title != title {
				t.Errorf("Wrong value: %+v", result)
			}
			if result.Body != body {
				t.Errorf("Wrong value: %+v", result)
			}
		}
	})

	t.Run("Update a post that does not exist", func(t *testing.T) {
		id := 999
		title := "new title"
		body := "new body"
		err := postRepo.Update(ctx, id, title, body)
		if err != nil {
			t.Errorf("Repository returns error: %v", err.Error())
		}
	})
}

func TestDelete(t *testing.T) {
	ctx := context.Background()
	postRepo := do.MustInvoke[repository.PostRepository](testInjector)

	t.Run("Delete a post successfully", func(t *testing.T) {
		id := 1
		err := postRepo.Delete(ctx, id)
		if err != nil {
			t.Errorf("Repository returns error: %v", err.Error())
		}

		result, err := postRepo.GetDetail(ctx, id)
		if err != nil {
			t.Errorf("Repository returns error: %v", err.Error())
		}
		if !result.IsEmpty() {
			t.Errorf("Not nil %+v", result)
		}
	})

	t.Run("Delete a post that does not exist", func(t *testing.T) {
		id := 999
		err := postRepo.Delete(ctx, id)
		if err != nil {
			t.Errorf("Repository returns error: %v", err.Error())
		}
	})
}
