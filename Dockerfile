FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o wobble ./server/cmd

EXPOSE 3000

CMD ["./wobble"]