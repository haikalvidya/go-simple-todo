# Step 1: Modules caching
FROM golang:1.19.2-alpine3.16 as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Step 2: Builder
FROM golang:1.19.2-alpine3.16 as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /bin/app .

# Step 3: Final
FROM alpine:latest

# GOPATH for scratch images is /
COPY --from=builder /app/config/config.yaml /
COPY --from=builder /app/migrations /migrations
COPY --from=builder /bin/app /app
# CMD ["/app"]