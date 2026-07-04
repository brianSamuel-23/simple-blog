package internal

import (
	"context"

	"simple-blog.com/entity"
	"simple-blog.com/post/post"
	"simple-blog.com/user/user"
)

func (c *Comment) AddCommment(ctx context.Context, comment entity.Comment) (entity.Comment, error) {

	if comment.AuthorID != 0 {
		usr, err := c.usrRepo.GetOne(ctx, user.GetOneUserParam{ID: comment.AuthorID})

		if err != nil {
			return entity.Comment{}, err
		}

		comment.AuthorName = usr.Name
	}

	if _, err := c.pstRepo.GetOne(ctx, post.GetOnePostParam{ID: comment.PostID}); err != nil {
		return entity.Comment{}, err
	}

	cmt, err := c.cmtRepo.Save(ctx, comment)

	if err != nil {
		return entity.Comment{}, err
	}

	return cmt, nil

}
