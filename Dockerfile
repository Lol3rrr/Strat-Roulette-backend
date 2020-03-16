# Building the binary of the App

FROM golang:1.13 AS build

WORKDIR /go/src/strat-roulette-backend

COPY . .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .


# Moving the binary to the 'final Image' to make it smaller
FROM alpine:latest

WORKDIR /app

COPY --from=build /go/src/strat-roulette-backend/app .

EXPOSE 80

CMD ["./app"]