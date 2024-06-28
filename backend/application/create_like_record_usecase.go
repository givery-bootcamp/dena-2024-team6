package application

import (
	"context"
	"fmt"
	"myapp/domain/apperror"
	"myapp/domain/repository"

	"github.com/samber/do"
)

type CreateLikeRecordUsecase interface {
	Execute(ctx context.Context) error
}

func NewCreateLikeRecordUsecase(i *do.Injector) (CreateLikeRecordUsecase, error) {
	postRepo := do.MustInvoke[repository.PostRepository](i)
	likeRepo := do.MustInvoke[repository.LikeRepository](i)
	return &CreateLikeRecordUsecaseInteractor{
		postRepository: postRepo,
		likeRepository: likeRepo,
	}, nil
}

type CreateLikeRecordUsecaseInteractor struct {
	postRepository repository.PostRepository
	likeRepository repository.LikeRepository
}

// Execute implements CreateLikeRecordUsecase.
func (c *CreateLikeRecordUsecaseInteractor) Execute(ctx context.Context) error {
	posts, err := c.postRepository.List(ctx)
	if err != nil {
		return apperror.New(apperror.CodeInternalServer, "failed to get posts")
	}
	if len(posts) == 0 {
		return apperror.New(apperror.CodeNotFound, "there is no post")
	}

	for _, post := range posts {
		_, err := c.likeRepository.Create(ctx, post.ID)
		if err != nil {
			return apperror.New(apperror.CodeInternalServer, fmt.Sprintf("failed to create like record for post ID: %d", post.ID))
		}
	}
	return nil
}
