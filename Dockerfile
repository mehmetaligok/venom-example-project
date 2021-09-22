FROM golang:1.16.8-alpine as builder

WORKDIR /build

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download

COPY . /build

# Build grpc server
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o grpc-server ./src/cmd/grpc-server

# Build http server
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o http-server ./src/cmd/http-server

FROM alpine:3.12.3 as base

COPY --from=builder /build/grpc-server /grpc-server
COPY --from=builder /build/http-server /http-server
