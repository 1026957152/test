FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o main .
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
CMD ["/app/main"]



