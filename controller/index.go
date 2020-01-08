package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// IndexHandler render the index page
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "home/index.html", nil)
}
