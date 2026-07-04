package internal

import (
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"

	"simple-blog.com/user/user"
)

func (e *Echo) Register(c echo.Context) error {
	var req registerRequest
	if !bindAndValidate(c, &req) {
		return returnInvalidPayload(c)
	}

	req.Name = strings.TrimSpace(req.Name)
	req.Email = strings.ToLower(strings.TrimSpace(req.Email))

	if err := e.usr.Register(c.Request().Context(), user.RegisterRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}); err != nil {
		switch err {
		case user.ErrEmailAlreadyExists:
			return returnErrorResponse(c, http.StatusConflict, user.ErrEmailAlreadyExists.Error(), map[string]string{
				"field": "email",
				"error": user.ErrEmailAlreadyExists.Error(),
			})
		case user.ErrWeakPassword:
			return returnErrorResponse(c, http.StatusUnprocessableEntity, user.ErrWeakPassword.Error(), map[string]string{
				"field": "password",
				"error": user.ErrWeakPassword.Error(),
			})
		default:
			return returnErrorResponse(c, http.StatusInternalServerError, INTERNAL_SERVER_ERROR, map[string]string{
				"field": "-",
				"error": INTERNAL_SERVER_ERROR,
			})
		}
	}

	return returnSuccessResponse(c, http.StatusCreated, SUCCESSFUL_USER_REGISTRATION, nil, nil)
}

func (e *Echo) Login(c echo.Context) error {
	var req loginRequest
	if !bindAndValidate(c, &req) {
		return returnInvalidPayload(c)
	}

	req.Email = strings.ToLower(strings.TrimSpace(req.Email))

	res, err := e.usr.Login(c.Request().Context(), user.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		switch err {
		case user.ErrInvalidCredentials:
			return returnErrorResponse(c, http.StatusUnauthorized, user.ErrInvalidCredentials.Error(), map[string]string{
				"field": "email or password",
				"error": user.ErrInvalidCredentials.Error(),
			})
		default:
			return returnErrorResponse(c, http.StatusInternalServerError, INTERNAL_SERVER_ERROR, map[string]string{
				"field": "-",
				"error": INTERNAL_SERVER_ERROR,
			})
		}
	}

	return returnSuccessResponse(c, http.StatusOK, SUCCESSFUL_USER_LOGIN, loginResponse{
		AccessToken: res.AccessToken,
		ExpiresAt:   res.ExpiresAt.Format(time.RFC3339),
	}, nil)
}
