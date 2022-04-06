FROM cicd.harbor.vmwarecna.net/base-images/golang-build-1.16 as builder
ARG GIT_COMMIT
ARG GOPROXY

WORKDIR /src
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy in the go src
COPY ./ ./
# Build
RUN GO111MODULE=on CGO_ENABLED=0 go build -ldflags "-X main.GitCommit=$GIT_COMMIT" -a -o run-hello-expo

# Copy the controller-manager into a thin image
#FROM cicd.harbor.vmwarecna.net/base-images/golang-runtime
FROM alpine:3.14.6
WORKDIR /app
COPY --from=builder /src/run-hello-expo .
COPY --from=builder /src/config/conf.yaml ./config/
COPY --from=builder /src/templates ./templates
ENTRYPOINT ["/app/run-hello-expo"]