package application

import (
	"context"
	"myapp/domain/apperror"
	"myapp/domain/repository"

	"github.com/samber/do"
)

type UpdatePostUsecase interface {
	Execute(ctx context.Context, input UpdatePostUsecaseInput) error
}

type UpdatePostUsecaseInput struct {
	PostID int
	UserID int
	Title  string
	Body   string
}

func NewUpdatePostUsecase(i *do.Injector) (UpdatePostUsecase, error) {
	postRepo := do.MustInvoke[repository.PostRepository](i)
	return &updatePostUsecaseInteractor{
		postRepository: postRepo,
	}, nil
}

type updatePostUsecaseInteractor struct {
	postRepository repository.PostRepository
}

// Execute implements UpdatePostUsecase.
func (u *updatePostUsecaseInteractor) Execute(ctx context.Context, input UpdatePostUsecaseInput) error {
	post, err := u.postRepository.GetDetail(ctx, input.PostID)
	if err != nil {
		return apperror.New(apperror.CodeInternalServer, "投稿の取得に失敗しました")
	}
	if post.IsEmpty() {
		return apperror.New(apperror.CodeNotFound, "投稿が存在しません")
	}
	if post.UserID != input.UserID {
		return apperror.New(apperror.CodeForbidden, "この投稿を更新することはできません")
	}

	err = u.postRepository.Update(ctx, post.ID, input.Title, input.Body)
	if err != nil {
		return apperror.New(apperror.CodeInternalServer, "投稿の削除に失敗しました")
	}

	return nil

}
