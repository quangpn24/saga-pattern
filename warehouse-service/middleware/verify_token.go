package middleware

import (
	"go-base/util"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func VerifyToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		token := c.Request().Header.Get("Authorization")
		if token == "" {
			c.Error(echo.NewHTTPError(http.StatusBadRequest, "You need access permission"))
		}

		newToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.NewHTTPError(http.StatusForbidden, "Unexpected signing method: %v", token.Header["alg"])
			}
			signature := []byte(util.JWT_SECRET_KEY)
			return signature, nil
		})

		if err != nil {
			return c.JSON(http.StatusForbidden, err.Error())
		}

		claims, ok := newToken.Claims.(jwt.MapClaims)
		if !ok || !newToken.Valid {
			return c.JSON(http.StatusForbidden, "couldn't parse claims")
		}

		//check expired time
		if int64(claims[util.EXP_AT].(float64)) < time.Now().Local().Unix() {
			return c.JSON(http.StatusUnauthorized, "token expired")
		}
		return next(c)
	}
}
