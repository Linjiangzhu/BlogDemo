package controller

import (
	"BlogDemo/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// BlogHandler render blog main page
func BlogHandler(c *gin.Context) {
	data := model.RetriveAllPost()
	c.HTML(http.StatusOK, "blog/blog.html", data)
}

func GetBlogNew(c *gin.Context) {
	c.HTML(http.StatusOK, "blog/new.html", nil)
}

func PostHandler(c *gin.Context) {
	id := c.Param("postid")
	post, err := model.RetrivePostByID(id)
	if err != nil {
		fmt.Println(err.Error())
		c.Redirect(http.StatusMovedPermanently, "/blog")
		c.Abort()
	}
	c.HTML(http.StatusOK, "blog/post.html", post)
}

func NewBlogHandler(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	err := model.CreatePost(title, content)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.Redirect(http.StatusMovedPermanently, "/admin")
	c.Abort()
}

func DeletePostHandler(c *gin.Context) {
	id := c.Param("postid")
	err := model.DeletePost(id)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.Redirect(http.StatusMovedPermanently, "/admin")
	c.Abort()
}

func EditPageShow(c *gin.Context) {
	id := c.Param("postid")
	post, err := model.RetrivePostByID(id)
	if err != nil {
		fmt.Println(err.Error())
		c.Redirect(http.StatusMovedPermanently, "/admin")
		c.Abort()
	}
	c.HTML(http.StatusOK, "blog/edit.html", post)

}

func EditPostHandler(c *gin.Context) {
	id := c.Param("postid")
	title := c.PostForm("title")
	content := c.PostForm("content")
	err := model.EditPost(id, title, content)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.Redirect(http.StatusMovedPermanently, "/admin")
	c.Abort()
}
