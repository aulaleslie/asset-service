package main

import (
	"fmt"
	"log"
	"net"

	"github.com/aulaleslie/asset-service/pkg/asset_group"
	assetGroupPb "github.com/aulaleslie/asset-service/pkg/asset_group/pb"
	"github.com/aulaleslie/asset-service/pkg/asset_sub_group"
	assetSubGroupPb "github.com/aulaleslie/asset-service/pkg/asset_sub_group/pb"
	"github.com/aulaleslie/asset-service/pkg/config"
	"github.com/aulaleslie/asset-service/pkg/db"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config:", err)
	}

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing", err)
	}

	fmt.Println("Asset service on", c.Port)

	ags := asset_group.Server{
		H: h,
	}

	asgs := asset_sub_group.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()

	assetGroupPb.RegisterAssetGroupServiceServer(grpcServer, &ags)
	assetSubGroupPb.RegisterAssetSubGroupServiceServer(grpcServer, &asgs)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve", err)
	}
}
