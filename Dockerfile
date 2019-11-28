FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/varandrew/gin-blog
COPY . $GOPATH/src/github.com/varandrew/gin-blog
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./gin-blog"]