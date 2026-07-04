package internal

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"

	"simple-blog.com/contextkey"
	"simple-blog.com/entity"
	"simple-blog.com/post/post"
	"simple-blog.com/user/user"
)

// TODO : REQUESTER ID TO BE SWITCHED WITH LOGGED IN USER CLAIM
func (e *Echo) CreatePost(c echo.Context) error {
	var req createPostRequest
	if !bindAndValidate(c, &req) {
		return returnInvalidPayload(c)
	}

	res, err := e.pst.CreatePost(c.Request().Context(), entity.Post{
		Title:    req.Title,
		Content:  e.sanitizer.Sanitize(req.Content),
		AuthorID: c.Get(contextkey.REQUESTER_ID).(int),
	})
	if err != nil {
		switch err {
		case user.ErrUserNotFound:
			return returnErrorResponse(c, http.StatusBadRequest, user.ErrUserNotFound.Error(), map[string]string{
				"field": "logged in user_id",
				"error": user.ErrUserNotFound.Error(),
			})
		default:
			return returnErrorResponse(c, http.StatusInternalServerError, INTERNAL_SERVER_ERROR, map[string]string{
				"field": "-",
				"error": INTERNAL_SERVER_ERROR,
			})
		}

	}

	return returnSuccessResponse(c, http.StatusCreated, SUCCESSFUL_BLOG_POST_CREATION, map[string]string{
		"id": strconv.Itoa(res.ID),
	}, nil)
}

func (e *Echo) GetPostDetail(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return returnErrorResponse(c, http.StatusUnprocessableEntity, INVALID_POST_ID, map[string]string{
			"field": "id",
			"error": INVALID_POST_ID,
		})
	}

	postRes, err := e.pst.GetPost(c.Request().Context(), id)
	if err != nil {
		switch err {
		case post.ErrPostNotFound:
			return returnErrorResponse(c, http.StatusBadRequest, post.ErrPostNotFound.Error(), map[string]string{
				"field": "post id",
				"error": post.ErrPostNotFound.Error(),
			})
		default:
			return returnErrorResponse(c, http.StatusInternalServerError, INTERNAL_SERVER_ERROR, map[string]string{
				"field": "-",
				"error": INTERNAL_SERVER_ERROR,
			})
		}
	}

	res := getPostDetailResponse{
		ID:         postRes.ID,
		Title:      postRes.Title,
		Content:    postRes.Content,
		AuthorName: postRes.AuthorName,
		CreatedAt:  postRes.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  postRes.UpdatedAt.Format(time.RFC3339),
	}
	return returnSuccessResponse(c, http.StatusOK, SUCCESSFUL_BLOG_POST_DETAIL_FETCHING, res, nil)
}

func (e *Echo) ListPost(c echo.Context) error {

	var page int
	var err error
	if c.QueryParam("page") == "" {
		page = 1
	} else {
		page, err = strconv.Atoi(c.QueryParam("page"))
		if err != nil {
			return returnErrorResponse(c, http.StatusUnprocessableEntity, INVALID_PAGE_PARAMETER, map[string]string{
				"field": "page",
				"error": INVALID_PAGE_PARAMETER,
			})
		}

	}

	var perPage int
	if c.QueryParam("per_page") == "" {
		perPage = 5
	} else {
		perPage, err = strconv.Atoi(c.QueryParam("per_page"))

		if err != nil {
			returnErrorResponse(c, http.StatusUnprocessableEntity, INVALID_PER_PAGE_PARAMETER, map[string]string{
				"field": "per_page",
				"error": INVALID_PER_PAGE_PARAMETER,
			})
		}
	}

	res, err := e.pst.ListPost(c.Request().Context(), post.ListPostParam{
		Page:    page,
		PerPage: perPage,
		Order:   c.QueryParam("order"),
		OrderBy: c.QueryParam("order_by"),
	})
	if err != nil {
		return returnErrorResponse(c, http.StatusInternalServerError, INTERNAL_SERVER_ERROR, map[string]string{
			"field": "-",
			"error": INTERNAL_SERVER_ERROR,
		})
	}

	posts := make([]getPostListModel, len(res.Posts))
	for i, p := range res.Posts {
		posts[i] = getPostListModel{
			ID:         p.ID,
			Title:      p.Title,
			AuthorID:   p.AuthorID,
			AuthorName: p.AuthorName,
			CreatedAt:  p.CreatedAt.Format(time.RFC3339),
			UpdatedAt:  p.UpdatedAt.Format(time.RFC3339),
		}
	}

	totalPage := (res.TotalData + perPage - 1) / perPage

	return returnSuccessResponse(c, http.StatusOK, SUCCESSFUL_BLOG_POST_LIST_FETCHING, posts, map[string]int{
		"page":       page,
		"per_page":   perPage,
		"total_page": totalPage,
		"total_data": res.TotalData,
	})
}

func (e *Echo) UpdatePost(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return returnErrorResponse(c, http.StatusUnprocessableEntity, INVALID_POST_ID, map[string]string{
			"field": "id",
			"error": INVALID_POST_ID,
		})
	}

	var req updatePostRequest
	if !bindAndValidate(c, &req) {
		return returnInvalidPayload(c)
	}

	if err := e.pst.UpdatePost(c.Request().Context(), post.UpdatePostParam{
		ID:          id,
		Title:       req.Title,
		Content:     req.Content,
		RequesterID: c.Get(contextkey.REQUESTER_ID).(int),
	}); err != nil {
		switch err {
		case post.ErrPostNotFound:
			return returnErrorResponse(c, http.StatusBadRequest, post.ErrPostNotFound.Error(), map[string]string{
				"field": "post id",
				"error": post.ErrPostNotFound.Error(),
			})
		case post.ErrForbidenPost:
			return returnErrorResponse(c, http.StatusForbidden, post.ErrForbidenPost.Error(), map[string]string{
				"field": "logged in user_id",
				"error": post.ErrForbidenPost.Error(),
			})
		default:
			return returnErrorResponse(c, http.StatusInternalServerError, INTERNAL_SERVER_ERROR, map[string]string{
				"field": "-",
				"error": INTERNAL_SERVER_ERROR,
			})
		}

	}

	return returnSuccessResponse(c, http.StatusOK, SUCCESSFUL_BLOG_POST_UPDATE, map[string]string{
		"id": strconv.Itoa(id),
	}, nil)
}

func (e *Echo) DeletePost(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return returnErrorResponse(c, http.StatusUnprocessableEntity, INVALID_POST_ID, map[string]string{
			"field": "id",
			"error": INVALID_POST_ID,
		})
	}

	if err := e.pst.DeletePost(c.Request().Context(), id, c.Get(contextkey.REQUESTER_ID).(int)); err != nil {
		switch err {
		case post.ErrPostNotFound:
			return returnErrorResponse(c, http.StatusBadRequest, post.ErrPostNotFound.Error(), map[string]string{
				"field": "post id",
				"error": post.ErrPostNotFound.Error(),
			})
		case post.ErrForbidenPost:
			return returnErrorResponse(c, http.StatusForbidden, post.ErrForbidenPost.Error(), map[string]string{
				"field": "logged in user_id",
				"error": post.ErrForbidenPost.Error(),
			})
		default:
			return returnErrorResponse(c, http.StatusInternalServerError, INTERNAL_SERVER_ERROR, map[string]string{
				"field": "-",
				"error": INTERNAL_SERVER_ERROR,
			})
		}
	}

	return returnSuccessResponse(c, http.StatusOK, SUCCESSFUL_BLOG_POST_DELETTION, nil, nil)
}
