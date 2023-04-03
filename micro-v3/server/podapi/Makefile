.PHONY: proto
proto:
	docker run --rm -v ${pwd}:/app -w /app gaotingwang/protoc:v3 -I ./ --go_out=./ --micro_out=./ ./proto/podApi/podApi.proto