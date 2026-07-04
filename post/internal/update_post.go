package internal

import (
	"context"

	"simple-blog.com/post/post"
)

func (p *Post) UpdatePost(ctx context.Context, req post.UpdatePostParam) error {
	exstPost, err := p.postRepo.GetOne(ctx, post.GetOnePostParam{ID: req.ID})

	if err != nil {
		return err
	}

	if exstPost.AuthorID != req.RequesterID {
		return post.ErrForbidenPost
	}

	exstPost.Content = req.Content
	exstPost.Title = req.Title

	_, err = p.postRepo.Save(ctx, exstPost)

	if err != nil {
		return err
	}

	return nil
}
