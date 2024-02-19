FROM golang:latest AS builder
COPY . /github.com/andy-ahmedov/task_manager_client/
WORKDIR /github.com/andy-ahmedov/task_manager_client/
RUN go mod download

# Instructions for running without further assembly:
# RUN go build -o ./bin/client_init cmd/main.go
# CMD ["./bin/client_init"]
RUN GOOS=linux go build -o ./.bin/client_init cmd/main.go

FROM alpine:3.19
WORKDIR /root/
COPY --from=builder /github.com/andy-ahmedov/task_manager_client/.bin/client_init .
RUN apk add libc6-compat
# CMD ["ls", "-la"]
CMD ["./client_init", "getall"]