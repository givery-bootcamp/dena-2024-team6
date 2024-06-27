package application

import (
	"context"
	"myapp/domain/apperror"
	"myapp/domain/repository"

	"github.com/samber/do"
)

type DeletePostUsecase interface {
	Execute(ctx context.Context, input DeletePostUsecaseInput) error
}

type DeletePostUsecaseInput struct {
	PostID int
	UserID int
}

func NewDeletePostUsecase(i *do.Injector) (DeletePostUsecase, error) {
	postRepo := do.MustInvoke[repository.PostRepository](i)
	return &DeletePostUsecaseInteractor{
		postRepository: postRepo,
	}, nil
}

type DeletePostUsecaseInteractor struct {
	postRepository repository.PostRepository
}

// Execute implements DeletePostUsecase.
func (c *DeletePostUsecaseInteractor) Execute(ctx context.Context, input DeletePostUsecaseInput) error {
	post, err := c.postRepository.GetDetail(ctx, input.PostID)
	if err != nil {
		return apperror.New(apperror.CodeInternalServer, "failed to fetch post")
	}
	if post.IsEmpty() {
		return apperror.New(apperror.CodeForbidden, "cannot access to posts")
	}
	if post.UserID != input.UserID {
		return apperror.New(apperror.CodeForbidden, "this post is not yours")
	}

	err = c.postRepository.Delete(ctx, post.ID)
	if err != nil {
		return apperror.New(apperror.CodeInternalServer, "failed to Delete post")
	}

	return nil
}
