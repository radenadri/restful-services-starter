package controllers

import (
	"boilerplate/app/db"
	"boilerplate/app/models"
	"boilerplate/app/utils"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// ShowUser godoc
// @Summary      Get current user
// @Description  Get current user with all posts
// @Tags         User
// @Accept       json
// @Headers      Content-Type application/json
// @Param	 page query int true "Page number" default(1)
// @Param	 perPage query int true "Number of posts per page" default(10)
// @Param	 id path int true "Post ID"
// @Produce      json
// @Security	 ApiKeyAuth
// @Success      200  {array}   models.User
// @Failure      400  {object}  utils.Response
// @Failure      404  {object}  utils.Response
// @Failure      500  {object}  utils.Response
// @Router       /users/{id} [get]
func ShowUser(c *fiber.Ctx) error {

	// get query params page and perPage
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("perPage", "10"))

	// get users first, return error if not found
	var user models.User
	if err := db.DB.First(&user, c.Params("id")).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.Response{
			Status:  false,
			Message: "User not found",
			Data:    []interface{}{},
		})
	}

	// find posts by user id, and paginate
	var posts []models.Post
	db.DB.Where("user_id = ?", user.ID).Scopes(db.Paginate(page, perPage)).Order("id desc").Find(&posts)

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		Status:  true,
		Message: "Successfully fetched users with post",
		Data:    map[string]interface{}{"user": user, "posts": posts},
	})
}
