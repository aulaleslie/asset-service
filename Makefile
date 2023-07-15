proto-asset-group:
	protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative pkg/asset_group/pb/asset_group.proto

proto-asset-sub-group:
	protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative pkg/asset_sub_group/pb/asset_sub_group.proto

proto:
	$(MAKE) proto-asset-group && \
    $(MAKE) proto-asset-sub-group

server:
	go run cmd/main.go	