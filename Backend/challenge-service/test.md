# Run

1. `go mod init backend/challenge-service`
2. `go get .`
3. `go run .`

# Test

## Create User

curl -X POST \
  http://localhost:8080/challenge/createchallenge \
  -H "Content-Type: multipart/form-data" \
  -F "testcase=@/home/manuja/DisProject/ProblemSolvingPlatform/Backend/challenge-service/testfiles/testcase.zip" \
  -F "template=@/home/manuja/DisProject/ProblemSolvingPlatform/Backend/challenge-service/testfiles/template.zip" \
  -F "readme=@/home/manuja/DisProject/ProblemSolvingPlatform/Backend/challenge-service/testfiles/readme.md" \
  -F "id=2" \
  -F "title=Example Challenge 2" \
  -F "difficulty=easy" \
  -F "authorid=1234" 

## Get Challenges

curl -X GET http://localhost:8080/challenge/getchallenges

## Get Challenge By ID

curl -X GET http://localhost:8080/challenge/getchallengebyid/1