FROM golang:1.24.2-alpine AS devlop

WORKDIR /app

RUN go install github.com/air-verse/air@latest
CMD ["air", "-c", ".air.toml"]
