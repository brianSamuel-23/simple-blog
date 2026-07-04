package internal

import (
	"context"

	"simple-blog.com/post/post"
)

func (p *Post) DeletePost(ctx context.Context, ID int, requesterID int) error {
	exstPost, err := p.postRepo.GetOne(ctx, post.GetOnePostParam{ID: ID})

	if err != nil {
		return err
	}

	if exstPost.AuthorID != requesterID {
		return post.ErrForbidenPost
	}

	err = p.postRepo.Delete(ctx, ID)

	if err != nil {
		return err
	}

	return nil
}
