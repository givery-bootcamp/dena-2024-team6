package usecases

import (
	"myapp/internal/entities"
	"myapp/internal/interfaces"
)

type CommentUsecase struct {
	repository interfaces.PostsRepository
}

func NewCommentUsecase(r interfaces.PostsRepository) *CommentUsecase {
	return &CommentUsecase{
		repository: r,
	}
}

func (u *CommentUsecase) Execute(userID int, postID int, body string) (*entities.Post, error) {
	return u.repository.CreateComment(userID, postID, body)
}
