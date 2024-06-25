package usecases

import (
	"myapp/internal/interfaces"
)

type DeleteCommentUsecase struct {
	repository interfaces.PostsRepository
}

func NewDeleteCommentUsecase(r interfaces.PostsRepository) *DeleteCommentUsecase {
	return &DeleteCommentUsecase{
		repository: r,
	}
}

func (u *DeleteCommentUsecase) Execute(userID int, commentID int) error {
	return u.repository.DeleteComment(userID, commentID)
}
