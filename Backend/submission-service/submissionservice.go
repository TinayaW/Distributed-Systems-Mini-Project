package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Config struct {
	Database struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		DBName   string `json:"dbname"`
		SSLMode  string `json:"sslmode"`
	} `json:"database"`
	Server struct {
		Address string `json:"address"`
		Port    int    `json:"port"`
	} `json:"server"`
	Services struct {
		ApiGateway string `json:"apigateway"`
	} `json:"services"`
}

type Submission struct {
	ID            int     `json:"id"`
	Score         float64 `json:"score"`
	ChallengeID   int     `json:"challengeId"`
	UserID        int     `json:"userId"`
	FileName      string  `json:"fileName"`
	FileExtension string  `json:"fileExtension"`
	File          []byte  `json:"file"`
}

type Challenge struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	TemplateFile  []byte `json:"templatefile"`
	ReadmeFile    []byte `json:"readmefile"`
	Difficulty    string `json:"difficulty"`
	TestCasesFile []byte `json:"testfasesfile"`
	AuthorID      int    `json:"authorid"`
}

var db *sql.DB

func main() {

	configFile, openErr := os.Open("config.json")
	if openErr != nil {
		log.Fatal("Error opening config file:", openErr)
	}
	defer configFile.Close()

	var config Config
	decodeErr := json.NewDecoder(configFile).Decode(&config)
	if decodeErr != nil {
		log.Fatal("Error decoding config JSON:", decodeErr)
	}

	dbInfo := "host=" + config.Database.Host +
		" port=" + strconv.Itoa(config.Database.Port) +
		" user=" + config.Database.User +
		" password=" + config.Database.Password +
		" dbname=" + config.Database.DBName +
		" sslmode=" + config.Database.SSLMode

	var err error
	db, err = sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.GET("/submission/user/:userid", getUserSubmissions)
	router.GET("/submission/challenge/:challengeid", getChallengeSubmissions)
	router.GET("/submission/:id", getSubmissionById)
	router.POST("/submission/upload", func(c *gin.Context) {
		uploadSubmission(c, &config)
	})

	serverAddress := config.Server.Address + ":" + strconv.Itoa(config.Server.Port)
	router.Run(serverAddress)
}

func getUserSubmissions(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	userId := c.Param("userid")

	rows, err := db.Query("SELECT id, score, userId, challengeId, fileName, fileExtension, file FROM submission WHERE userId = $1", userId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var submissions []Submission
	for rows.Next() {
		var a Submission
		err := rows.Scan(&a.ID, &a.Score, &a.UserID, &a.ChallengeID, &a.FileName, &a.FileExtension, &a.File)
		if err != nil {
			log.Fatal(err)
		}

		// err = saveFilesLocally(a)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		submissions = append(submissions, a)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, submissions)
}

func getChallengeSubmissions(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	challengeId := c.Param("challengeid")

	rows, err := db.Query("SELECT id, score, userId, challengeId, fileName, fileExtension, file FROM submission WHERE challengeId = $1", challengeId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var submissions []Submission
	for rows.Next() {
		var a Submission
		err := rows.Scan(&a.ID, &a.Score, &a.UserID, &a.ChallengeID, &a.FileName, &a.FileExtension, &a.File)
		if err != nil {
			log.Fatal(err)
		}

		// err = saveFilesLocally(a)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		submissions = append(submissions, a)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, submissions)
}

func getSubmissionById(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	id := c.Param("id")

	rows, err := db.Query("SELECT id, score, userId, challengeId, fileName, fileExtension, file FROM submission WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var submission Submission
	if rows.Next() {
		err := rows.Scan(&submission.ID, &submission.Score, &submission.UserID, &submission.ChallengeID, &submission.FileName, &submission.FileExtension, &submission.File)
		if err != nil {
			log.Fatal(err)
		}

		// err = saveFilesLocally(submission)
		// if err != nil {
		// 	log.Fatal(err)
		// }
	} else {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "submission not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, submission)
}

// func saveFilesLocally(submission Submission) error {
// 	err := os.WriteFile(fmt.Sprintf("submission%d.zip", submission.ID), submission.File, 0644)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func uploadSubmission(c *gin.Context, config *Config) {

	if _, err := os.Stat("temp"); os.IsNotExist(err) {
		os.Mkdir("temp", 0755)
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	openFile, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer openFile.Close()

	fileBytes, err := io.ReadAll(openFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}

	id := c.PostForm("id")
	userId := c.PostForm("userId")
	challengeId := c.PostForm("challengeId")
	fileName := c.PostForm("fileName")
	fileExtension := c.PostForm("fileExtension")

	challengeResp, err := http.Get(config.Services.ApiGateway + "/challenge/" + challengeId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch challenge: " + err.Error()})
		return
	}
	defer challengeResp.Body.Close()

	if challengeResp.StatusCode != http.StatusOK {
		c.JSON(http.StatusNotFound, gin.H{"error": "Challenge not found"})
		return
	}

	var challenge Challenge
	err = json.NewDecoder(challengeResp.Body).Decode(&challenge)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode challenge: " + err.Error()})
		return
	}

	err = saveChallengeFilesLocally(challenge)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save challenge files: " + err.Error()})
		return
	}

	_, err = exec.Command("/bin/bash", "-c", "unzip -o temp/testcases.zip -d temp").Output()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unzip testcases: " + err.Error()})
		return
	}

	err = saveSubmissionFilesLocally(Submission{File: fileBytes})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save submission files: " + err.Error()})
		return
	}

	_, err = exec.Command("/bin/bash", "-c", "unzip -o temp/submission.zip -d temp").Output()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unzip submission: " + err.Error()})
		return
	}

	score, err := exec.Command("python3", "temp/testcase.py").Output()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to run testcases: " + err.Error()})
		return
	}

	trimmedScore := strings.TrimSpace(string(score))
	scoreFloat, err := strconv.ParseFloat(trimmedScore, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse score: " + err.Error()})
		return
	}
	_, err = db.Exec("INSERT INTO submission (id, score, userId, challengeId, fileName, fileExtension, file) VALUES ($1, $2, $3, $4, $5, $6, $7)", id, scoreFloat, userId, challengeId, fileName, fileExtension, fileBytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = os.RemoveAll("temp")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete temp folder: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Submission created successfully"})
}

func saveChallengeFilesLocally(challenge Challenge) error {
	err := os.WriteFile("temp/testcases.zip", challenge.TestCasesFile, 0644)
	if err != nil {
		return err
	}
	return nil
}

func saveSubmissionFilesLocally(submission Submission) error {
	err := os.WriteFile("temp/submission.zip", submission.File, 0644)
	if err != nil {
		return err
	}
	return nil
}
