package userusecase

import "github.com/ghn-rs/cloud-strife-user/internal/repository/database"

type UseCase struct {
	User *database.User
}
