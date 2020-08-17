package main

import (
	"blogapp/src/middlewares"
	"blogapp/src/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

// HomePage handler
func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "rest api for a simple blog app",
	})
}

// // SpitOut Simple Func
// func SpitOut(c *gin.Context) {
// 	name := c.Query("name")
// 	age := c.Query("age")

// 	c.JSON(200, gin.H{
// 		"name": name,
// 		"age":  age,
// 	})
// }

// // SpitOut2 one more way to get variable information
// func SpitOut2(c *gin.Context) {
// 	name := c.Param("name")
// 	age := c.Param("age")

// 	c.JSON(200, gin.H{
// 		"name": name,
// 		"age":  age,
// 	})
// }

// // ReqBody struct
// type ReqBody struct {
// 	Name string
// 	Age  int
// }

// // SpitOut3 value extract from POST Req
// func SpitOut3(c *gin.Context) {

// 	body := c.Request.Body
// 	// value, _ := ioutil.ReadAll(body)

// 	var val ReqBody
// 	json.NewDecoder(body).Decode(&val)

// 	// fmt.Println(val.Name)
// 	// fmt.Println(val.Age)

// 	c.JSON(200, gin.H{
// 		"name": val.Name,
// 		"age":  val.Age,
// 	})
// }

func main() {

	fmt.Println("spinning up the server...")

	server := gin.Default()
	server.GET("/", HomePage)
	// server.GET("/spitout", SpitOut)
	// server.GET("/spitout2/:name/:age", SpitOut2)
	// server.POST("/spitout3", SpitOut3)

	BlogGroup := server.Group("blogs")
	{
		BlogGroup.GET("/", routes.GetAllBlogs)
		BlogGroup.POST("/", middlewares.AuthMiddleware(), routes.PostBlog)
		BlogGroup.GET("/:id", routes.GetBlog)
		BlogGroup.PUT("/:id", middlewares.AuthMiddleware(), routes.PutBlog)
		BlogGroup.DELETE("/:id", middlewares.AuthMiddleware(), routes.DeleteBlog)
	}

	server.Run()
}
