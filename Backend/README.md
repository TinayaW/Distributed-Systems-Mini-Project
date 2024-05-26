# Run

1. `docker-compose up --build` 

# Test

## 1. User service

### 1.1. Create user

curl -X POST \
  http://localhost:8083/user/create \
  -H 'Content-Type: application/json' \
  -d '{
    "id": 1,
    "username": "MrManchi",
    "fullname": "Manuja",
    "userpassword": "test"
  }'

### 1.2. Get users

curl -X GET http://localhost:8083/user/users

### 1.3. Get user by ID

curl -X GET http://localhost:8083/user/1

### 1.4. Get user by username

curl -X GET http://localhost:8083/user/username/MrManchi

### 1.5. Update user

curl -X PUT \
  http://localhost:8083/user/update/1 \
  -H 'Content-Type: application/json' \
  -d '{
    "username": "MrDew",
    "fullname": "Dewmina Kariyawasam"
}'

### 1.6. Delete user

curl -X DELETE \
  http://localhost:8083/user/delete/1

## 2. Challenge service

### 2.1. Create challenge

curl -X POST \
  http://localhost:8083/challenge/create \
  -H "Content-Type: multipart/form-data" \
  -F "testcase=@/home/manuja/DisProject/ProblemSolvingPlatform/Backend/challenge-service/testfiles/testcase_80.zip" \
  -F "template=@/home/manuja/DisProject/ProblemSolvingPlatform/Backend/challenge-service/testfiles/template.zip" \
  -F "readme=@/home/manuja/DisProject/ProblemSolvingPlatform/Backend/challenge-service/testfiles/readme.md" \
  -F "id=1" \
  -F "title=Example Challenge 1" \
  -F "difficulty=easy" \
  -F "authorid=1001" 

### 2.2. Get challenges

curl -X GET http://localhost:8083/challenge/challenges

### 2.3. Get user challenges

curl -X GET http://localhost:8083/challenge/challenges/user/782077

### 2.3. Get challenge by ID

curl -X GET http://localhost:8083/challenge/1

### 2.4. Get challenge by difficulty

curl -X GET http://localhost:8083/challenge/difficulty/easy

### 2.5. Update challenge

curl -X PUT http://localhost:8083/challenge/update/1 \
  -H "Content-Type: multipart/form-data" \
  -F "testcase=@/home/manuja/DisProject/ProblemSolvingPlatform/Backend/challenge-service/testfiles/testcase_100.zip" \
  -F "template=@/home/manuja/DisProject/ProblemSolvingPlatform/Backend/challenge-service/testfiles/template.zip" \
  -F "readme=@/home/manuja/DisProject/ProblemSolvingPlatform/Backend/challenge-service/testfiles/readme.md" \
  -F "title=Changed Challenge Title" \
  -F "difficulty=Medium" \
  -F "authorid=1235"

### 2.6. Delete challenge

curl -X DELETE \
  http://localhost:8083/challenge/delete/1

## 3. Submission service

### 3.1. Upload submission

curl -X POST \
  http://localhost:8083/submission/upload \
  -H "Content-Type: multipart/form-data" \
  -F "file=@/home/manuja/DisProject/ProblemSolvingPlatform/Backend/submission-service/testfiles/submission.zip" \
  -F "id=1" \
  -F "fileName=submission" \
  -F "fileExtension=zip" \
  -F "userId=1006" \
  -F "challengeId=1" 

### 3.2. Get submissions by userID

curl -X GET http://localhost:8083/submission/user/1006

### 3.3. Get submissions by challengeID

curl -X GET http://localhost:8083/submission/challenge/1

### 3.4. Get submission by ID

curl -X GET http://localhost:8083/submission/1