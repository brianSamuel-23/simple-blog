package internal

import (
	"context"

	"simple-blog.com/comment/comment"
)

func (c *Comment) ListComment(ctx context.Context, param comment.ListCommentParam) (comment.GetPaginatedCommentResponse, error) {
	return c.cmtRepo.GetPaginated(ctx, comment.GetPaginatedCommentParam{PostID: param.PostID, Page: param.Page, PerPage: param.PerPage, Order: param.Order, OrderBy: param.Order})
}
