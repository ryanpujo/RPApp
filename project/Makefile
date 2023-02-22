USER_BINARY=userApp


docker_run: user_binary
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

proto_user:
	cd ../user-service && protoc --go_out=user-proto --proto_path=proto proto/*.proto --go-grpc_out=user-proto