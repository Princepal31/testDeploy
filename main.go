package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Data struct {
	Name string `json:"fullName"`
	Age  int    `json:"age"`
}

var data = []Data{
	{"prince", 23},
	{"Noone", 24},
	{"thirdPerson", 25},
}

func main() {
	router := gin.Default()
	protectedRoute := router.Group("/protected")
	protectedRoute.Use(loggerMiddleWare, authenticationMiddleware)
	protectedRoute.GET("/showData/:id", showData)
	protectedRoute.POST("/saveData", saveData)
	router.Run("localhost:9010")
}

func showData(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("id : ", id)
	c.IndentedJSON(http.StatusOK, data)
}

func saveData(c *gin.Context) {
	data := Data{}
	if err := c.BindJSON(&data); err != nil {
		return
	}
	c.IndentedJSON(http.StatusCreated, data)
}

func loggerMiddleWare(c *gin.Context) {
	fmt.Println("Logging the data")
	c.Next()
}

func authenticationMiddleware(c *gin.Context) {
	fmt.Println("token is there")
	c.Next()
}
