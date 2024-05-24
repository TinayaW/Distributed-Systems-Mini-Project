*If not use docker compose use below methods.*

# Run

1. `go mod init backend/submission-service`
2. `go get .`
3. `go run .`

# Docker run

1. `docker build --tag submission-service .`
2. `docker run --publish 8082:8082 submission-service`

# Test

### 1. Upload submission

curl -X POST \
  http://localhost:8082/submission/upload \
  -H "Content-Type: multipart/form-data" \
  -F "file=@/home/manuja/DisProject/ProblemSolvingPlatform/Backend/submission-service/testfiles/submission.zip" \
  -F "id=1" \
  -F "fileName=submission" \
  -F "fileExtension=zip" \
  -F "userId=1006" \
  -F "challengeId=1" 

### 2. Get submissions by userID

curl -X GET http://localhost:8082/submission/user/1

### 3. Get submissions by challengeID

curl -X GET http://localhost:8082/submission/challenge/1

### 4. Get submission by ID

curl -X GET http://localhost:8082/submission/1