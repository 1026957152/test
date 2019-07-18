FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o main .
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
CMD ["/app/main"]



# docker build --build-arg  http_proxy=http://192.168.10.101:1089 --build-arg https_proxy=http://192.168.10.101:1089 --network=host -t docker.yulinmei.cn/my-go-app .