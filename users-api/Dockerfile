FROM golang:alpine AS builder

ENV GIN_MODE=release
ENV PORT=8081

ENV DB_HOST="localhost:3306"
ENV DB_DATABASE="usersdb"
ENV DB_USERNAME="root"
ENV DB_PASSWORD="root"

ENV KEYCLOAK_URL="http://localhost:8082/"
ENV KEYCLOAK_REALM="money-house"
ENV KEYCLOAK_CLIENT_ID="users-api"
ENV KEYCLOAK_CLIENT_SECRET="aBC5irvTeiVOhF4nNBRN2ZZ0JLh35N8L"

WORKDIR /go/src

COPY . .

RUN go mod download && go mod verify
RUN go build cmd/api/main.go

EXPOSE $PORT

FROM alpine

WORKDIR /app
COPY --from=builder /go/src/main .
COPY --from=builder /go/src/.env .
COPY --from=builder /go/src/docs/* ./docs/*

ENTRYPOINT ["./main"]