build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/superryanguo/kick \
	  user_service/proto/user.proto
	docker build -t outlet_client -f ./outlet_client/Dockerfile ./outlet_client

run:
	docker run outlet_client

