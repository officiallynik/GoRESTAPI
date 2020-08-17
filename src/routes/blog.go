package routes

import (
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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
