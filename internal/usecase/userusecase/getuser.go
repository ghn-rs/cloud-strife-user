package userusecase

import (
	"context"
	"errors"
	pb "github.com/ghn-rs/cloud-strife-user/proto/gen"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
)

func (uc *UseCase) GetUser(ctx context.Context, req *pb.GetUserRequest, resp *pb.UserResponse) (err error) {

	user, err := uc.User.GetUserById(ctx, req.Id)
	switch {
	case err == nil:
	case errors.Is(err, gorm.ErrRecordNotFound):
		resp.Base.Status = uint32(codes.NotFound)
		resp.Base.Message = "User not found"
		return
	default:
		resp.Base.Status = uint32(codes.Internal)
		resp.Base.Message = "Failed to retrieve user"
		return
	}

	resp.Base.Status = uint32(codes.OK)
	resp.Base.Message = "User retrieved successfully"
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
	return
}
