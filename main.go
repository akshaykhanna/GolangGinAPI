package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

var router *gin.Engine

type Users struct {
	Id        int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Firstname string `gorm:"not null" form:"firstname" json:"firstname"`
	Lastname  string `gorm:"not null" form:"lastname" json:"lastname"`
}

func main() {
	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		v1.POST("/users", PostUser)
		v1.GET("/users", GetUsers)
		v1.GET("/users/:id", GetUser)
		v1.PUT("/users/:id", UpdateUser)
		v1.DELETE("/users/:id", DeleteUser)
	}

	r.Run(":8080")
}

func PostUser(c *gin.Context) {
	// The futur code…
}

func GetUsers(c *gin.Context) {
	var users = []Users{
		Users{Id: 1, Firstname: "Oliver", Lastname: "Queen"},
		Users{Id: 2, Firstname: "Malcom", Lastname: "Merlyn"},
	}

	c.JSON(200, users)

	// curl -i http://localhost:8080/api/v1/users
}

func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	user_id, _ := strconv.ParseInt(id, 0, 64)

	if user_id == 1 {
		content := gin.H{"id": user_id, "firstname": "Oliver", "lastname": "Queen"}
		c.JSON(200, content)
	} else if user_id == 2 {
		content := gin.H{"id": user_id, "firstname": "Malcom", "lastname": "Merlyn"}
		c.JSON(200, content)
	} else {
		content := gin.H{"error": "user with id#" + id + " not found"}
		c.JSON(404, content)
	}

	// curl -i http://localhost:8080/api/v1/users/1
}

func UpdateUser(c *gin.Context) {
	// The futur code…
}

func DeleteUser(c *gin.Context) {
	// The futur code…
}