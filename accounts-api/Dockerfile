FROM golang:alpine AS builder

ENV GIN_MODE=release
ENV PORT=8083

ENV DB_HOST="localhost:3306"
ENV DB_DATABASE="ACCOUNTSDB"
ENV DB_USERNAME="root"
ENV DB_PASSWORD="root"

ENV USERS_API_HOST="http://localhost:8081"
ENV BANK_API_HOST="http://localhost:8084"

WORKDIR /go/src

COPY . .

RUN go mod download && go mod verify
RUN go build cmd/api/main.go

EXPOSE $PORT

FROM alpine

WORKDIR /app
COPY --from=builder /go/src/main .
COPY --from=builder /go/src/.env .
COPY --from=builder /go/src/pkg/utils/aliases.txt ./pkg/utils/aliases.txt
COPY --from=builder /go/src/docs/* ./docs/*

ENTRYPOINT ["./main"]