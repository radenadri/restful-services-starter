package middlewares

import (
	"boilerplate/app/config"
	"boilerplate/app/db"
	"boilerplate/app/models"
	"boilerplate/app/utils"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

func EnableJWT() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(config.JWT_SECRET)},
		ErrorHandler: JwtError,
	})
}

func GenerateJWT(user models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	refreshToken := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	refreshClaims := refreshToken.Claims.(jwt.MapClaims)

	claims["user_id"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	refreshClaims["user_id"] = user.ID
	refreshClaims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()

	t, err := token.SignedString([]byte(config.JWT_SECRET))
	if err != nil {
		return "Error generating access token : ", err
	}

	r, err := refreshToken.SignedString([]byte(config.JWT_SECRET))
	if err != nil {
		return "Error generating refresh token : ", err
	}

	return t + "|" + r, nil
}

func GenerateRefreshJWT(user models.User) (string, error) {
	refreshToken := jwt.New(jwt.SigningMethodHS256)

	refreshClaims := refreshToken.Claims.(jwt.MapClaims)

	refreshClaims["user_id"] = user.ID
	refreshClaims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()

	r, err := refreshToken.SignedString([]byte(config.JWT_SECRET))
	if err != nil {
		return "Error generating refresh token : ", err
	}

	return r, nil
}

func FindUserByToken(c *fiber.Ctx) (*models.User, error) {
	// Get header "Authorization"
	auth := c.Request().Header.Peek("Authorization")

	if auth == nil {
		return nil, nil
	}

	// Split token
	token := strings.Split(string(auth), " ")

	if len(token) != 2 {
		return nil, nil
	}

	// Get token
	t := token[1]

	// Parse token
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(t, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWT_SECRET), nil
	})

	if err != nil {
		return nil, err
	}

	// Get user id
	userId := int(claims["user_id"].(float64))

	// Check if user exists
	var user models.User
	if err := db.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}

	// Return user
	return &user, nil
}

func JwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(utils.Response{
				Status:  false,
				Message: "Missing or malformed JWT",
				Data:    nil,
			})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(utils.Response{
			Status:  false,
			Message: "Invalid or expired JWT",
			Data:    nil,
		})
}
