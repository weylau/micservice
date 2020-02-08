#!/bin/bash
rm -rf user-service
go build .
docker build -t registry.cn-shenzhen.aliyuncs.com/weylau/mic-user-service:latest -f Dockerfile .
docker push registry.cn-shenzhen.aliyuncs.com/weylau/mic-user-service:latest