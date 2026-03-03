package main

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
)

func hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", hash)
}


func dbConnect() (*sql.DB, error) {
	dbUser := "kai"
	dbPassword := "admin"
	dbName := "srvdb"
	dbHost := "localhost"
	dbPort := "3306"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	return sql.Open("mysql", dsn)
}

func main() {
	var err error
	db, err = sql.Open("mysql", "kai:Ricachonda18$$@tcp(localhost:3306)/srvdb")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//--------------------------LOGIN-AUTH-------------------------------
	r := gin.Default()
	r.POST("/api/login", loginAuth)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	//--------------------------LOGIN-AUTH-------------------------------

	r.GET("/api/employees", func(c *gin.Context) {
		getEmployees(c.Writer, c.Request)
	})
	r.POST("/api/employees/create", func(c *gin.Context) {
		createEmployee(c.Writer, c.Request)
	})
	r.PUT("/api/employees/update", func(c *gin.Context) {
		updateEmployee(c.Writer, c.Request)
	})
	r.DELETE("/api/employees/delete", func(c *gin.Context) {
		deleteEmployee(c.Writer, c.Request)
	})

	r.GET("/api/users", func(c *gin.Context) {
		getUsers(c.Writer, c.Request)
	})
	r.POST("/api/users/create", func(c *gin.Context) {
		createUser(c.Writer, c.Request)
	})
	r.PUT("/api/users/update", func(c *gin.Context) {
		updateUser(c.Writer, c.Request)
	})
	r.DELETE("/api/users/delete", func(c *gin.Context) {
		deleteUser(c.Writer, c.Request)
	})

	// Enable CORS
	corsOptions := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)

	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", corsOptions(r)))
}
