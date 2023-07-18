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

	if in.AsgName == "" {
		return &pb.CUDResponse{
			Status: false,
			Data: &pb.Data{
				Code:    422,
				BgpName: []string{"Name cannot be blank"},
			},
		}, nil

	}

	asgParentGroup, err := json.Marshal(in.AsgParentGroup)
	if err != nil {
		return nil, err
	}

	assetSubGroup.AsgParentGroup = string(asgParentGroup)
	assetSubGroup.AsgName = in.AsgName

	if err = s.H.DB.Create(&assetSubGroup).Error; err != nil {
		return &pb.CUDResponse{
			Status: false,
			Data:   &pb.Data{Code: http.StatusInternalServerError, Message: err.Error()},
		}, nil
	}

	return &pb.CUDResponse{
		Status: true,
		Data: &pb.Data{
			Code:    200,
			Message: "Asset SubGroup created",
		},
	}, nil
}

func (s *Server) Update(ctx context.Context, in *pb.CreateUpdateRequest) (*pb.CUDResponse, error) {
	var assetSubGroup models.AssetSubGroup

	if result := s.H.DB.First(&assetSubGroup, in.Id); result.Error != nil {
		return &pb.CUDResponse{
			Status: false,
			Data: &pb.Data{
				Name:    "Not Found",
				Message: "Page not found",
				Code:    404,
			},
		}, nil
	}
	if in.AsgName == "" {
		return &pb.CUDResponse{
			Status: false,
			Data: &pb.Data{
				BgpName: []string{"Name cannot be blank."},
				Code:    422,
			},
		}, nil
	}

	asgParentGroup, err := json.Marshal(in.AsgParentGroup)
	if err != nil {
		return nil, err
	}
	assetSubGroup.AsgParentGroup = string(asgParentGroup)
	assetSubGroup.AsgName = in.AsgName

	if result := s.H.DB.Save(&assetSubGroup); result.Error != nil {
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
			Message: "Asset SubGroup updated",
			Code:    200,
		},
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
				Message: result.Error.Error(),
				Code:    http.StatusInternalServerError,
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
	// Prepare a variable to store the results
	var assetSubGroups []*pb.AssetSubGroup

	// Build the query based on the request parameters
	query := s.H.DB

	// Pagination
	if in.PerPage > 0 && in.Page > 0 {
		offset := (in.Page - 1) * in.PerPage
		query = query.Offset(int(offset)).Limit(int(in.PerPage))
	}

	// Sorting
	if in.Sort != "" {
		query = query.Order(in.Sort)
	}

	// Search
	if in.Search != "" {
		query = query.Where("asg_name LIKE ?", "%"+in.Search+"%")
	}

	// Retrieve the asset subgroups from the database
	results := query.Find(&assetSubGroups)

	if results.Error != nil {
		// Return an error response if there was an issue with the query
		return nil, results.Error
	}

	// Expand fields if requested
	// if len(in.Expand) > 0 {

	// }

	// Prepare the response
	response := &pb.ReadResponse{
		Success: true,
		Data: &pb.ReadResponseData{
			Items: assetSubGroups,
		},
	}

	return response, nil
	//return &pb.ReadResponse{}, nil
}
