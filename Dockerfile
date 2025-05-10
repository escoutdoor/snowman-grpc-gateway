FROM golang:1.23-alpine3.21 as builder

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY "go.sum" "go.mod" ./
RUN go mod download

COPY . .

# RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/app ./cmd/snowman/main.go
#
# FROM alpine:3.21
# WORKDIR /app
#
# COPY --from=builder /app/bin/app .
# ENTRYPOINT ["/app/app"]

CMD ["air"]
