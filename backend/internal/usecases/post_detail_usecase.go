package usecases

import (
	"myapp/internal/entities"
	"myapp/internal/interfaces"
)

type PostDetailUsecase struct {
	repository interfaces.PostsRepository
}

func NewPostDetailUsecase(r interfaces.PostsRepository) *PostDetailUsecase {
	return &PostDetailUsecase{
		repository: r,
	}
}

func (u *PostDetailUsecase) Execute(postID int) (*entities.Post, error) {
	return u.repository.Get(postID)
}
