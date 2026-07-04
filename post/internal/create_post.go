package internal

import (
	"context"

	"simple-blog.com/entity"
	"simple-blog.com/user/user"
)

func (p *Post) CreatePost(ctx context.Context, post entity.Post) (entity.Post, error) {

	usr, err := p.usrRepo.GetOne(ctx, user.GetOneUserParam{ID: post.AuthorID})

	if err != nil {
		return entity.Post{}, err
	}

	post.AuthorName = usr.Name

	return p.postRepo.Save(ctx, post)
}
