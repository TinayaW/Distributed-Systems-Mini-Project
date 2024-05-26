package router

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	_ "github.com/lib/pq"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Password string `json:"userpassword"`
}

func SetupRouter(consulClient *api.Client, db *sql.DB) *gin.Engine {
	router := gin.Default()

	router.GET("/user/users", getUsers(db))
	router.GET("/user/:id", getUserById(db))
	router.POST("/user/create", createUser(db))
	router.PUT("/user/update/:id", updateUser(db))
	router.DELETE("/user/delete/:id", deleteUser(db))
	router.GET("/user/username/:username", getUserByUsername(db))

	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	return router
}

func getUsers(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		rows, err := db.Query("SELECT id, username, fullname, userpassword FROM userdata")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var users []User
		for rows.Next() {
			var a User
			err := rows.Scan(&a.ID, &a.Username, &a.Fullname, &a.Password)
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
}

func getUserById(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		id := c.Param("id")

		rows, err := db.Query("SELECT id, username, fullname, userpassword FROM userdata WHERE id = $1", id)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var user User
		if rows.Next() {
			err := rows.Scan(&user.ID, &user.Username, &user.Fullname, &user.Password)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.IndentedJSON(http.StatusOK, user)
	}
}

func getUserByUsername(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		id := c.Param("username")

		rows, err := db.Query("SELECT id, username, fullname, userpassword FROM userdata WHERE username = $1", id)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var user User
		if rows.Next() {
			err := rows.Scan(&user.ID, &user.Username, &user.Fullname, &user.Password)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.IndentedJSON(http.StatusOK, user)
	}
}

func createUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userAlbum User
		if err := c.BindJSON(&userAlbum); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		stmt, err := db.Prepare("INSERT INTO userdata (id, username, fullname, userpassword) VALUES ($1, $2, $3, $4)")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		if _, err := stmt.Exec(userAlbum.ID, userAlbum.Username, userAlbum.Fullname, userAlbum.Password); err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusCreated, userAlbum)
	}
}

func updateUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		id := c.Param("id")
		var user User
		if err := c.BindJSON(&user); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		stmt, err := db.Prepare("UPDATE userdata SET username=$2, fullname=$3 WHERE id=$1")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		result, err := stmt.Exec(id, user.Username, user.Fullname)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
			return
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get rows affected"})
			return
		}

		if rowsAffected == 0 {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
	}
}

func deleteUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		id := c.Param("id")

		stmt, err := db.Prepare("DELETE FROM userdata WHERE id=$1")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		result, err := stmt.Exec(id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
			return
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get rows affected"})
			return
		}

		if rowsAffected == 0 {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	}
}
