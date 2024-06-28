package store_test

import (
	"context"
	"fmt"
	"myapp/domain/repository"
	"testing"

	"github.com/samber/do"
)

func TestCreateComment(t *testing.T) {
	ctx := context.Background()
	commentRepo := do.MustInvoke[repository.CommentRepository](testInjector)

	t.Run("Create a comment successfully", func(t *testing.T) {
		postID := 1
		userID := 1
		body := "body"
		id, err := commentRepo.Create(ctx, postID, userID, body)
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

func TestGetByID(t *testing.T) {
	ctx := context.Background()
	commentRepo := do.MustInvoke[repository.CommentRepository](testInjector)

	t.Run("Get a comment by ID successfully", func(t *testing.T) {
		id := 1
		result, err := commentRepo.GetByID(ctx, id)
		if err != nil {
			t.Errorf("Repository returns error: %v", err.Error())
		}
		if result.IsEmpty() {
			t.Error("Nil")
		} else {
			fmt.Printf("ID: %d, PostID: %d, UserID: %d, Body: %s\n", result.ID, result.PostID, result.UserID, result.Body)
		}
	})

	t.Run("Getting a comment that does not exist", func(t *testing.T) {
		id := 999
		result, err := commentRepo.GetByID(ctx, id)
		if err != nil {
			t.Errorf("Repository returns error: %v", err.Error())
		}
		if !result.IsEmpty() {
			t.Errorf("Not nil %+v", result)
		}
	})

	t.Run("Getting a comment with a negative ID", func(t *testing.T) {
		id := -1
		result, err := commentRepo.GetByID(ctx, id)
		if err != nil {
			t.Errorf("Repository returns error: %v", err.Error())
		}
		if !result.IsEmpty() {
			t.Errorf("Not nil %+v", result)
		}
	})
}

func TestListComments(t *testing.T) {
	ctx := context.Background()
	commentRepo := do.MustInvoke[repository.CommentRepository](testInjector)

	t.Run("List comments successfully", func(t *testing.T) {
		postID := 1
		results, err := commentRepo.List(ctx, postID)
		if err != nil {
			t.Errorf("Repository returns error: %v", err.Error())
		}
		if len(results) == 0 {
			t.Error("Empty")
		} else {
			for _, result := range results {
				fmt.Printf("ID: %d, PostID: %d, UserID: %d, Body: %s\n", result.ID, result.PostID, result.UserID, result.Body)
			}
		}
	})

	t.Run("List comments with a post ID that does not exist", func(t *testing.T) {
		postID := 999
		results, err := commentRepo.List(ctx, postID)
		if err != nil {
			t.Errorf("Repository returns error: %v", err.Error())
		}
		if len(results) != 0 {
			t.Errorf("Not empty %+v", results)
		}
	})

	t.Run("List comments with a negative post ID", func(t *testing.T) {
		postID := -1
		results, err := commentRepo.List(ctx, postID)
		if err != nil {
			t.Errorf("Repository returns error: %v", err.Error())
		}
		if len(results) != 0 {
			t.Errorf("Not empty %+v", results)
		}
	})
}

func TestUpdateComment(t *testing.T) {
	ctx := context.Background()
	commentRepo := do.MustInvoke[repository.CommentRepository](testInjector)

	t.Run("Update a comment successfully", func(t *testing.T) {
		postID := 1
		userID := 1
		commentID := 1
		body := "body"

		err := commentRepo.Update(ctx, postID, userID, commentID, body)
		if err != nil {
			t.Errorf("Repository returns error: %v", err.Error())
		}

		result, err := commentRepo.GetByID(ctx, commentID)
		if err != nil {
			t.Errorf("Repository returns error: %v", err.Error())
		}
		if result.IsEmpty() {
			t.Error("Nil")
		}
		if result.Body != body {
			t.Errorf("Wrong value: %+v", result)
		}
	})

	t.Run("Update a comment that does not exist", func(t *testing.T) {
		postID := 1
		userID := 1
		commentID := 999
		body := "body"

		err := commentRepo.Update(ctx, postID, userID, commentID, body)
		if err == nil {
			t.Errorf("Repository did not return error")
		}
	})
}

func TestDeleteComment(t *testing.T) {
	ctx := context.Background()
	commentRepo := do.MustInvoke[repository.CommentRepository](testInjector)

	t.Run("Delete a comment successfully", func(t *testing.T) {
		id := 1
		err := commentRepo.Delete(ctx, id)
		if err != nil {
			t.Errorf("Repository returns error: %v", err.Error())
		}

		result, err := commentRepo.GetByID(ctx, id)
		if err != nil {
			t.Errorf("Repository returns error: %v", err.Error())
		}
		if !result.IsEmpty() {
			t.Errorf("Not nil %+v", result)
		}
	})

	t.Run("Delete a comment that does not exist", func(t *testing.T) {
		id := 999
		err := commentRepo.Delete(ctx, id)
		if err == nil {
			t.Errorf("Repository did not return error")
		}
	})

	t.Run("Delete a comment with a negative ID", func(t *testing.T) {
		id := -1
		err := commentRepo.Delete(ctx, id)
		if err == nil {
			t.Errorf("Repository did not return error")
		}

		_, err = commentRepo.GetByID(ctx, id)
		if err == nil {
			t.Errorf("Repository did not return error")
		}
	})
}
