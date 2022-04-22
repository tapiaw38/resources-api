FROM golang:1.13.4

LABEL maintainer="Walter <tapiaw38@gmail.com>"

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
COPY .env .

RUN go build

CMD ["./resources-api"]