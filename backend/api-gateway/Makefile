gen-api-gateway:
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	pkg/driver/pb/driver.proto

	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	pkg/trip/pb/trip.proto

run-server:
	go run cmd/main.go