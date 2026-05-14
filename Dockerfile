FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main cmd/api/main.go

FROM golang:1.24-alpine AS development
WORKDIR /app

RUN apk add --no-cache git bash
COPY --from=builder /app /app
EXPOSE 3001
CMD ["go", "run", "cmd/api/main.go"]

FROM alpine:latest AS production
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs

EXPOSE 3001

CMD ["./main"]
