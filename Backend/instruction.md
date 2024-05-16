# Docker run

1. `docker-compose up --build`

# Run

`go mod init backend/user-service`

1. `go get .`
2. `go run .`

# Test

## Create Album

curl -X POST \
  http://localhost:8080/createuser \
  -H 'Content-Type: application/json' \
  -d '{
    "id": "1",
    "username": "MrManchi",
    "fullname": "Manuja Dewmina"
  }'

## Get Album

curl -X GET http://localhost:8080/getusers
