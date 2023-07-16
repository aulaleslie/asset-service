package asset_sub_group

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aulaleslie/asset-service/pkg/asset_sub_group/pb"
	"github.com/aulaleslie/asset-service/pkg/db"
	"github.com/aulaleslie/asset-service/pkg/models"
)

type Server struct {
	pb.UnimplementedAssetSubGroupServiceServer
	H db.Handler
}

func (s *Server) Create(ctx context.Context, in *pb.CreateUpdateRequest) (*pb.CUDResponse, error) {
	var assetSubGroup models.AssetSubGroup

	asgParentGroup, err := json.Marshal(in.AsgParentGroup)
	if err != nil {
		return nil, err
	}

	assetSubGroup.AsgParentGroup = string(asgParentGroup)

	if in.AsgName == "" {
		return &pb.CUDResponse{
			Status: false,
			Data:   &pb.Data{Code: 422, Message: "Group Name is required"},
		}, nil

	}
	assetSubGroup.AsgName = in.AsgName

	if result := s.H.DB.Create(&assetSubGroup); result.Error != nil {
		return &pb.CUDResponse{
			Status: false,
			Data:   &pb.Data{Code: http.StatusInternalServerError, Message: result.Error.Error()},
		}, nil
	}
	return &pb.CUDResponse{}, nil
}

func (s *Server) Update(ctx context.Context, in *pb.CreateUpdateRequest) (*pb.CUDResponse, error) {
	var assetSubGroup models.AssetSubGroup

	asgParentGroup, err := json.Marshal(&assetSubGroup)
	if err != nil {
		return nil, err
	}
	assetSubGroup.AsgParentGroup = string(asgParentGroup)

	if result := s.H.DB.Find(&assetSubGroup, in.Id); result.Error != nil {
		return &pb.CUDResponse{
			Status: false,
			Data:   &pb.Data{Code: http.StatusNotFound, Message: "id not found"},
		}, nil
	}

	assetSubGroup.AsgName = in.AsgName
	assetSubGroup.AsgParentGroup = string(asgParentGroup)

	return &pb.CUDResponse{
		Status: false,
		Data:   &pb.Data{Code: http.StatusOK, Message: "Update success"},
	}, nil
}

func (s *Server) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.CUDResponse, error) {
	var assetSubGroup models.AssetSubGroup

	if result := s.H.DB.Find(&assetSubGroup, in.Id); result.Error != nil {
		return &pb.CUDResponse{
			Status: false,
			Data: &pb.Data{
				Name:    "Not Found",
				Code:    http.StatusNotFound,
				Message: "Page not found",
			},
		}, nil
	}

	if result := s.H.DB.Delete(&assetSubGroup); result.Error != nil {
		return &pb.CUDResponse{
			Status: false,
			Data: &pb.Data{
				Name:    "Unauthorized",
				Message: "Your request was made with invalid credentials.",
				Code:    401,
			},
		}, nil
	}

	return &pb.CUDResponse{
		Status: true,
		Data: &pb.Data{
			Code:    200,
			Message: "Group Successfully Deleted",
		},
	}, nil
}

func (s *Server) Read(ctx context.Context, in *pb.ReadRequest) (*pb.ReadResponse, error) {
	// TODO create query. need information about AssetsGroup map query.
	return &pb.ReadResponse{}, nil
}
