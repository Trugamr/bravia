# Build on the runner's native architecture and cross-compile for the target
# platform, avoiding emulation of the Go toolchain in multi-platform builds.
FROM --platform=$BUILDPLATFORM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG TARGETOS
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH \
    go build -trimpath -ldflags="-s -w" -o /app/remote ./cmd/remote

FROM gcr.io/distroless/static-debian12 AS release

COPY --from=builder /app/remote /usr/local/bin/remote

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/usr/local/bin/remote"]
