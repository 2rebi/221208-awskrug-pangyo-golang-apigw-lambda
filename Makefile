
PROJECT_DIR := $(shell pwd)

install:
	go mod download all

build:
	cd $(PROJECT_DIR)/1_helloworld && \
        GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o main main.go && \
        zip function.zip main && rm main
	cd $(PROJECT_DIR)/2_apigw_lambda && \
        GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o main main.go && \
        zip function.zip main && rm main
	cd $(PROJECT_DIR)/3_apigw_lambda_route && \
        GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o main main.go && \
        zip function.zip main && rm main
	cd $(PROJECT_DIR)/4_case_http_net && \
        GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o main main.go && \
        zip function.zip main && rm main
	cd $(PROJECT_DIR)/5_case_echo_framework && \
        GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o main main.go && \
        zip function.zip main && rm main

local-clean:
	-rm -r $(PROJECT_DIR)/out

local-build: local-clean build
	mkdir $(PROJECT_DIR)/out
	mv $(PROJECT_DIR)/1_helloworld/function.zip $(PROJECT_DIR)/out/1_helloworld.zip
	mv $(PROJECT_DIR)/2_apigw_lambda/function.zip $(PROJECT_DIR)/out/2_apigw_lambda.zip
	mv $(PROJECT_DIR)/3_apigw_lambda_route/function.zip $(PROJECT_DIR)/out/3_apigw_lambda_route.zip
	mv $(PROJECT_DIR)/4_case_http_net/function.zip $(PROJECT_DIR)/out/4_case_http_net.zip
	mv $(PROJECT_DIR)/5_case_echo_framework/function.zip $(PROJECT_DIR)/out/5_case_echo_framework.zip