FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shorturl .

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

ENV TZ=Asia/Shanghai

RUN adduser -D -g '' appuser

COPY --from=builder /app/shorturl /app/shorturl

USER appuser

WORKDIR /app

EXPOSE 8080

CMD ["./shorturl"]