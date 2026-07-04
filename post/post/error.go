package post

import "errors"

var (
	ErrPostNotFound = errors.New("Blog post not found!")
	ErrForbidenPost = errors.New("Blog post does not belong to the requester!")
)
