package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// HomePage handler
func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "rest api for a simple blog app",
	})
}

// SpitOut Simple Func
func SpitOut(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

// SpitOut2 one more way to get variable information
func SpitOut2(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

// ReqBody struct
type ReqBody struct {
	Name string
	Age  int
}

// SpitOut3 value extract from POST Req
func SpitOut3(c *gin.Context) {

	body := c.Request.Body
	// value, _ := ioutil.ReadAll(body)

	var val ReqBody
	json.NewDecoder(body).Decode(&val)

	// fmt.Println(val.Name)
	// fmt.Println(val.Age)

	c.JSON(200, gin.H{
		"name": val.Name,
		"age":  val.Age,
	})
}

// Blog interface
type Blog struct {
	Title    string
	Author   string
	BlogData string
	ID       string
}

// AllBlogs global variable
// var AllBlogs []Blog
var AllBlogs = make(map[string]Blog)

// GetAllBlogs ==> get all blogs
func GetAllBlogs(c *gin.Context) {

	c.JSON(200, gin.H{
		"data": AllBlogs,
	})
}

// GetBlog ==> get a particular blog
func GetBlog(c *gin.Context) {
	BlogID := c.Param("id")

	if val, ok := AllBlogs[BlogID]; ok {
		c.JSON(200, gin.H{
			"data": val,
		})
	} else {
		c.JSON(200, gin.H{
			"data": "Sorry... No such blog exists",
		})
	}

}

// PostBlog ==> post a new blog
func PostBlog(c *gin.Context) {
	body := c.Request.Body

	var val Blog
	json.NewDecoder(body).Decode(&val)

	UUIDNew := uuid.New()
	val.ID = strings.Replace(UUIDNew.String(), "-", "", -1)

	AllBlogs[val.ID] = val

	c.JSON(201, gin.H{
		"data": val,
	})
}

// PutBlog ==> edit a particular blog
func PutBlog(c *gin.Context) {
	BlogID := c.Param("id")

	if OldVal, ok := AllBlogs[BlogID]; ok {
		body := c.Request.Body
		var NewVal Blog
		json.NewDecoder(body).Decode(&NewVal)

		if len(NewVal.Title) != 0 {
			OldVal.Title = NewVal.Title
		}

		if len(NewVal.Author) != 0 {
			OldVal.Author = NewVal.Author
		}

		if len(NewVal.BlogData) != 0 {
			OldVal.BlogData = NewVal.BlogData
		}

		AllBlogs[BlogID] = OldVal

		c.JSON(202, gin.H{
			"data": OldVal,
		})
	} else {
		c.JSON(204, gin.H{
			"data": "Sorry... No such blog exists",
		})
	}

}

// DeleteBlog ==> delete a particular blog
func DeleteBlog(c *gin.Context) {
	BlogID := c.Param("id")

	if _, ok := AllBlogs[BlogID]; ok {
		delete(AllBlogs, BlogID)

		c.JSON(202, gin.H{
			"data": "Deleted Successfully",
		})
	} else {
		c.JSON(204, gin.H{
			"data": "Sorry... No such blog exists",
		})
	}
}

func main() {

	fmt.Println("spinning up the server...")

	server := gin.Default()
	server.GET("/", HomePage)
	server.GET("/spitout", SpitOut)
	server.GET("/spitout2/:name/:age", SpitOut2)
	server.POST("/spitout3", SpitOut3)

	server.GET("/all_blogs", GetAllBlogs)
	server.POST("/new_blog", PostBlog)
	server.GET("/blog/:id", GetBlog)
	server.PUT("/blog/:id", PutBlog)
	server.DELETE("/blog/:id", DeleteBlog)

	server.Run()
}
