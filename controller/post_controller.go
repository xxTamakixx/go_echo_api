package controller

import (
	"go_echo_api/model"
	"go_echo_api/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IPostController interface {
	GetAllPosts(c echo.Context) error
	GetPostById(c echo.Context) error
	CreatePost(c echo.Context) error
	UpdatePost(c echo.Context) error
	DeletePost(c echo.Context) error
}

type postController struct {
	pu usecase.IPostUsecase
}

func NewPostController(pu usecase.IPostUsecase) IPostController {
	return &postController{pu}
}

func (pc *postController) GetAllPosts(c echo.Context) error {
	postsRes, err := pc.pu.GetAllPosts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, postsRes)
}

func (pc *postController) GetPostById(c echo.Context) error {
	id := c.Param("id")
	postId, _ := strconv.Atoi(id)
	postRes, err := pc.pu.GetPostById(uint(postId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, postRes)
}

func (pc *postController) CreatePost(c echo.Context) error {
	post := model.Post{}
	if err := c.Bind(&post); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	postRes, err := pc.pu.CreatePost(post)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, postRes)
}

func (pc *postController) UpdatePost(c echo.Context) error {
	id := c.Param("id")
	postId, _ := strconv.Atoi(id)

	post := model.Post{}
	if err := c.Bind(&post); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	postRes, err := pc.pu.UpdatePost(post, uint(postId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, postRes)
}

func (pc *postController) DeletePost(c echo.Context) error {
	id := c.Param("id")
	postId, _ := strconv.Atoi(id)

	err := pc.pu.DeletePost(uint(postId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
