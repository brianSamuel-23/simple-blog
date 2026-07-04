package internal

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"

	"simple-blog.com/comment/comment"
	"simple-blog.com/contextkey"
	"simple-blog.com/entity"
	"simple-blog.com/post/post"
	"simple-blog.com/user/user"
)

func (e *Echo) AddComment(c echo.Context) error {
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return returnErrorResponse(c, http.StatusUnprocessableEntity, INVALID_POST_ID, map[string]string{
			"field": "id",
			"error": INVALID_POST_ID,
		})
	}

	var req addCommentRequest
	if !bindAndValidate(c, &req) {
		return returnInvalidPayload(c)
	}

	authorID, _ := c.Get(contextkey.REQUESTER_ID).(int)

	res, err := e.cmt.AddCommment(c.Request().Context(), entity.Comment{
		PostID:   postID,
		AuthorID: authorID,
		Content:  e.sanitizer.Sanitize(req.Content),
	})
	if err != nil {
		switch err {
		case post.ErrPostNotFound:
			return returnErrorResponse(c, http.StatusBadRequest, post.ErrPostNotFound.Error(), map[string]string{
				"field": "post id",
				"error": post.ErrPostNotFound.Error(),
			})
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

	return returnSuccessResponse(c, http.StatusCreated, SUCCESSFUL_COMMENT_CREATION, map[string]string{
		"id": strconv.Itoa(res.ID),
	}, nil)
}

func (e *Echo) ListComment(c echo.Context) error {
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse{Message: "invalid id"})
	}

	var page int
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

	res, err := e.cmt.ListComment(c.Request().Context(), comment.ListCommentParam{
		PostID:  postID,
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

	comments := make([]getCommentListModel, len(res.Comments))
	for i, cm := range res.Comments {
		comments[i] = getCommentListModel{
			ID:         cm.ID,
			PostID:     cm.PostID,
			AuthorID:   cm.AuthorID,
			AuthorName: cm.AuthorName,
			Content:    cm.Content,
			CreatedAt:  cm.CreatedAt.Format(time.RFC3339),
		}
	}

	totalPage := (res.TotalData + perPage - 1) / perPage
	return returnSuccessResponse(c, http.StatusOK, SUCCESSFUL_COMMENT_LIST_FETCHING, comments, map[string]int{
		"page":       page,
		"per_page":   perPage,
		"total_data": res.TotalData,
		"total_page": totalPage,
	})
}
