FROM golang:1.13.7 as base

# Set necessary environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . .

# Build the application
RUN go build -o main main.go

FROM busybox as app

WORKDIR /service
COPY --from=base /build/main .

# Export necessary port
EXPOSE 8080

ENTRYPOINT ["/service/main"]
