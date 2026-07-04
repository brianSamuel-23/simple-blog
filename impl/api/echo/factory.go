package echo

import (
	"github.com/labstack/echo/v4"

	"simple-blog.com/comment"
	"simple-blog.com/impl/api/echo/internal"
	"simple-blog.com/impl/sanitizer"
	"simple-blog.com/post"
	"simple-blog.com/user"
)

type Post interface {
	CreatePost(c echo.Context) error
	GetPostDetail(c echo.Context) error
	ListPost(c echo.Context) error
	UpdatePost(c echo.Context) error
	DeletePost(c echo.Context) error
}

type Comment interface {
	AddComment(c echo.Context) error
	ListComment(c echo.Context) error
}

type User interface {
	Login(c echo.Context) error
	Register(c echo.Context) error
}

type API interface {
	Post
	Comment
	User
}

func New(pst post.Post, cmt comment.Comment, usr user.User, sanitizer sanitizer.Sanitizer) API {
	return internal.NewEcho(pst, cmt, usr, sanitizer)
}
