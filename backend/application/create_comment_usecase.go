package application

import (
	"context"
	"myapp/domain/apperror"
	"myapp/domain/model"
	"myapp/domain/repository"

	"github.com/samber/do"
)

type CreateCommentUsecase interface {
	Execute(ctx context.Context, input CreateCommentUsecaseInput) (CreateCommentUsecaseOutput, error)
}

type CreateCommentUsecaseInput struct {
	PostID int
	UserID int
	Body   string
}

type CreateCommentUsecaseOutput struct {
	Comment model.Comment
}

func NewCreateCommentUsecase(i *do.Injector) (CreateCommentUsecase, error) {
	postRepo := do.MustInvoke[repository.PostRepository](i)
	commentRepo := do.MustInvoke[repository.CommentRepository](i)
	return &createCommentUsecaseInteractor{
		postRepository:    postRepo,
		commentRepository: commentRepo,
	}, nil
}

type createCommentUsecaseInteractor struct {
	postRepository    repository.PostRepository
	commentRepository repository.CommentRepository
}

// Execute implements CreateCommentUsecase.
func (c *createCommentUsecaseInteractor) Execute(ctx context.Context, input CreateCommentUsecaseInput) (CreateCommentUsecaseOutput, error) {
	post, err := c.postRepository.GetDetail(ctx, input.PostID)
	if err != nil {
		return CreateCommentUsecaseOutput{}, apperror.New(apperror.CodeInternalServer, "failed to fetch post")
	}
	if post.IsEmpty() {
		return CreateCommentUsecaseOutput{}, apperror.New(apperror.CodeForbidden, "cannot access to comments of this post id")
	}

	comment, err := c.commentRepository.Create(ctx, post.ID, input.UserID, input.Body)
	if err != nil {
		return CreateCommentUsecaseOutput{}, apperror.New(apperror.CodeInternalServer, "failed to create comment")
	}

	return CreateCommentUsecaseOutput{
		Comment: comment,
	}, nil
}
