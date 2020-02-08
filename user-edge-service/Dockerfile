#源镜像
FROM golang:1.13

#设置工作目录
WORKDIR $GOPATH/src

RUN mkdir user-edge-service

COPY user-edge-service $GOPATH/src/user-edge-service
COPY config.ini $GOPATH/src/user-edge-service


#切换工作目录
WORKDIR $GOPATH/src/user-edge-service
RUN mkdir log


#暴露端口
EXPOSE 9091

#最终运行docker的命令
ENTRYPOINT  ["./user-edge-service"]