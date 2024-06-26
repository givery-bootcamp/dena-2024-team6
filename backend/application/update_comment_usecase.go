package application

import (
	"context"
	"myapp/domain/apperror"
	"myapp/domain/repository"

	"github.com/samber/do"
)

type UpdateCommentUsecase interface {
	Execute(ctx context.Context, input UpdateCommentUsecaseInput) error
}

type UpdateCommentUsecaseInput struct {
	PostID    int
	UserID    int
	CommentID int
	Body      string
}

func NewUpdateCommentUsecase(i *do.Injector) (UpdateCommentUsecase, error) {
	postRepo := do.MustInvoke[repository.PostRepository](i)
	commentRepo := do.MustInvoke[repository.CommentRepository](i)
	return &UpdateCommentUsecaseInteractor{
		postRepository:    postRepo,
		commentRepository: commentRepo,
	}, nil
}

type UpdateCommentUsecaseInteractor struct {
	postRepository    repository.PostRepository
	commentRepository repository.CommentRepository
}

// Execute implements UpdateCommentUsecase.
func (c *UpdateCommentUsecaseInteractor) Execute(ctx context.Context, input UpdateCommentUsecaseInput) error {
	post, err := c.postRepository.GetDetail(ctx, input.PostID)
	if err != nil {
		return apperror.New(apperror.CodeInternalServer, "failed to fetch post")
	}
	if post.IsEmpty() {
		return apperror.New(apperror.CodeForbidden, "cannot access to comments of this post id")
	}

	err = c.commentRepository.Update(ctx, post.ID, input.UserID, input.CommentID, input.Body)
	if err != nil {
		return apperror.New(apperror.CodeInternalServer, "failed to Update comment")
	}

	return nil
}
