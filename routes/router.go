package routes

import (
	"BlogDemo/controller"
	"BlogDemo/middleware"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

// NewRouter return a new gin router
func NewRouter() *gin.Engine {
	r = gin.Default()
	r.Use(middleware.MethodOverride(r))
	r.LoadHTMLGlob("./view/**/*")

	r.Static("/assets", "./assets")
	r.GET("/", controller.IndexHandler)
	r.GET("/blog", controller.BlogHandler)
	r.POST("/blog", controller.NewBlogHandler)

	r.GET("/blog/:postid", controller.PostHandler)
	r.DELETE("/blog/:postid", controller.DeletePostHandler)
	r.PUT("/blog/:postid", controller.EditPostHandler)
	r.GET("/blog/:postid/edit", controller.EditPageShow)

	r.GET("/login", controller.LoginGet)
	r.POST("/login", controller.LoginPost)

	r.GET("/admin", controller.AdminIndexHandler)
	r.GET("/admin/blog/new", controller.GetBlogNew)
	// r.GET("/blog/api/post/:postid")

	return r
}
