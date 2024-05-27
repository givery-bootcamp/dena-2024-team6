package usecases

import (
	"myapp/internal/entities"
	"myapp/internal/interfaces"
)

type PostsUsecase struct {
	repository interfaces.PostsRepository
}

func NewPostsUsecase(r interfaces.PostsRepository) *PostsUsecase {
	return &PostsUsecase{
		repository: r,
	}
}

func (u *PostsUsecase) Execute() ([]*entities.Post, error) {
	return u.repository.Get()
}
