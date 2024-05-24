*If not use docker compose use below methods.*

# Run

1. `go mod init backend/user-service`
2. `go get .`
3. `go run .`

# Docker run

1. `docker build --tag user-service .`
2. `docker run --publish 8080:8080 user-service`

# Test

### 1. Create user

curl -X POST \
  http://localhost:8080/user/create \
  -H 'Content-Type: application/json' \
  -d '{
    "id": 2,
    "username": "MrManchi",
    "fullname": "Manuja"
  }'

### 2. Get users

curl -X GET http://localhost:8080/user/users

### 3. Get user by ID

curl -X GET http://localhost:8080/user/1

### 4. Update user

curl -X PUT \
  http://localhost:8080/user/update/3 \
  -H 'Content-Type: application/json' \
  -d '{
    "username": "MrD",
    "fullname": "D Kariyawasam"
}'

### 5. Delete user

curl -X DELETE \
  http://localhost:8080/user/delete/1