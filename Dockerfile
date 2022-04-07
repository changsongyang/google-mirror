FROM golang:alpine AS builder

LABEL org.opencontainers.image.source https://github.com/yangchuansheng/google-mirror
LABEL maintainer="disqus-php-api Docker Maintainers https://icloudnative.io"

RUN apk update
RUN apk add --no-cache git

WORKDIR /app

ADD . /app

RUN go get github.com/dsnet/compress/brotli
RUN CGO_ENABLED=0 go build -o /app/google-mirror

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/google-mirror /app/google-mirror

EXPOSE 3000

CMD ["/app/google-mirror"]
