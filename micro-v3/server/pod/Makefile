
.PHONY: proto
proto:
	docker run --rm -v ${pwd}:/app -w /app gaotingwang/protoc:v3 -I ./ --go_out=./ --micro_out=./ ./proto/pod/pod.proto

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o pod *.go

.PHONY: docker
docker:
	docker build -t pod:latest .

.PHONY: docker-run
docker-run:
	docker run -p 8081:8081 -p 9091:9091 -p 9191:9191 -v ../../docker/data/log:/micro.log pod