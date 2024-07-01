package application

import (
	"context"
	"myapp/domain/apperror"
	"myapp/domain/model"
	"myapp/domain/repository"

	"github.com/samber/do"
)

type ListSpeedsUsecase interface {
	Execute(ctx context.Context) (ListSpeedsUsecaseOutput, error)
}

type ListSpeedsUsecaseOutput struct {
	Speeds []model.Speed
}

func NewListSpeedsUsecase(i *do.Injector) (ListSpeedsUsecase, error) {
	postRepository := do.MustInvoke[repository.PostRepository](i)
	commentRepository := do.MustInvoke[repository.CommentRepository](i)
	likeRepository := do.MustInvoke[repository.LikeRepository](i)

	return &listSpeedsUsecaseInteractor{
		postRepository:    postRepository,
		commentRepository: commentRepository,
		likeRepository:    likeRepository,
	}, nil
}

type listSpeedsUsecaseInteractor struct {
	postRepository    repository.PostRepository
	commentRepository repository.CommentRepository
	likeRepository    repository.LikeRepository
}

// Execute implements ListSpeedsUsecase.
func (l *listSpeedsUsecaseInteractor) Execute(ctx context.Context) (ListSpeedsUsecaseOutput, error) {
	posts, err := l.postRepository.List(ctx)
	if err != nil {
		return ListSpeedsUsecaseOutput{}, apperror.New(apperror.CodeInternalServer, "failed to get posts")
	}
	if len(posts) == 0 {
		return ListSpeedsUsecaseOutput{}, apperror.New(apperror.CodeNotFound, "there is no post")
	}

	var speeds []model.Speed
	for _, post := range posts {
		comments, err := l.commentRepository.List(ctx, post.ID)
		if err != nil {
			return ListSpeedsUsecaseOutput{}, apperror.New(apperror.CodeInternalServer, "failed to get comments")
		}

		likes, err := l.likeRepository.Get(ctx, post.ID)
		if err != nil {
			return ListSpeedsUsecaseOutput{}, apperror.New(apperror.CodeInternalServer, "failed to get likes")
		}

		speeds = append(speeds, model.Speed{
			PostID: post.ID,
			Speed:  (len(comments) * 35) + (likes * 3),
		})
	}

	return ListSpeedsUsecaseOutput{
		Speeds: speeds,
	}, nil
}
