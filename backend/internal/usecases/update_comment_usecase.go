package usecases

import (
	"myapp/internal/entities"
	"myapp/internal/interfaces"
)

type UpdateCommentUsecase struct {
	repository interfaces.PostsRepository
}

func NewUpdateCommentUsecase(r interfaces.PostsRepository) *UpdateCommentUsecase {
	return &UpdateCommentUsecase{
		repository: r,
	}
}

func (u *UpdateCommentUsecase) Execute(userID int, commentID int, postID int, body string) (*entities.Post, error) {
	return u.repository.UpdateComment(userID, commentID, postID, body)
}
