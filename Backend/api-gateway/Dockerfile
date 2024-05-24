FROM golang:1.22.1 As builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o api-gateway ./cmd/api-gateway/

EXPOSE 8083

CMD ["./api-gateway"]