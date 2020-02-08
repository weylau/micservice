#源镜像
FROM golang:1.13

#设置工作目录
WORKDIR $GOPATH/src

RUN mkdir user-service

COPY user-service $GOPATH/src/user-service


#切换工作目录
WORKDIR $GOPATH/src/user-service

RUN mkdir log

#暴露端口
EXPOSE 9090

#最终运行docker的命令
ENTRYPOINT  ["nohup","./user-service",">log/log.log", "&"]