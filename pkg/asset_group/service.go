package asset_group

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/aulaleslie/asset-service/pkg/asset_group/pb"
	"github.com/aulaleslie/asset-service/pkg/db"
	"github.com/aulaleslie/asset-service/pkg/models"
)

type Server struct {
	pb.UnimplementedAssetGroupServiceServer
	H db.Handler
}

func (s *Server) Create(_ context.Context, in *pb.CreateUpdateRequest) (*pb.CUDResponse, error) {
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

	agrBusinessGroup := strings.Join(in.AgrBusinessGroup, ",")
	assetsGroup.AgrBusinessGroup = agrBusinessGroup
	assetsGroup.AgrGroupName = in.AgrGroupName

	if err := s.H.DB.Create(&assetsGroup).Error; err != nil {
		return &pb.CUDResponse{
			Status: false,
			Data: &pb.Data{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}, nil
	}

	if len(in.SubGroup) > 0 {
		for _, subGroupId := range in.SubGroup {
			var subGroup models.AssetSubGroup

			err := s.H.DB.First(&subGroup, "asg_id = ?", subGroupId).Error
			if err != nil {
				return &pb.CUDResponse{
					Status: false,
					Data: &pb.Data{
						Code:    http.StatusInternalServerError,
						Message: err.Error(),
					},
				}, nil
			}

			var asgParentGroups []string
			if subGroup.AsgParentGroup != "" {
				asgParentGroups = strings.Split(subGroup.AsgParentGroup, ",")
				asgParentGroups = append(asgParentGroups, subGroupId)
			} else {
				asgParentGroups = append(asgParentGroups, subGroupId)
			}

			subGroup.AsgParentGroup = strings.Join(asgParentGroups, ",")

			if err := s.H.DB.Save(&subGroup).Error; err != nil {
				return &pb.CUDResponse{
					Status: false,
					Data: &pb.Data{
						Code:    http.StatusInternalServerError,
						Message: err.Error(),
					},
				}, nil
			}
		}
	}

	return &pb.CUDResponse{
		Status: true,
		Data: &pb.Data{
			Code:    http.StatusCreated,
			Message: "Asset Group Created successfully",
		},
	}, nil
}

func (s *Server) Update(_ context.Context, in *pb.CreateUpdateRequest) (*pb.CUDResponse, error) {
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

	agrBusinessGroup := strings.Join(in.AgrBusinessGroup, ",")
	assetGroup.AgrBusinessGroup = agrBusinessGroup
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

	if len(in.SubGroup) > 0 {
		for _, subGroupId := range in.SubGroup {
			var subGroup models.AssetSubGroup

			err := s.H.DB.First(&subGroup, "asg_id = ?", subGroupId).Error
			if err != nil {
				return &pb.CUDResponse{
					Status: false,
					Data: &pb.Data{
						Code:    http.StatusInternalServerError,
						Message: err.Error(),
					},
				}, nil
			}

			var asgParentGroups []string
			if subGroup.AsgParentGroup != "" {
				asgParentGroups = strings.Split(subGroup.AsgParentGroup, ",")
				asgParentGroups = append(asgParentGroups, subGroupId)
			} else {
				asgParentGroups = append(asgParentGroups, subGroupId)
			}

			subGroup.AsgParentGroup = strings.Join(asgParentGroups, ",")

			if err := s.H.DB.Save(&subGroup).Error; err != nil {
				return &pb.CUDResponse{
					Status: false,
					Data: &pb.Data{
						Code:    http.StatusInternalServerError,
						Message: err.Error(),
					},
				}, nil
			}
		}
	}

	return &pb.CUDResponse{
		Status: true,
		Data: &pb.Data{
			Message: "Asset group updated",
			Code:    200,
		},
	}, nil
}

func (s *Server) Delete(_ context.Context, in *pb.DeleteRequest) (*pb.CUDResponse, error) {
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

func (s *Server) Read(_ context.Context, in *pb.ReadRequest) (*pb.ReadResponse, error) {
	var assetGroups []*models.AssetsGroup

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
		query = query.Where("agr_group_name LIKE ?", "%"+in.Search+"%")
	}

	// Retrieve the asset subgroups from the database
	err := query.Find(&assetGroups).Error
	if err != nil {
		// Return an error response if there was an issue with the query
		return nil, err
	}

	var items []*pb.AssetGroup
	for _, result := range assetGroups {
		var businessGroups []*models.BusinessGroup
		var subGroups []*models.AssetSubGroup

		businessGroupIds := strings.Split(result.AgrBusinessGroup, ",")
		err := s.H.DB.Where("bgp_id IN (?)", businessGroupIds).Find(&businessGroups).Error
		if err != nil {

		}

		err = s.H.DB.Where("asg_parent_group LIKE ?", "%"+fmt.Sprint(result.AgrID)+",%").Find(&subGroups).Error
		if err != nil {

		}

		var businessGroupResponses []*pb.BusinessGroup
		for _, businessGroup := range businessGroups {
			businessGroupResponses = append(businessGroupResponses, &pb.BusinessGroup{
				BgpId:           fmt.Sprint(businessGroup.BgpID),
				BgpLabel:        businessGroup.BgpLabel,
				BgpName:         businessGroup.BgpName,
				BgpOrganization: fmt.Sprint(businessGroup.BgpOrganization),
			})
		}

		var subGroupResponses []*pb.SubGroupName
		for _, subGroup := range subGroups {
			subGroupResponses = append(subGroupResponses, &pb.SubGroupName{
				AsgId:   fmt.Sprint(subGroup.AsgID),
				AsgName: subGroup.AsgName,
			})
		}

		item := &pb.AssetGroup{
			AgrId:            fmt.Sprint(result.AgrID),
			AgrGroupName:     result.AgrGroupName,
			AgrBusinessGroup: result.AgrBusinessGroup,
			AgrOrganization:  fmt.Sprint(result.AgrOrganization),
			BusinessGroups:   businessGroupResponses,
			SubGroupsNames:   subGroupResponses,
			AssetTypes:       nil,
			Values:           nil,
		}
		items = append(items, item)
	}

	return &pb.ReadResponse{
		Success: true,
		Data: &pb.ReadResponseData{
			Items:  items,
			XLinks: nil,
			XMeta:  nil,
		},
	}, nil
}
