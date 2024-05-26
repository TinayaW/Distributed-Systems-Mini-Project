package router

import (
	"database/sql"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	_ "github.com/lib/pq"
)

type Challenge struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	TemplateFile  []byte `json:"templatefile"`
	ReadmeFile    []byte `json:"readmefile"`
	Difficulty    string `json:"difficulty"`
	TestCasesFile []byte `json:"testfasesfile"`
	AuthorID      int    `json:"authorid"`
}

type ChallengeMin struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Difficulty string `json:"difficulty"`
	AuthorID   int    `json:"authorid"`
}

func SetupRouter(consulClient *api.Client, db *sql.DB) *gin.Engine {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20

	router.GET("/challenge/challenges", getChallenges(db))
	router.GET("/challenge/challenges/user/:authorid", getUserChallenges(db))
	router.GET("/challenge/:id", getChallengeById(db))
	router.GET("/challenge/difficulty/:difficulty", getChallengesByDifficulty(db))
	router.POST("/challenge/create", createChallenge(db))
	router.PUT("/challenge/update/:id", updateChallenge(db))
	router.DELETE("/challenge/delete/:id", deleteChallenge(db))

	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	return router
}

func getChallenges(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		rows, err := db.Query("SELECT id, title, difficulty, authorid FROM challenge")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var challenges []ChallengeMin
		for rows.Next() {
			var a ChallengeMin
			err := rows.Scan(&a.ID, &a.Title, &a.Difficulty, &a.AuthorID)
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
}

func getUserChallenges(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		authorid := c.Param("authorid")

		rows, err := db.Query("SELECT id, title, difficulty, authorid FROM challenge WHERE authorid = $1", authorid)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var challenges []ChallengeMin
		for rows.Next() {
			var a ChallengeMin
			err := rows.Scan(&a.ID, &a.Title, &a.Difficulty, &a.AuthorID)
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
}

func getChallengesByDifficulty(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		difficulty := c.Param("difficulty")

		rows, err := db.Query("SELECT id, title, difficulty, authorid FROM challenge WHERE difficulty = $1", difficulty)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var challenges []ChallengeMin
		for rows.Next() {
			var a ChallengeMin
			err := rows.Scan(&a.ID, &a.Title, &a.Difficulty, &a.AuthorID)
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
}

func getChallengeById(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		id := c.Param("id")

		rows, err := db.Query("SELECT id, title, templatefile, readmefile, difficulty, testcasesfile, authorid FROM challenge WHERE id = $1", id)
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

			// err = saveFilesLocally(challenge)
			// if err != nil {
			// 	log.Fatal(err)
			// }
		} else {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Challenge not found"})
			return
		}

		c.IndentedJSON(http.StatusOK, challenge)
	}
}

// func saveFilesLocally(challenge Challenge) error {
// 	err := os.WriteFile(fmt.Sprintf("template_%d.zip", challenge.ID), challenge.TemplateFile, 0644)
// 	if err != nil {
// 		return err
// 	}

// 	err = os.WriteFile(fmt.Sprintf("readme_%d.md", challenge.ID), challenge.ReadmeFile, 0644)
// 	if err != nil {
// 		return err
// 	}

// 	err = os.WriteFile(fmt.Sprintf("testcases_%d.zip", challenge.ID), challenge.TestCasesFile, 0644)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func createChallenge(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		testcaseFile, err := c.FormFile("testcase")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		templateFile, err := c.FormFile("template")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		readmeFile, err := c.FormFile("readme")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		testcase, err := testcaseFile.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open testcase file"})
			return
		}
		defer testcase.Close()

		testcaseBytes, err := io.ReadAll(testcase)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read testcase file"})
			return
		}

		template, err := templateFile.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open template file"})
			return
		}
		defer template.Close()
		templateBytes, err := io.ReadAll(template)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read template file"})
			return
		}

		readme, err := readmeFile.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open readme file"})
			return
		}
		defer readme.Close()
		readmeBytes, err := io.ReadAll(readme)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read readme file"})
			return
		}

		id := c.PostForm("id")
		title := c.PostForm("title")
		difficulty := c.PostForm("difficulty")
		authorid := c.PostForm("authorid")

		_, err = db.Exec("INSERT INTO challenge (id, title, templatefile, readmefile, difficulty, testcasesfile, authorid) VALUES ($1, $2, $3, $4, $5, $6, $7)", id, title, templateBytes, readmeBytes, difficulty, testcaseBytes, authorid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Challenge created successfully"})
	}
}

func updateChallenge(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		testcaseFile, err := c.FormFile("testcase")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		templateFile, err := c.FormFile("template")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		readmeFile, err := c.FormFile("readme")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		testcase, err := testcaseFile.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open testcase file"})
			return
		}
		defer testcase.Close()

		testcaseBytes, err := io.ReadAll(testcase)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read testcase file"})
			return
		}

		template, err := templateFile.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open template file"})
			return
		}
		defer template.Close()
		templateBytes, err := io.ReadAll(template)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read template file"})
			return
		}

		readme, err := readmeFile.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open readme file"})
			return
		}
		defer readme.Close()
		readmeBytes, err := io.ReadAll(readme)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read readme file"})
			return
		}

		title := c.PostForm("title")
		difficulty := c.PostForm("difficulty")
		authorid := c.PostForm("authorid")

		id := c.Param("id")

		stmt, err := db.Prepare("UPDATE challenge SET title=$2, templatefile=$3, readmefile=$4, difficulty=$5, testcasesfile=$6, authorid=$7 WHERE id=$1")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		result, err := stmt.Exec(id, title, templateBytes, readmeBytes, difficulty, testcaseBytes, authorid)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update challenge"})
			return
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get rows affected"})
			return
		}

		if rowsAffected == 0 {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Challenge not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Challenge updated successfully"})
	}
}

func deleteChallenge(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		id := c.Param("id")

		stmt, err := db.Prepare("DELETE FROM challenge WHERE id=$1")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		result, err := stmt.Exec(id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete challenge"})
			return
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get rows affected"})
			return
		}

		if rowsAffected == 0 {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Challenge not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Challenge deleted successfully"})
	}
}
