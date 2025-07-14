package userusecase

import (
	"context"
	pb "github.com/ghn-rs/cloud-strife-user/proto/gen"
	coreproto "github.com/ghn-rs/corelib/proto/gen"
	"google.golang.org/grpc/codes"
)

func (uc *UseCase) HealthCheck(ctx context.Context, req *pb.EmptyRequest, resp *coreproto.GenericResponse) (err error) {
	resp.Base.Status = uint32(codes.OK)
	resp.Base.Message = "User service is healthy"
	return
}
