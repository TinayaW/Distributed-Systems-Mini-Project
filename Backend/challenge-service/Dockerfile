FROM golang:1.22.1 As builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o challenge-service ./cmd/challenge-service/

EXPOSE 8081

CMD ["./challenge-service"]
 