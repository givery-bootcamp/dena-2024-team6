package application

import (
	"context"
	"myapp/domain/apperror"
	"myapp/domain/repository"

	"github.com/samber/do"
)

type DeleteCommentUsecase interface {
	Execute(ctx context.Context, input DeleteCommentUsecaseInput) error
}

type DeleteCommentUsecaseInput struct {
	CommentID int
}

func NewDeleteCommentUsecase(i *do.Injector) (DeleteCommentUsecase, error) {
	commentRepo := do.MustInvoke[repository.CommentRepository](i)
	return &DeleteCommentUsecaseInteractor{
		commentRepository: commentRepo,
	}, nil
}

type DeleteCommentUsecaseInteractor struct {
	commentRepository repository.CommentRepository
}

// Execute implements DeleteCommentUsecase.
func (c *DeleteCommentUsecaseInteractor) Execute(ctx context.Context, input DeleteCommentUsecaseInput) error {
	comment, err := c.commentRepository.GetByID(ctx, input.CommentID)
	if err != nil {
		return apperror.New(apperror.CodeInternalServer, "failed to fetch comment")
	}
	if comment.IsEmpty() {
		return apperror.New(apperror.CodeForbidden, "cannot access to comments")
	}

	err = c.commentRepository.Delete(ctx, comment.ID)
	if err != nil {
		return apperror.New(apperror.CodeInternalServer, "failed to Delete comment")
	}

	return nil
}
