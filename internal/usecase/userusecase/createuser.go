package userusecase

import (
	"context"
	"errors"
	"github.com/ghn-rs/cloud-strife-user/internal/model/entity"
	pb "github.com/ghn-rs/cloud-strife-user/proto/gen"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
	"time"
)

func (uc *UseCase) CreateUser(ctx context.Context, req *pb.CreateUserRequest, resp *pb.UserResponse) (err error) {
	payload := &entity.User{
		Username:  req.Username,
		Password:  req.Password,
		Email:     req.Email,
		Role:      req.Role,
		Status:    req.Status,
		LastLogin: time.Now(),
	}

	id, err := uc.User.InsertUser(ctx, payload)
	if err != nil {
		resp.Base.Status = uint32(codes.Internal)
		resp.Base.Message = "Failed to create user"
		return
	}

	user, err := uc.User.GetUserById(ctx, id)
	switch {
	case err == nil:
	case errors.Is(err, gorm.ErrRecordNotFound):
		resp.Base.Status = uint32(codes.NotFound)
		resp.Base.Message = "User not found after creation"
		return
	default:
		resp.Base.Status = uint32(codes.Internal)
		resp.Base.Message = "Failed to retrieve user after creation"
		return
	}

	resp.Base.Status = uint32(codes.OK)
	resp.Base.Message = "User created successfully"
	resp.User = &pb.User{
		Id:          user.Id,
		Username:    user.Username,
		Email:       user.Email,
		Role:        user.Role,
		Status:      user.Status,
		LastLogin:   user.LastLogin.Unix(),
		CreatedDate: user.CreatedDate.Unix(),
		UpdatedDate: user.UpdatedDate.Unix(),
	}
	return nil
}
