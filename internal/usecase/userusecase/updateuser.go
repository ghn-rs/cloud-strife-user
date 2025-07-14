package userusecase

import (
	"context"
	"errors"
	"github.com/ghn-rs/cloud-strife-user/internal/model/entity"
	pb "github.com/ghn-rs/cloud-strife-user/proto/gen"
	"github.com/ghn-rs/corelib/src/helper"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
)

func (uc *UseCase) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest, resp *pb.UserResponse) (err error) {

	existingUser, err := uc.User.GetUserById(ctx, req.Id)
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

	payload := &entity.User{
		EssentialEntity: existingUser.EssentialEntity,
		Password:        existingUser.Password,
		LastLogin:       existingUser.LastLogin,
	}
	helper.SetIfNotNil(&payload.Email, req.Email)
	helper.SetIfNotNil(&payload.Username, req.Username)
	helper.SetIfNotNil(&payload.Role, req.Role)
	helper.SetIfNotNil(&payload.Status, req.Status)

	err = uc.User.UpdateUser(ctx, payload)
	if err != nil {
		resp.Base.Status = uint32(codes.Internal)
		resp.Base.Message = "Failed to update user"
		return
	}

	updatedUser, err := uc.User.GetUserById(ctx, req.Id)
	if err != nil {
		resp.Base.Status = uint32(codes.Internal)
		resp.Base.Message = "Failed to retrieve updated user"
		return
	}

	resp.Base.Status = uint32(codes.OK)
	resp.Base.Message = "User updated successfully"
	resp.User = &pb.User{
		Id:          updatedUser.Id,
		Username:    updatedUser.Username,
		Email:       updatedUser.Email,
		Role:        updatedUser.Role,
		Status:      updatedUser.Status,
		LastLogin:   updatedUser.LastLogin.Unix(),
		CreatedDate: updatedUser.CreatedDate.Unix(),
		UpdatedDate: updatedUser.UpdatedDate.Unix(),
	}
	return
}
