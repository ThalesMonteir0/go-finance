package handlers

import (
	"errors"
	"financial-go/internal/entity"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"strings"
)

const (
	SECRET_KEY = "SECRET_KEY"
)

func JwtMiddleware(next, stop echo.HandlerFunc, userRepo entity.UserRepositoryInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		err, celNumber := validJwtAndReturnCelClaim(c)
		if err != nil {
			return stop(c)
		}

		user, errFind := userRepo.FindUserByCellphone(celNumber)
		if errFind != nil {
			return stop(c)
		}

		c.Set("userID", user.ID)

		return next(c)
	}
}

func getHeaderAuthorization(c echo.Context) (string, error) {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return "", errors.New("token is null")
	}

	return strings.TrimPrefix(token, "Bearer "), nil
}

func validJwtAndReturnCelClaim(c echo.Context) (error, string) {
	secretKey := os.Getenv(SECRET_KEY)
	tokenString, err := getHeaderAuthorization(c)
	if err != nil {
		return err, ""
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secretKey), nil
		}
		return nil, errors.New("invalid token")
	})
	if err != nil {
		return err, ""
	}

	if !token.Valid {
		return err, ""
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("claims is invalid"), ""
	}

	return nil, claims["cel"].(string)
}

func SomeErrorHandler(c echo.Context) error {
	return c.JSON(
		http.StatusInternalServerError,
		map[string]any{"message": "Something went wrong !"},
	)
}
