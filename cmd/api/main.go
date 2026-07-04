package main

import (
	"simple-blog.com/comment"
	"simple-blog.com/config"
	httpapi "simple-blog.com/impl/api/echo"
	"simple-blog.com/impl/hasher"
	"simple-blog.com/impl/repo/mysql"
	"simple-blog.com/impl/sanitizer"
	"simple-blog.com/post"
	"simple-blog.com/user"
)

func main() {

	db := config.NewMysql()

	repo := mysql.New(db)

	post := post.New(repo)
	cmt := comment.New(repo)

	sanitizer := sanitizer.New()
	hasher := hasher.New()
	tokens := config.NewJWT()
	usr := user.New(repo, hasher, tokens)
	api := httpapi.New(post, cmt, usr, sanitizer)

	e := config.NewEcho()

	e.POST("/posts", api.CreatePost)
	e.GET("/posts", api.ListPost)
	e.GET("/posts/:id", api.GetPostDetail)
	e.PUT("/posts/:id", api.UpdatePost)
	e.DELETE("/posts/:id", api.DeletePost)

	e.POST("/posts/:id/comments", api.AddComment)
	e.GET("/posts/:id/comments", api.ListComment)

	e.POST("/register", api.Register)
	e.POST("/login", api.Login)

	e.Logger.Fatal(e.Start(":8080"))
}
