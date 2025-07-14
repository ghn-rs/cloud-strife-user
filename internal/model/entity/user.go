package entity

import (
	"github.com/ghn-rs/corelib/src/entity"
	"time"
)

type User struct {
	LastLogin time.Time `json:"last_login"`
	entity.EssentialEntity
	Username string `gorm:"type:varchar(255);uniqueIndex;not null" json:"username"`
	Password string `gorm:"type:text;not null" json:"-"`
	Email    string `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Role     string `gorm:"not null;default:'user'" json:"role"`
	Status   int32  `gorm:"not null;default:0" json:"status"`
}

func (User) TableName() string {
	return "users"
}
