package usecases

import (
	"myapp/internal/entities"
	"myapp/internal/interfaces"
)

type PostUsecase struct {
	repository interfaces.PostsRepository
}

func NewPostUsecase(r interfaces.PostsRepository) *PostUsecase {
	return &PostUsecase{
		repository: r,
	}
}

func (u *PostUsecase) Execute(userID int, title string, body string) (*entities.Post, error) {
	return u.repository.Create(userID, title, body)
}
