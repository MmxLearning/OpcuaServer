FROM golang:alpine AS builder

WORKDIR /build

COPY . .

RUN go env -w GO111MODULE=auto \
    && go env -w CGO_ENABLED=0 \
    && set -ex \
    && go build -ldflags "-s -w" -o runner ./cmd/server

FROM golang:alpine

RUN apk update && \
    apk upgrade --no-cache && \
    rm -rf /var/cache/apk/*

COPY --from=builder  /build/runner /usr/bin/runner
RUN chmod +x /usr/bin/runner

ENTRYPOINT [ "/usr/bin/runner" ]