package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
)

var db *sql.DB
var secretKey = "apikey"

func generateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func loginAuth(c *gin.Context) {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var passwordHash string
	err := db.QueryRow("SELECT password FROM users WHERE username=?", request.Username).Scan(&passwordHash)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	if passwordHash != hashPassword(request.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token, err := generateToken(request.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// func main() {
// 	var err error
// 	db, err = sql.Open("mysql", "kai:Ricachonda18$$@tcp(localhost:3306)/srvdb")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	r := gin.Default()

// 	r.GET("/api/users", getUsers)
// 	r.POST("/api/users/create", createUser)
// 	r.PUT("/api/users/update", updateUser)
// 	r.DELETE("/api/users/delete", deleteUser)

// 	// Enable CORS
// 	r.Use(cors.New(cors.Config{
// 		AllowOrigins:     []string{"http://localhost:9000"},
// 		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
// 		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
// 		AllowCredentials: true,
// 	}))

// 	r.POST("/api/login", login)
// 	r.Run(":8080")
// }
