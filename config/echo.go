package config

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"simple-blog.com/contextkey"
)

type publicRoute struct{ method, path string }

var publicRoutes = []publicRoute{
	{http.MethodPost, "/register"},
	{http.MethodGet, "/posts/:id/comments"},
	{http.MethodPost, "/posts/:id/comments"},
	{http.MethodPost, "/login"},
}

func shouldSkipAuth(c echo.Context) bool {
	if c.Request().URL.Path == "/healthz" {
		return true
	}
	for _, r := range publicRoutes {
		if c.Request().Method == r.method && c.Path() == r.path {
			return true
		}
	}
	return false
}

type CustomValidator struct {
	validator *validator.Validate
}

func NewValidator() *CustomValidator {
	v := validator.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" || name == "" {
			return fld.Name
		}
		return name
	})

	return &CustomValidator{validator: v}
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			fields := make(map[string]string, len(ve))
			for _, fe := range ve {
				fields[fe.Field()] = validationMessage(fe)
			}
			return echo.NewHTTPError(http.StatusUnprocessableEntity, fields)
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
func NewEcho() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Validator = NewValidator()
	e.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(jwtSecret()),
		SigningMethod: echojwt.AlgorithmHS256,
		Skipper:       shouldSkipAuth,
	}))

	e.Use(userContextMiddleware())
	return e
}

func validationMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "is required"
	case "email":
		return "must be a valid email address"
	case "min":
		return fmt.Sprintf("must be at least %s characters", fe.Param())
	case "max":
		return fmt.Sprintf("must be at most %s characters", fe.Param())
	default:
		return fmt.Sprintf("failed validation on '%s'", fe.Tag())
	}
}

func claimsFromContext(c echo.Context) (jwt.MapClaims, bool) {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok || token == nil {
		return nil, false
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	return claims, ok
}
func userContextMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if shouldSkipAuth(c) {
				return next(c)
			}

			claims, ok := claimsFromContext(c)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "No token found")
			}

			if sub, ok := claims["sub"].(string); ok {
				if id, err := strconv.Atoi(sub); err == nil {
					c.Set(contextkey.REQUESTER_ID, id)
				}
			}

			if name, ok := claims["name"].(string); ok {
				c.Set(contextkey.REQUESTER_NAME, name)
			}
			return next(c)
		}
	}
}
