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

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
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
	router.GET("/user/getusers", getUsers)
	router.GET("/user/getuserbyid/:id", getUserById)
	router.POST("/user/createuser", createUser)

	serverAddress := config.Server.Address + ":" + strconv.Itoa(config.Server.Port)
	router.Run(serverAddress)
}

func getUsers(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	rows, err := db.Query("SELECT id, username, fullname FROM userdata")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var a User
		err := rows.Scan(&a.ID, &a.Username, &a.Fullname)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, a)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, users)
}

func getUserById(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	id := c.Param("id")

	rows, err := db.Query("SELECT id, username, fullname FROM userdata WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var user User
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Fullname)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func createUser(c *gin.Context) {

	var userAlbum User
	if err := c.BindJSON(&userAlbum); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	stmt, err := db.Prepare("INSERT INTO userdata (id, username, fullname) VALUES ($1, $2, $3)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(userAlbum.ID, userAlbum.Username, userAlbum.Fullname); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, userAlbum)
}
