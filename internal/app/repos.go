package app

import (
	"github.com/ghn-rs/cloud-strife-user/internal/repository/database"
)

type Repos struct {
	User *database.User
}

func NewRepos(dependencies *Dependencies) *Repos {

	return &Repos{
		User: database.NewUser(dependencies.DatabaseClient.Db),
	}
}
