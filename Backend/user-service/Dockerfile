FROM golang:1.22.1 As builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o user-service .

EXPOSE 8080

CMD ["./user-service"]
