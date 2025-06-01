package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/pneubeck/NFCommunicator_Backend/models"
)

var Db *sql.DB

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error is occurred  on .env file please check")
	}
	host := os.Getenv("HOST")
	port, _ := strconv.Atoi(os.Getenv("PORT")) // don't forget to convert int since port is int type.
	user := os.Getenv("USERNAME")
	dbname := os.Getenv("DB_NAME")
	pass := os.Getenv("PASSWORD")
	//TODO: Enable ssl in production
	psqlSetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, pass)
	db, errSql := sql.Open("postgres", psqlSetup)
	if errSql != nil {
		fmt.Println("There is an error while connecting to the database ", errSql)
		panic(errSql)
	} else {
		Db = db
		fmt.Println("Successfully connected database!")
	}
	router := gin.Default()
	router.GET("/NextUserId", getNextUserId)
	router.POST("/PostMessage", postMessage)
	router.Run(":8080")
}

func getNextUserId(context *gin.Context) {
	var nextUserId int
	tx, err := Db.BeginTx(context, nil)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, err)
	}
	defer tx.Rollback()
	_, err = tx.Exec("Update LastUserId SET LastUserId = LastUserId + 1")
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, err)
	}
	if err := tx.QueryRow("SELECT LastUserId FROM LastUserId").Scan(&nextUserId); err != nil {
		context.IndentedJSON(http.StatusInternalServerError, err)
	}
	tx.Commit()
	context.JSON(http.StatusOK, nextUserId)
}

func postMessage(c *gin.Context) {
	var newTodo models.Message
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// addedTodo := services.AddTodoItem(&newTodo)

	// c.JSON(http.StatusCreated, addedTodo)

}
