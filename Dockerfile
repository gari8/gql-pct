FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN apk update &&  apk add git
RUN go install github.com/cosmtrek/air@latest

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]