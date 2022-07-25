FROM golang:1.18.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o entrypoint ./cmd/server/main.go

FROM alpine

COPY --from=builder /app/entrypoint /entrypoint

CMD [ "/entrypoint" ]