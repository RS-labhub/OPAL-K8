# Build the manager binary -> doc
FROM golang:1.16 as builder

WORKDIR /workspace
# Copy the Go Modules manifests -> doc
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much -> doc
# and so that source changes don't invalidate our downloaded layer -> doc
RUN go mod download

# Copy the go source -> doc
COPY main.go main.go
COPY api/ api/
COPY controllers/ controllers/

# Build -> doc
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o manager main.go

# Use distroless as minimal base image to package the manager binary -> doc
# Refer to https://github.com/GoogleContainerTools/distroless for more details -> doc
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/manager .
USER 65532:65532

ENTRYPOINT ["/manager"]

# FROM golang:1.16 as builder
# WORKDIR /workspace
# COPY . .
# RUN go mod tidy
# RUN go build -o opal-controller main.go

# FROM gcr.io/distroless/base-debian10
# WORKDIR /
# COPY --from=builder /workspace/opal-controller .
# USER nonroot:nonroot
# ENTRYPOINT ["/opal-controller"]
