FROM golang:1.22.5

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV GIN_MODE=release

RUN go build -o ./map-pinner cmd/main.go

EXPOSE ${SERVER_PORT}

CMD ["./map-pinner"]
