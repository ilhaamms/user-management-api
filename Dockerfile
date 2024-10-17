FROM golang:1.22-alpine as builder

WORKDIR /app

ENV MYSQL_USER=root
ENV MYSQL_PASSWORD=ilham
ENV MYSQL_DATABASE=user_management_api
ENV MYSQL_HOST=mysql
ENV MYSQL_PORT=3306

COPY go.mod go.sum ./
RUN go mod download

RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest


COPY . .


RUN go build -o /app/main .

FROM alpine:3.12

WORKDIR /app

COPY --from=builder /app/config /app/config
COPY --from=builder /app/main /app/main
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate

EXPOSE 8080

CMD migrate -database "mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DATABASE}" -path config/db/migrations up && /app/main