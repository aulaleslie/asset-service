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

	if in.AgrGroupName == "" {
		return &pb.CUDResponse{
			Status: false,
			Data: &pb.Data{
				Code:    422,
				BgpName: []string{"Name cannot be blank"},
			},
		}, nil
	}

	agrBusinessGroup, err := json.Marshal(in.AgrBusinessGroup)
	if err != nil {
		return nil, err
	}
	assetsGroup.AgrBusinessGroup = string(agrBusinessGroup)
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
			Code:    http.StatusCreated,
			Message: "Asset Group Created successfully",
		},
	}, nil
}

func (s *Server) Update(ctx context.Context, in *pb.CreateUpdateRequest) (*pb.CUDResponse, error) {
	var assetGroup models.AssetsGroup

	if result := s.H.DB.First(&assetGroup, in.Id); result.Error != nil {
		return &pb.CUDResponse{
			Status: false,
			Data: &pb.Data{
				Name:    "Not Found",
				Message: "Page not found",
				Code:    404,
			},
		}, nil
	}

	if in.AgrGroupName == "" {
		return &pb.CUDResponse{
			Status: false,
			Data: &pb.Data{
				BgpName: []string{"Name cannot be blank."},
				Code:    422,
			},
		}, nil
	}
	agrBusinessGroup, err := json.Marshal(in.AgrBusinessGroup)
	if err != nil {
		return nil, err
	}
	assetGroup.AgrBusinessGroup = string(agrBusinessGroup)
	assetGroup.AgrGroupName = in.AgrGroupName

	if result := s.H.DB.Save(&assetGroup); result.Error != nil {
		return &pb.CUDResponse{
			Status: false,
			Data: &pb.Data{
				Message: result.Error.Error(),
				Code:    500,
			},
		}, nil
	}

	return &pb.CUDResponse{
		Status: true,
		Data: &pb.Data{
			Message: "Asset group updated",
			Code:    200,
		},
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
				Message: result.Error.Error(),
				Code:    http.StatusInternalServerError,
			},
		}, nil
	}

	return &pb.CUDResponse{
		Status: true,
		Data: &pb.Data{
			Message: "Group Successfully Deleted",
			Code:    200,
		},
	}, nil
}

func (s *Server) Read(ctx context.Context, in *pb.ReadRequest) (*pb.ReadResponse, error) {
	// TODO need to confirm Omar about BusinessGroup, SubGroup, AssetType, Value (lack information)
	return &pb.ReadResponse{}, nil
}
