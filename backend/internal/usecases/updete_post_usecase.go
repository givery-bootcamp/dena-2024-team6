package usecases

import (
	"myapp/internal/entities"
	"myapp/internal/interfaces"
)

type UpdatePostUsecase struct {
	repository interfaces.PostsRepository
}

func NewUpdatePostUsecase(r interfaces.PostsRepository) *UpdatePostUsecase {
	return &UpdatePostUsecase{
		repository: r,
	}
}

func (u *UpdatePostUsecase) Execute(userID int, postID int, title string, body string) (*entities.Post, error) {
	return u.repository.Update(userID, postID, title, body)
}
