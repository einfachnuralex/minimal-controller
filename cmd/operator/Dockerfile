FROM --platform=$BUILDPLATFORM golang:alpine AS build

ARG TARGETOS TARGETARCH

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN --mount=type=cache,target=/root/.cache/go-build \
  --mount=type=cache,target=/go/pkg \
  GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o operator -ldflags "-s -w" ./cmd/operator

FROM alpine

COPY --from=build /src/operator /bin/operator

ENTRYPOINT [ "operator" ]
