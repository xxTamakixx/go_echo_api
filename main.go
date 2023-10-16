package main

import (
	"go_echo_api/controller"
	"go_echo_api/db"
	"go_echo_api/repository"
	"go_echo_api/router"
	"go_echo_api/usecase"
)

func main() {
	db := db.NewDB()
	postRepository := repository.NewPostRepository(db)
	postUsecase := usecase.NewPostUsecase(postRepository)
	postController := controller.NewPostController(postUsecase)
	e := router.NewRouter(postController)
	e.Logger.Fatal(e.Start(":8080"))
}
