package controllers

import (
	"boilerplate/app/config"
	pg "boilerplate/app/db"
	"boilerplate/app/middlewares"
	"boilerplate/app/models"
	"boilerplate/app/utils"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Login godoc
// @Summary      Perform login
// @Description  Login with email and password
// @Tags         Auth
// @Accept       json
// @Headers      Content-Type application/json
// @Produce      json
// @Param        request body models.LoginRequest true "Login request"
// @Success      200  {array}   models.LoginResponse
// @Failure      400  {object}  utils.Response
// @Failure      404  {object}  utils.Response
// @Failure      500  {object}  utils.Response
// @Router       /login [post]
func Login(c *fiber.Ctx) error {
	// Get and parse user input
	loginRequest := new(models.LoginRequest)
	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			Status:  false,
			Message: "Invalid request payload",
			Data:    nil,
		})
	}

	// Validate user input
	validationErrors := utils.GlobalValidator.Validate(loginRequest)
	if len(validationErrors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			Status:  false,
			Message: "Validation errors",
			Data:    validationErrors,
		})
	}

	// Check if user exists
	var user models.User
	if err := pg.DB.Where("email = ?", loginRequest.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.Response{
			Status:  false,
			Message: "User not found",
			Data:    nil,
		})
	}

	// Check if password is correct
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginRequest.PasswordHash)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			Status:  false,
			Message: "Invalid email or password",
			Data:    nil,
		})
	}

	// Generate JWT
	token, err := middlewares.GenerateJWT(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Response{
			Status:  false,
			Message: "Could not login user",
			Data:    nil,
		})
	}

	// explode token with | as delimiter
	tokenList := strings.Split(token, "|")

	// set cookie
	c.Cookie(&fiber.Cookie{
		Name:    "refreshToken",
		Value:   tokenList[1],
		Expires: time.Now().Add(time.Hour * 24),
	})

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		Status:  true,
		Message: "User logged in successfully",
		Data: models.LoginResponse{
			Token:        tokenList[0],
			RefreshToken: tokenList[1],
		},
	})
}

// Register godoc
// @Summary      Attempt register
// @Description  Register to the system
// @Tags         Auth
// @Accept       json
// @Headers      Content-Type application/json
// @Produce      json
// @Param        request body models.RegisterRequest true "Register request"
// @Success      200  {array}   models.User
// @Failure      400  {object}  utils.Response
// @Failure      404  {object}  utils.Response
// @Failure      500  {object}  utils.Response
// @Router       /register [post]
func Register(c *fiber.Ctx) error {
	// Get and parse user input
	registerRequest := new(models.RegisterRequest)
	if err := c.BodyParser(registerRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			Status:  false,
			Message: "Invalid request payload",
			Data:    nil,
		})
	}

	// Validate user input
	validationErrors := utils.GlobalValidator.Validate(registerRequest)
	if len(validationErrors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			Status:  false,
			Message: "Validation errors",
			Data:    validationErrors,
		})
	}

	// Generate password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerRequest.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Response{
			Status:  false,
			Message: "Could not register user",
			Data:    nil,
		})
	}

	// Create user
	user := models.User{
		Name:         registerRequest.Name,
		Email:        registerRequest.Email,
		PasswordHash: string(hashedPassword),
	}

	// Save user
	if err := pg.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Response{
			Status:  false,
			Message: "Could not save user",
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		Status:  true,
		Message: "User registered successfully",
		Data:    user,
	})
}

