package userusecase

import (
	"context"
	pb "github.com/ghn-rs/cloud-strife-user/proto/gen"
	"google.golang.org/grpc/codes"
)

func (uc *UseCase) PaginateUser(ctx context.Context, req *pb.PaginateUserRequest, resp *pb.UsersResponse) (err error) {

	users, total, err := uc.User.PaginateUsers(ctx, req.Page, req.Limit)
	if err != nil {
		resp.Base.Status = uint32(codes.Internal)
		resp.Base.Message = "Failed to paginate users"
		return
	}

	var pbUsers []*pb.User
	for _, user := range users {
		pbUsers = append(pbUsers, &pb.User{
			Id:          user.Id,
			Username:    user.Username,
			Email:       user.Email,
			Role:        user.Role,
			Status:      user.Status,
			LastLogin:   user.LastLogin.Unix(),
			CreatedDate: user.CreatedDate.Unix(),
			UpdatedDate: user.UpdatedDate.Unix(),
		})
	}

	resp.Base.Status = uint32(codes.OK)
	resp.Base.Message = "Users paginated successfully"
	resp.Users = pbUsers
	resp.TotalCount = int32(total)
	return
}

func (uc *UseCase) ListUsers(ctx context.Context, req *pb.EmptyRequest, resp *pb.UsersResponse) (err error) {

	users, err := uc.User.GetAllUsers(ctx)
	if err != nil {
		resp.Base.Status = uint32(codes.Internal)
		resp.Base.Message = "Failed to list users"
		return
	}

	var pbUsers []*pb.User
	for _, user := range users {
		pbUsers = append(pbUsers, &pb.User{
			Id:          user.Id,
			Username:    user.Username,
			Email:       user.Email,
			Role:        user.Role,
			Status:      user.Status,
			LastLogin:   user.LastLogin.Unix(),
			CreatedDate: user.CreatedDate.Unix(),
			UpdatedDate: user.UpdatedDate.Unix(),
		})
	}

	resp.Base.Status = uint32(codes.OK)
	resp.Base.Message = "Users listed successfully"
	resp.Users = pbUsers
	resp.TotalCount = int32(len(users))
	return
}
