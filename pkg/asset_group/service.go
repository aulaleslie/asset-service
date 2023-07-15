package asset_group

import (
	"context"

	"github.com/aulaleslie/asset-service/pkg/asset_group/pb"
	"github.com/aulaleslie/asset-service/pkg/db"
)

type Server struct {
	pb.UnimplementedAssetGroupServiceServer
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

// func (s *Server) CreateAsset(ctx context.Context, req *pb.CreateAssetRequest) (*pb.CreateAssetResponse, error) {
// 	var asset models.Asset

// 	asset.Name = req.Name
// 	asset.Stock = req.Stock
// 	asset.Price = req.Price

// 	if result := s.H.DB.Create(&asset); result.Error != nil {
// 		return &pb.CreateAssetResponse{
// 			Status: http.StatusConflict,
// 			Error:  result.Error.Error(),
// 		}, nil
// 	}

// 	return &pb.CreateAssetResponse{
// 		Status: http.StatusCreated,
// 		Id:     asset.Id,
// 	}, nil
// }

// func (s *Server) FindOne(ctx context.Context, req *pb.FindOneRequest) (*pb.FindOneResponse, error) {
// 	var asset models.Asset

// 	if result := s.H.DB.First(&asset, asset.Id); result.Error != nil {
// 		return &pb.FindOneResponse{
// 			Status: http.StatusNotFound,
// 			Error:  result.Error.Error(),
// 		}, nil
// 	}

// 	data := &pb.FindOneData{
// 		Id:    asset.Id,
// 		Name:  asset.Name,
// 		Stock: asset.Stock,
// 		Price: asset.Price,
// 	}

// 	return &pb.FindOneResponse{
// 		Status: http.StatusOK,
// 		Data:   data,
// 	}, nil
// }

// func (s *Server) DecreaseStock(ctx context.Context, req *pb.DecreaseStockRequest) (*pb.DecreaseStockResponse, error) {
// 	var asset models.Asset

// 	if result := s.H.DB.First(&asset, req.Id); result.Error != nil {
// 		return &pb.DecreaseStockResponse{
// 			Status: http.StatusNotFound,
// 			Error:  result.Error.Error(),
// 		}, nil
// 	}

// 	if asset.Stock <= 0 {
// 		return &pb.DecreaseStockResponse{
// 			Status: http.StatusConflict,
// 			Error:  "Stock too low",
// 		}, nil
// 	}

// 	var log models.StockDecreaseLog

// 	if result := s.H.DB.Where(&models.StockDecreaseLog{OrderId: req.OrderId}).First(&log); result.Error != nil {
// 		return &pb.DecreaseStockResponse{
// 			Status: http.StatusConflict,
// 			Error:  "Stock already decrease",
// 		}, nil
// 	}

// 	asset.Stock = asset.Stock - 1

// 	s.H.DB.Save(&asset)

// 	log.OrderId = req.OrderId
// 	log.AssetRefer = asset.Id

// 	s.H.DB.Create(&log)

// 	return &pb.DecreaseStockResponse{
// 		Status: http.StatusOK,
// 	}, nil
// }
