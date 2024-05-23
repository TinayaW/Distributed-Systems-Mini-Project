FROM golang:1.22.1 As builder

RUN apt-get update && apt-get install -y unzip

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o submission-service .

EXPOSE 8082

CMD ["./submission-service"]
