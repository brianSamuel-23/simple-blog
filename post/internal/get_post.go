package internal

import (
	"context"

	"simple-blog.com/entity"
	"simple-blog.com/post/post"
)

func (p *Post) GetPost(ctx context.Context, ID int) (entity.Post, error) {
	return p.postRepo.GetOne(ctx, post.GetOnePostParam{ID: ID})
}
