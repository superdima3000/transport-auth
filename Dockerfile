FROM golang:1.26.1-bookworm AS builder

WORKDIR /app


RUN apt-get update && apt-get install -y --no-install-recommends \
    gcc \
    libc6-dev \
    git \
    ca-certificates \
    libsqlite3-dev \
 && rm -rf /var/lib/apt/lists/*

COPY go.mod go.sum ./
ENV GOPROXY=direct
ENV GOSUMDB=off
RUN go mod download

RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . .

RUN /go/bin/swag init --parseDependency --parseInternal --dir ./ --generalInfo cmd/main.go

RUN CGO_ENABLED=1 GOOS=linux go build -ldflags="-s -w" -o server ./cmd/main.go

FROM golang:1.26.1-bookworm

WORKDIR /app

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates libsqlite3-0 \
 && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/server .
COPY --from=builder /app/docs ./docs
COPY configuration/ /app/configuration/

VOLUME ["/app/data"]

EXPOSE 8080

CMD ["./server"]