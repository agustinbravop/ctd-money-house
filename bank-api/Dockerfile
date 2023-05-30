FROM golang:alpine AS builder

ENV GIN_MODE=release
ENV PORT=8084

WORKDIR /go/src

COPY . .

RUN go mod download && go mod verify
RUN go build cmd/api/main.go

EXPOSE $PORT

FROM alpine

WORKDIR /app
COPY --from=builder /go/src/main .
COPY --from=builder /go/src/.env .

ENTRYPOINT ["./main"]