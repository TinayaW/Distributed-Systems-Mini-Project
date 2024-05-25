package router

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

	"backend/submission-service/config"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	_ "github.com/lib/pq"
)

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

func SetupRouter(consulClient *api.Client, db *sql.DB, configdata *config.Config) *gin.Engine {
	router := gin.Default()

	router.GET("/submission/user/:userid", getUserSubmissions(db))
	router.GET("/submission/challenge/:challengeid", getChallengeSubmissions(db))
	router.GET("/submission/:id", getSubmissionById(db))
	router.POST("/submission/upload", uploadSubmission(db, configdata))

	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	return router
}

func getUserSubmissions(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
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
}

func getChallengeSubmissions(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
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
}

func getSubmissionById(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
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
}

// func saveFilesLocally(submission Submission) error {
// 	err := os.WriteFile(fmt.Sprintf("submission%d.zip", submission.ID), submission.File, 0644)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func uploadSubmission(db *sql.DB, config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {

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
