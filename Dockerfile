FROM golang:1.13-alpine

LABEL maintainer="Bharath Kuppala"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 9098

CMD ["./main"]
