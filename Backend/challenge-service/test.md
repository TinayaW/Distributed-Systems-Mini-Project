# Run

1. `go mod init backend/challenge-service`
2. `go get .`
3. `go run .`

# Test

## Create User

curl -X POST \
  http://localhost:8080/challenge/createchallenge \
  -H 'Content-Type: application/json' \
  -d '{
    "id": 1,
    "title": "Example Challenge",
    "templatefile": "example_template_content",
    "readmefile": "example_readme_content",
    "difficulty": "medium",
    "testcasesfile": "example_test_cases_content",
    "authorid": 123
  }'

## Get Challenges

curl -X GET http://localhost:8080/challenge/getchallenges

## Get Challenge By ID

curl -X GET http://localhost:8080/challenge/getchallengebyid/1