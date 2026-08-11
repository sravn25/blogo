# ---- Build stage ----
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o blogo .

# ---- Final stage ----
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/blogo .

EXPOSE 8080 

CMD ["./blogo"]
