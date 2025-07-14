package userusecase

import (
	"context"
	"errors"
	pb "github.com/ghn-rs/cloud-strife-user/proto/gen"
	coreproto "github.com/ghn-rs/corelib/proto/gen"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
)

func (uc *UseCase) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest, resp *coreproto.GenericResponse) (err error) {

	_, err = uc.User.GetUserById(ctx, req.Id)
	switch {
	case err == nil:
	case errors.Is(err, gorm.ErrRecordNotFound):
		resp.Base.Status = uint32(codes.NotFound)
		resp.Base.Message = "User not found"
	default:
		resp.Base.Status = uint32(codes.Internal)
		resp.Base.Message = "Failed to retrieve user"
		return
	}

	err = uc.User.DeleteUser(ctx, req.Id)
	if err != nil {
		resp.Base.Status = uint32(codes.Internal)
		resp.Base.Message = "Failed to delete user"
		return
	}

	resp.Base.Status = uint32(codes.OK)
	resp.Base.Message = "User deleted successfully"
	return
}
