FROM golang:1.16.4-alpine as dev

WORKDIR /app

COPY . .

CMD [ "go", "run", "main.go"]

FROM golang:1.16.4-alpine as build

WORKDIR /build

COPY . .

RUN go build -o test-go-app ./main.go

FROM alpine as prd

WORKDIR /app

COPY --from=build /build/test-go-app .

RUN apk add --no-cache mysql-client

RUN addgroup go \
    && adduser -D -G go go \
    && chown -R go:go /app/test-go-app

CMD ["./test-go-app"]