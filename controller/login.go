package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginGet render admin page
func LoginGet(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login.html", nil)
}

func LoginPost(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/admin")
	c.Abort()
}
