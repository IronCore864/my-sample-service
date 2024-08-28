FROM golang:alpine AS build-env
WORKDIR $GOPATH/src/github.com/ironcore864/my-sample-service
COPY . .
RUN apk add git
RUN go get ./... && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o my-sample-service

FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/github.com/ironcore864/my-sample-service/my-sample-service /app/
CMD ["./my-sample-service"]
USER 1000
EXPOSE 8080/tcp
