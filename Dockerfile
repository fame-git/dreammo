#
# Build go project
#
FROM golang:1.22.5-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go get -u github.com/swaggo/swag/cmd/swag && go install github.com/swaggo/swag/cmd/swag@latest && swag init

RUN go build -o myapp .

#
# Runtime container
#
FROM alpine:latest  

WORKDIR /app

COPY --from=builder /app/myapp .

COPY view ./view

EXPOSE 8000

ENTRYPOINT ["./myapp"]