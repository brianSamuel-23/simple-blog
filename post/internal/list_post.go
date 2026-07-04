package internal

import (
	"context"

	"simple-blog.com/post/post"
)

func (p *Post) ListPost(ctx context.Context, param post.ListPostParam) (post.GetPaginatedPostResponse, error) {
	return p.postRepo.GetPaginated(ctx, post.GetPaginatedPostParam{Page: param.Page, PerPage: param.PerPage, Order: param.Order, OrderBy: param.OrderBy})
}
