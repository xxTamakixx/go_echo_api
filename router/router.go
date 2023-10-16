package router

import (
	"go_echo_api/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(pc controller.IPostController) *echo.Echo {
	e := echo.New()
	p := e.Group("/posts")
	p.GET("", pc.GetAllPosts)
	p.GET("/:id", pc.GetPostById)
	p.POST("", pc.CreatePost)
	p.PUT("/:id", pc.UpdatePost)
	p.DELETE("/:id", pc.DeletePost)
	return e
}
