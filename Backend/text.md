# Run

1. `go get .`
2. `go run .`

# Test

## Create Album

curl -X POST \
  http://localhost:8080/albums \
  -H 'Content-Type: application/json' \
  -d '{
    "id": "1",
    "title": "Example Album",
    "artist": "Example Artist",
    "price": 9.99
  }'

## Get Album

curl -X GET http://localhost:8080/albums
