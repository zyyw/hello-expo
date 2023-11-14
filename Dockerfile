FROM golang:1.21.4 as builder
ARG GIT_COMMIT
ARG GOPROXY

WORKDIR /src

# Copy in the go src
COPY ./ ./

RUN go mod download

# Build
RUN GO111MODULE=on CGO_ENABLED=0 go build -ldflags "-X main.GitCommit=$GIT_COMMIT" -a -o run-hello-expo

# Copy the binary, config, and static files into a thin image
FROM alpine:3.15

WORKDIR /app
COPY --from=builder /src/run-hello-expo .
COPY --from=builder /src/config/conf.yaml ./config/
COPY --from=builder /src/templates ./templates

WORKDIR /app
ENTRYPOINT ["/app/run-hello-expo", "-c", "/app/config/conf.yaml"]