// RefreshToken godoc
// @Summary      Refresh a user access token
// @Description  Refresh a user access token and return the new access token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.RefreshResponse
// @Failure      400  {object}  utils.Response
// @Failure      404  {object}  utils.Response
// @Failure      500  {object}  utils.Response
// @Router       /refresh [post]
func RefreshToken(c *fiber.Ctx) error {
	// Get cookie
	cookie := c.Cookies("refreshToken")

	// Parse token
	token, err := jwt.ParseWithClaims(cookie, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWT_SECRET), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.Response{
			Status:  false,
			Message: "Could not refresh token",
			Data:    nil,
		})
	}

	// Get claims
	claims := token.Claims.(jwt.MapClaims)

	// Check if token is expired
	if !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.Response{
			Status:  false,
			Message: "Could not refresh token",
			Data:    nil,
		})
	}

	// Get user id
	userId := int(claims["user_id"].(float64))

	// Check if user exists
	var user models.User
	if err := pg.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.Response{
			Status:  false,
			Message: "Could not refresh token",
			Data:    nil,
		})
	}

	// Generate new token
	refreshToken, err := middlewares.GenerateRefreshJWT(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Response{
			Status:  false,
			Message: "Could not refresh token",
			Data:    nil,
		})
	}

	// set cookie
	c.Cookie(&fiber.Cookie{
		Name:    "refreshToken",
		Value:   refreshToken,
		Expires: time.Now().Add(time.Hour * 24),
	})

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		Status:  true,
		Message: "Token refreshed successfully",
		Data: models.RefreshResponse{
			RefreshToken: refreshToken,
		},
	})
}

// Profile godoc
// @Summary      Get profile
// @Description  Get current user profile
// @Tags         Auth
// @Accept       json
// @Headers      Content-Type application/json
// @Produce      json
// @Security	 ApiKeyAuth
// @Success      200  {array}   models.User
// @Failure      400  {object}  utils.Response
// @Failure      404  {object}  utils.Response
// @Failure      500  {object}  utils.Response
// @Router       /me [get]
func Profile(c *fiber.Ctx) error {
	// Get header "Authorization"
	auth := c.Request().Header.Peek("Authorization")

	// Split token
	splitToken := strings.Split(string(auth), "Bearer ")
	auth = []byte(splitToken[1])

	// Parse token
	token, err := jwt.ParseWithClaims(string(auth), jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWT_SECRET), nil
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Response{
			Status:  false,
			Message: "Error while parsing token",
			Data:    err.Error(),
		})
	}

	// Get user
	var user models.User
	if err := pg.DB.Where("id = ?", token.Claims.(jwt.MapClaims)["user_id"]).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.Response{
			Status:  false,
			Message: "User not found",
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		Status:  true,
		Message: "User profile retrieved successfully",
		Data:    user,
	})
}

// UpdateProfile godoc
// @Summary      Update profile
// @Description  Update current user profile
// @Tags         Auth
// @Accept       json
// @Headers      Content-Type application/json
// @Produce      json
// @Security	 ApiKeyAuth
// @Param        request body models.UpdateUserRequest true "Update profile request"
// @Success      200  {array}   models.User
// @Failure      400  {object}  utils.Response
// @Failure      404  {object}  utils.Response
// @Failure      500  {object}  utils.Response
// @Router       /me [put]
func UpdateProfile(c *fiber.Ctx) error {
	// Get header "Authorization"
	auth := c.Request().Header.Peek("Authorization")

	// Split token
	splitToken := strings.Split(string(auth), "Bearer ")
	auth = []byte(splitToken[1])

	// Parse token
	token, err := jwt.ParseWithClaims(string(auth), jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWT_SECRET), nil
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Response{
			Status:  false,
			Message: "Error while parsing token",
			Data:    err.Error(),
		})
	}

	// Get user input
	updateUserRequest := new(models.UpdateUserRequest)
	if err := c.BodyParser(updateUserRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			Status:  false,
			Message: "Invalid request payload",
			Data:    nil,
		})
	}

	// Validate user input
	validationErrors := utils.GlobalValidator.Validate(updateUserRequest)
	if len(validationErrors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			Status:  false,
			Message: "Validation errors",
			Data:    validationErrors,
		})
	}

	// Get user
	var user models.User
	if err := pg.DB.Where("id = ?", token.Claims.(jwt.MapClaims)["user_id"]).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.Response{
			Status:  false,
			Message: "User not found",
			Data:    nil,
		})
	}

	// Update user
	user.Name = updateUserRequest.Name
	if err := pg.DB.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Response{
			Status:  false,
			Message: "Error while updating user",
			Data:    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		Status:  true,
		Message: "User profile updated successfully",
		Data:    user,
	})
}
