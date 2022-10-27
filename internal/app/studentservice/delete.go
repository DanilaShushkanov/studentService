package studentservice

import (
	"context"
	api "github.com/danilashushkanov/student-service/pkg/studentServiceApi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) DeleteStudent(ctx context.Context, req *api.GetStudentRequest) (*api.SimpleResponse, error) {
	if err := validateStudentIDRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	s.logger.Info("delete student")
	err := s.studentRepository.Delete(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, "error delete student")
	}

	return &api.SimpleResponse{}, nil
}
