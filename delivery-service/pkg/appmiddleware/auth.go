package auth

import (
	"delivery-service/pkg/constant"
	"delivery-service/pkg/utils"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Authentication struct {
	SkipperPath []string
	KeyLookup   string
	AuthScheme  string
}

func NewAuthentication(keyLookup string, authScheme string, skipperPath []string) *Authentication {
	return &Authentication{
		SkipperPath: skipperPath,
		KeyLookup:   keyLookup,
		AuthScheme:  authScheme,
	}
}

func (a *Authentication) Middleware() echo.MiddlewareFunc {
	skipper := func(c echo.Context) bool {
		return utils.ContainFirst(a.SkipperPath, c.Path())
	}
	return middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		Skipper:    skipper,
		KeyLookup:  a.KeyLookup,
		AuthScheme: a.AuthScheme,
		Validator:  a.ValidateAccessToken,
	})
}

func (a *Authentication) ValidateAccessToken(token string, c echo.Context) (bool, error) {
	if token == "" {
		c.Error(echo.NewHTTPError(http.StatusBadRequest, "You need access permission"))
	}

	newToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, echo.NewHTTPError(http.StatusForbidden, "Unexpected signing method: %v", token.Header["alg"])
		}
		signature := []byte(constant.JWT_SECRET_KEY)
		return signature, nil
	})

	if err != nil {
		return false, c.JSON(http.StatusForbidden, err.Error())
	}

	claims, ok := newToken.Claims.(jwt.MapClaims)
	if !ok || !newToken.Valid {
		return false, c.JSON(http.StatusForbidden, "couldn't parse claims")
	}

	//check expired time
	if int64(claims["exp_at"].(float64)) < time.Now().Local().Unix() {
		return false, c.JSON(http.StatusUnauthorized, "token expired")
	}
	return true, nil
}
