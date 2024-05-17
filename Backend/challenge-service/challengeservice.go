package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

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
	router.GET("/challenge/getchallenges", getChallenges)
	router.GET("/challenge/getchallengebyid/:id", getChallengeById)
	router.POST("/challenge/createchallenge", createChallenge)

	serverAddress := config.Server.Address + ":" + strconv.Itoa(config.Server.Port)
	router.Run(serverAddress)
}

func getChallenges(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	rows, err := db.Query("SELECT id, title, templatefile, readmefile, difficulty, testcasesfile, authorid FROM challenge")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var challenges []Challenge
	for rows.Next() {
		var a Challenge
		err := rows.Scan(&a.ID, &a.Title, &a.TemplateFile, &a.ReadmeFile, &a.Difficulty, &a.TestCasesFile, &a.AuthorID)
		if err != nil {
			log.Fatal(err)
		}
		challenges = append(challenges, a)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, challenges)
}

func getChallengeById(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	id := c.Param("id")

	rows, err := db.Query("SELECT  id, title, templatefile, readmefile, difficulty, testcasesfile, authorid FROM challenge WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var challenge Challenge
	if rows.Next() {
		err := rows.Scan(&challenge.ID, &challenge.Title, &challenge.TemplateFile, &challenge.ReadmeFile, &challenge.Difficulty, &challenge.TestCasesFile, &challenge.AuthorID)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Challenge not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, challenge)
}

func createChallenge(c *gin.Context) {

	var challengeAlbum Challenge
	log.Println(c.Request.Body)
	if err := c.ShouldBindJSON(&challengeAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stmt, err := db.Prepare("INSERT INTO challenge (id, title, templatefile, readmefile, difficulty, testcasesfile, authorid) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(challengeAlbum.ID, challengeAlbum.Title, challengeAlbum.Difficulty, challengeAlbum.AuthorID); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, challengeAlbum)
}
