BUILD_NAME=petapp

build:
	GOARCH=amd64 GOOS=darwin go build -o ${BUILD_NAME}-darwin main.go
 	GOARCH=amd64 GOOS=linux go build -o ${BUILD_NAME}-linux main.go
 	GOARCH=amd64 GOOS=window go build -o ${BUILD_NAME}-windows main.go

start:
	./${BUILD_NAME}

build_run: build start

clean:
	go clean
	rm ${BUILD_NAME}-darwin ${BUILD_NAME}-linux ${BUILD_NAME}-windows

# start postgres(docker-compose) and go server
docker-upd:
	docker-compose up -d
docker-down:
	docker-compose down
dev:
	docker-compose up -d
	go run main.go

run: main.go
	nodemon --exec "go run" main.go