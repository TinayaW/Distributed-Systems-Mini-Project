FROM golang:1.22.1 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -o myapp

FROM alpine:3
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/myapp /myapp

EXPOSE 8080

CMD ["/myapp"]
