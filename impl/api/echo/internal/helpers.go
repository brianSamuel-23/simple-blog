package internal

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type errorResponse struct {
	Message string `json:"message"`
}

func bindAndValidate(c echo.Context, req any) bool {
	if err := c.Bind(req); err != nil {
		return false
	}

	if err := c.Validate(req); err != nil {
		return false
	}
	return true
}

func returnSuccessResponse(c echo.Context, statusCode int, message string, data any, metadata any) error {
	return c.JSON(statusCode, &ApiResponseModel{
		Message:  message,
		Data:     data,
		Metadata: metadata,
	})
}

func returnErrorResponse(c echo.Context, statusCode int, message string, errors any) error {
	return c.JSON(statusCode, &ApiResponseModel{
		Message: message,
		Error:   errors,
	})
}

func returnInvalidPayload(c echo.Context) error {
	return c.JSON(http.StatusBadRequest, &ApiResponseModel{
		Message: INVALID_REQUEST_PAYLOAD,
		Error: map[string]string{
			"field": "payload body",
			"error": INVALID_REQUEST_PAYLOAD,
		},
	})
}
