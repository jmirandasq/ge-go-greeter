# Stage 1: build
FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o server ./cmd/api/main.go

# Stage 2: run
FROM gcr.io/distroless/static:nonroot

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080

USER nonroot:nonroot

CMD ["./server"]
