package asset_group

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aulaleslie/asset-service/pkg/asset_group/pb"
	"github.com/aulaleslie/asset-service/pkg/db"
	"github.com/aulaleslie/asset-service/pkg/models"
)

type Server struct {
	pb.UnimplementedAssetGroupServiceServer
	H db.Handler
}

func (s *Server) Create(ctx context.Context, in *pb.CreateUpdateRequest) (*pb.CUDResponse, error) {
	var assetsGroup models.AssetsGroup

	agrBusinessGroup, err := json.Marshal(in.AgrBusinessGroup)
	if err != nil {
		return nil, err
	}
	assetsGroup.AgrBusinessGroup = string(agrBusinessGroup)

	if in.AgrGroupName == "" {
		return &pb.CUDResponse{
			Status: false,
			Data: &pb.Data{
				Code:    422,
				Message: "Group Name is required",
			},
		}, nil
	}
	assetsGroup.AgrGroupName = in.AgrGroupName

	if err = s.H.DB.Create(&assetsGroup).Error; err != nil {
		return &pb.CUDResponse{
			Status: false,
			Data: &pb.Data{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}, nil
	}

	return &pb.CUDResponse{
		Status: true,
		Data: &pb.Data{
			Code: http.StatusCreated,
		},
	}, nil
}

func (s *Server) Update(ctx context.Context, in *pb.CreateUpdateRequest) (*pb.CUDResponse, error) {
	var assetsGroup models.AssetsGroup

	agrBusinessGroup, err := json.Marshal(in.AgrBusinessGroup)
	if err != nil {
		return nil, err
	}
	assetsGroup.AgrBusinessGroup = string(agrBusinessGroup)

	if in.AgrGroupName == "" {
		return &pb.CUDResponse{}, nil
	}

	assetsGroup.AgrGroupName = in.AgrGroupName

	if res := s.H.DB.Save(&assetsGroup); res.Error != nil {
		return &pb.CUDResponse{
			Status: false,
			Data: &pb.Data{
				Code: http.StatusConflict,
			},
		}, nil
	}

	return &pb.CUDResponse{
		Status: true,
	}, nil
}

func (s *Server) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.CUDResponse, error) {
	var assetGroup models.AssetsGroup

	if result := s.H.DB.Find(&assetGroup, in.Id); result.Error != nil {
		return &pb.CUDResponse{
			Status: false,
			Data: &pb.Data{
				Name:    "Not Found",
				Code:    http.StatusNotFound,
				Message: "Page not found",
			},
		}, nil
	}

	if result := s.H.DB.Delete(&assetGroup); result.Error != nil {
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
	// TODO need to confirm Omar about BusinessGroup, SubGroup, AssetType, Value (lack information)
	return &pb.ReadResponse{}, nil
}
