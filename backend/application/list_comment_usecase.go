package application

import (
	"context"
	"myapp/domain/apperror"
	"myapp/domain/model"
	"myapp/domain/repository"

	"github.com/samber/do"
)

type ListCommentsUsecase interface {
	Execute(ctx context.Context, input ListCommentsUsecaseInput) (ListCommentsUsecaseOutput, error)
}

type ListCommentsUsecaseInput struct {
	PostID int
}

type ListCommentsUsecaseOutput struct {
	Comments []model.Comment
}

func NewListCommentsUsecase(i *do.Injector) (ListCommentsUsecase, error) {
	postRepo := do.MustInvoke[repository.PostRepository](i)
	commentRepo := do.MustInvoke[repository.CommentRepository](i)
	return &listCommentsUsecaseInteractor{
		postRepository:    postRepo,
		commentRepository: commentRepo,
	}, nil
}

type listCommentsUsecaseInteractor struct {
	postRepository    repository.PostRepository
	commentRepository repository.CommentRepository
}

// Execute implements ListCommentsUsecase.
func (l *listCommentsUsecaseInteractor) Execute(ctx context.Context, input ListCommentsUsecaseInput) (ListCommentsUsecaseOutput, error) {
	post, err := l.postRepository.GetDetail(ctx, input.PostID)
	if err != nil {
		return ListCommentsUsecaseOutput{}, apperror.New(apperror.CodeInternalServer, "failed to fetch post")
	}
	if post.IsEmpty() {
		return ListCommentsUsecaseOutput{}, apperror.New(apperror.CodeForbidden, "cannot access to comments of this post id")
	}

	comments, err := l.commentRepository.List(ctx, post.ID)
	if err != nil {
		return ListCommentsUsecaseOutput{}, apperror.New(apperror.CodeInternalServer, "failed to fetch comments")
	}

	return ListCommentsUsecaseOutput{
		Comments: comments,
	}, nil
}
