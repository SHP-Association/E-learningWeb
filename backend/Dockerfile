FROM golang:1.24-alpine AS builder

WORKDIR /src

RUN apk add --no-cache build-base

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /out/web ./cmd/web

FROM alpine:3.20

WORKDIR /app

RUN addgroup -S app && adduser -S -G app app \
    && apk add --no-cache ca-certificates tzdata curl sqlite-libs

COPY --from=builder /out/web /app/web
COPY config/config.example.yaml /app/config/config.example.yaml

RUN chown -R app:app /app
USER app

EXPOSE 8000

HEALTHCHECK --interval=30s --timeout=5s --start-period=20s --retries=3 \
  CMD curl -fsS http://127.0.0.1:${HTTP_PORT:-8000}/api/health || exit 1

CMD ["/app/web"]
