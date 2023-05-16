USER_BINARY=userApp
BROKER_BINARY=brokerApp


docker_run: user_binary broker_binary
	@echo "stop all running container"
	docker-compose down -v
	@echo "building container(when required) and start docker containers"
	docker-compose up --build 
	@echo "docker image is built and started"

docker_stop:
	docker-compose down

user_binary:
	@echo "building user-service binary"
	cd ../user-service && env GOOS=linux CGO_ENABLED=0 go build -o ${USER_BINARY} ./cmd
	@echo "user_binary is built and ready to be run"

broker_binary:
	@echo "building broker-service binary"
	cd ../broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd
	@echo "broker_binary is built and ready to be run"

proto_user:
	cd ../user-service && protoc --go_out=user-proto --proto_path=proto proto/*.proto --go-grpc_out=user-proto
	cd ../broker-service && protoc --go_out=user/user-proto --proto_path=user/proto user/proto/*.proto --go-grpc_out=user/user-proto

user_image: user_binary
	@echo "building user image"
	cd ../user-service && docker build -t ryanpujo/user-service . && docker push ryanpujo/user-service
	@echo "user image is built"

frontend_image:
	@echo "building frontend image"
	cd ../front-end && docker build -t ryanpujo/front-end . && docker push ryanpujo/front-end
	@echo "frontend image is built"

broker_image: broker_binary
	@echo "building broker image"
	cd ../broker-service && docker build -t ryanpujo/broker-service . && docker push ryanpujo/broker-service
	@echo "broker image is built"

docker_image: user_image broker_image

proto_product:
	cd ../product-service && protoc --go_out=product-proto --proto_path=proto proto/*.proto --go-grpc_out=product-proto
	cd ../broker-service && protoc --go_out=product/product-proto --proto_path=product/proto product/proto/*.proto --go-grpc_out=product/product-proto

test: user_test broker_test product_test

user_test:
	@echo "running test for user service"
	cd ../user-service && go test ./interface/repository ./interface/controller ./usecases/interactor --coverprofile=cover.out
	@echo "finished running all test"

broker_test:
	@echo "running test for broker service"
	cd ../broker-service && go test ./user/interface/controller --coverprofile=cover.out
	@echo "finished running all test"

product_test:
	@echo "run test for product service"
	cd ../product-service && go test ./internal/repository ./internal/service --coverprofile=cover.out
	@echo "finish running all test"

coverage_show:
	cd ../user-service && go tool cover -html=cover.out
	cd ../broker-service && go tool cover -html=cover.out
	cd ../product-service && go tool cover -html=cover.out