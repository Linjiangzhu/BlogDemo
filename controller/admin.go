package controller

import (
	"BlogDemo/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// adminIndexHandler render blog main page
func AdminIndexHandler(c *gin.Context) {
	data := model.RetriveAllPost()
	c.HTML(http.StatusOK, "admin/dashboard.html", data)
}
