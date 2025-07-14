package app

import (
	"github.com/ghn-rs/cloud-strife-user/internal/usecase/userusecase"
)

type UseCases struct {
	User userusecase.UseCase
}

func NewUseCases(repos *Repos) *UseCases {

	return &UseCases{
		User: userusecase.UseCase{
			User: repos.User,
		},
	}
}
