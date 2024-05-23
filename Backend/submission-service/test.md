# Run

1. `go mod init backend/submission-service`
2. `go get .`
3. `go run .`

# Test

## Upload Submission

curl -X POST \
  http://localhost:8082/submission/upload \
  -H "Content-Type: multipart/form-data" \
  -F "file=@/home/manuja/DisProject/ProblemSolvingPlatform/Backend/submission-service/testfiles/submission.zip" \
  -F "id=1" \
  -F "fileName=submission" \
  -F "fileExtension=zip" \
  -F "userId=1006" \
  -F "challengeId=1" 

## Get Submissions By UserID

curl -X GET http://localhost:8082/submission/user/1

## Get Submissions By ChallengeID

curl -X GET http://localhost:8082/submission/challenge/1

## Get Submission By ID

curl -X GET http://localhost:8082/submission/1