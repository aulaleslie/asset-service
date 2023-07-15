package asset_sub_group

import (
	"context"

	"github.com/aulaleslie/asset-service/pkg/asset_sub_group/pb"
	"github.com/aulaleslie/asset-service/pkg/db"
)

type Server struct {
	pb.UnimplementedAssetSubGroupServiceServer
	H db.Handler
}

func (s *Server) Create(ctx context.Context, in *pb.CreateUpdateRequest) (*pb.CUDResponse, error) {
	return &pb.CUDResponse{}, nil
}

func (s *Server) Update(ctx context.Context, in *pb.CreateUpdateRequest) (*pb.CUDResponse, error) {
	return &pb.CUDResponse{}, nil
}

func (s *Server) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.CUDResponse, error) {
	return &pb.CUDResponse{}, nil
}

func (s *Server) Read(ctx context.Context, in *pb.ReadRequest) (*pb.ReadResponse, error) {
	return &pb.ReadResponse{}, nil
}
