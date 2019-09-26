FROM golang
MAINTAINER alphejangs@gmail.com

VOLUME /go/src/leveldb-server
WORKDIR /go/src/leveldb-server
COPY . .
ENV GOPROXY https://goproxy.io
RUN go install -v ./...
EXPOSE 9650
ENTRYPOINT ["leveldb-server"]