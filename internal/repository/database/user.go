package database

import (
	"context"
	"github.com/ghn-rs/cloud-strife-user/internal/model/entity"
	"gorm.io/gorm"
	"time"
)

type User struct {
	DB *gorm.DB
}

func NewUser(DB *gorm.DB) *User {
	return &User{
		DB: DB,
	}
}

func (db *User) InsertUser(ctx context.Context, user *entity.User) (count int32, err error) {
	err = db.DB.WithContext(ctx).Create(&user).Error
	if err != nil {
		return 0, err
	}

	return user.Id, nil
}

func (db *User) GetUserById(ctx context.Context, id int32) (*entity.User, error) {
	var user entity.User
	err := db.DB.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		return &user, err
	}

	return &user, nil
}

func (db *User) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	var user entity.User
	err := db.DB.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if err != nil {
		return &user, err
	}

	return &user, nil
}

func (db *User) UpdateUser(ctx context.Context, user *entity.User) error {
	err := db.DB.WithContext(ctx).Where("id = ?", user.Id).Updates(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (db *User) UpdateUserLastLogin(ctx context.Context, id int32) error {
	err := db.DB.WithContext(ctx).Model(&entity.User{}).Where("id = ?", id).Update("last_login", time.Now()).Error
	if err != nil {
		return err
	}

	return nil
}

func (db *User) DeleteUser(ctx context.Context, id int32) error {
	var user entity.User
	user.Id = id
	err := db.DB.WithContext(ctx).Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (db *User) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
	err := db.DB.WithContext(ctx).Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

func (db *User) PaginateUsers(ctx context.Context, limit, offset int32) ([]entity.User, int64, error) {
	var users []entity.User
	var total int64

	err := db.DB.WithContext(ctx).Model(&entity.User{}).Count(&total).Error
	if err != nil {
		return users, 0, err
	}

	err = db.DB.WithContext(ctx).Limit(int(limit)).Offset(int(offset)).Find(&users).Error
	if err != nil {
		return users, total, err
	}

	return users, total, nil
}
