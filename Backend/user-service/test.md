# Run

1. `go mod init backend/user-service`
2. `go get .`
3. `go run .`

# Test

## Create User

curl -X POST \
  http://localhost:8080/user/create \
  -H 'Content-Type: application/json' \
  -d '{
    "id": 1,
    "username": "MrManchi",
    "fullname": "Manuja Dewmina"
  }'

## Get Users

curl -X GET http://localhost:8080/user/users

## Get User By ID

curl -X GET http://localhost:8080/user/1