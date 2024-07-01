package application

import (
	"context"
	"database/sql"
	"myapp/domain/apperror"
	"myapp/domain/repository"

	"github.com/samber/do"
)

type LikePostUsecase interface {
	Execute(ctx context.Context, input LikePostUsecaseInput) error
}

type LikePostUsecaseInput struct {
	PostID int
	Value  int
}

func NewLikePostUsecase(i *do.Injector) (LikePostUsecase, error) {
	postRepo := do.MustInvoke[repository.PostRepository](i)
	likeRepo := do.MustInvoke[repository.LikeRepository](i)
	return &LikePostUsecaseInteractor{
		postRepository: postRepo,
		likeRepository: likeRepo,
	}, nil
}

type LikePostUsecaseInteractor struct {
	postRepository repository.PostRepository
	likeRepository repository.LikeRepository
}

// Execute implements LikePostUsecase.
func (c *LikePostUsecaseInteractor) Execute(ctx context.Context, input LikePostUsecaseInput) error {
	post, _ := c.postRepository.GetDetail(ctx, input.PostID)
	if post.IsEmpty() {
		return apperror.New(apperror.CodeNotFound, "post is not found")
	}

	likes, err := c.likeRepository.Get(ctx, input.PostID)
	if err == sql.ErrNoRows {
		_, err = c.likeRepository.Create(ctx, input.PostID)
		likes = 0
		if err != nil {
			return apperror.New(apperror.CodeInternalServer, "failed to create likes record")
		}
	}
	if err != nil {
		return apperror.New(apperror.CodeInternalServer, "failed to fetch likes record")
	}
	likes += input.Value
	err = c.likeRepository.Update(ctx, input.PostID, likes)
	if err != nil {
		return apperror.New(apperror.CodeInternalServer, "failed to Update comment")
	}

	return nil
}
