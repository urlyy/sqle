package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/actiontech/sqle/sqle/utils"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// JWTTokenAdapter is a `echo` middleware,　by rewriting the header, the jwt token support header
// "Authorization: {token}" and "Authorization: Bearer {token}".
func JWTTokenAdapter() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth := c.Request().Header.Get(echo.HeaderAuthorization)
			if auth != "" && !strings.HasPrefix(auth, middleware.DefaultJWTConfig.AuthScheme) {
				c.Request().Header.Set(echo.HeaderAuthorization,
					fmt.Sprintf("%s %s", middleware.DefaultJWTConfig.AuthScheme, auth))
			}
			return next(c)
		}
	}
}

var errAuditPlanMisMatch = errors.New("audit plan name don't match the token")

// ScannerVerifier is a `echo` middleware. Every audit plan should be
// scanner-scoped which means that scanner-A should not push SQL to scanner-B.
func ScannerVerifier() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// JWT parser expect no 'Bearer' ahead of token, so
			// we cut the leading auth schema.
			auth := c.Request().Header.Get(echo.HeaderAuthorization)
			parts := strings.Split(auth, " ")
			token := parts[0]
			if len(parts) == 2 {
				token = parts[1]
			}

			apnInToken, err := utils.ParseAuditPlanName(token)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}

			apnInParam := c.Param("audit_plan_name")
			if apnInToken != apnInParam {
				return echo.NewHTTPError(http.StatusInternalServerError, errAuditPlanMisMatch.Error())
			}
			return next(c)
		}
	}
}

var JWTConfig = middleware.JWTConfig{
	ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
		keyFunc := func(t *jwt.Token) (interface{}, error) {
			if t.Method.Alg() != "HS256" {
				return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
			}
			return []byte(utils.JWTSecret), nil
		}

		// claims are of type `jwt.MapClaims` when token is created with `jwt.Parse`
		token, err := jwt.Parse(auth, keyFunc)
		if err != nil {
			return nil, err
		}
		if !token.Valid {
			return nil, errors.New("invalid token")
		}
		return token, nil
	},
}
