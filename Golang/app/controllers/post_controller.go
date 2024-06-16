package controllers

import (
	"boilerplate/app/db"
	"boilerplate/app/middlewares"
	"boilerplate/app/models"
	"boilerplate/app/utils"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// IndexPost godoc
// @Summary      Get all posts
// @Description  Get all posts and return them as a json
// @Tags         Post
// @Accept       json
// @Headers      Content-Type application/json
// @Param	 page query int true "Page number" default(1)
// @Param	 perPage query int true "Number of posts per page" default(10)
// @Produce      json
// @Security	 ApiKeyAuth
// @Success      200  {array}   models.Post
// @Failure      400  {object}  utils.Response
// @Failure      404  {object}  utils.Response
// @Failure      500  {object}  utils.Response
// @Router       /posts [get]
func IndexPost(c *fiber.Ctx) error {
	// Get all posts, get query params page and perPage
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("perPage", "10"))

	// Get all posts and paginate
	var posts []models.Post
	db.DB.Scopes(db.Paginate(page, perPage)).Order("id desc").Find(&posts)

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		Status:  true,
		Message: "Successfully fetched posts",
		Data:    posts,
	})
}

// ShowPost godoc
// @Summary      Get post by ID
// @Description  Get post by ID and return it as a json
// @Tags         Post
// @Accept       json
// @Headers      Content-Type application/json
// @Param	 id path int true "Post ID"
// @Produce      json
// @Security	 ApiKeyAuth
// @Success      200  {array}   models.Post
// @Failure      400  {object}  utils.Response
// @Failure      404  {object}  utils.Response
// @Failure      500  {object}  utils.Response
// @Router       /posts/{id} [get]
func ShowPost(c *fiber.Ctx) error {
	var post models.Post

	// Check if post exists
	if err := db.DB.First(&post, c.Params("id")).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.Response{
			Status:  false,
			Message: "Post not found",
			Data:    []interface{}{},
		})
	}

	db.DB.Where("id = ?", c.Params("id")).First(&post)

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		Status:  true,
		Message: "Successfully fetched post",
		Data:    post,
	})
}

// StorePost godoc
// @Summary	 Create new post
// @Description  Create a new post and return it
// @Tags         Post
// @Accept       json
// @Headers      Content-Type application/json
// @Produce      json
// @Security	 ApiKeyAuth
// @Success      200  {array}   models.Post
// @Failure      400  {object}  utils.Response
// @Failure      404  {object}  utils.Response
// @Failure      500  {object}  utils.Response
// @Router       /posts [post]
func StorePost(c *fiber.Ctx) error {
	// get and parse request body
	postRequest := new(models.CreatePostRequest)
	if err := c.BodyParser(postRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			Status:  false,
			Message: "Invalid request payload",
			Data:    []interface{}{},
		})
	}

	// Validate user input
	validationErrors := utils.GlobalValidator.Validate(postRequest)
	if len(validationErrors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			Status:  false,
			Message: "Validation errors",
			Data:    validationErrors,
		})
	}

	// Get authenticated user from token
	user, err := middlewares.FindUserByToken(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.Response{
			Status:  false,
			Message: "Unauthorized",
			Data:    []interface{}{},
		})
	}

	// Create new post
	post := models.Post{
		Title: postRequest.Title,
		Body:  postRequest.Body,
		User:  user,
	}

	if err := db.DB.Create(&post).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Response{
			Status:  false,
			Message: "Failed to create post",
			Data: map[string]interface{}{
				"error": err.Error(),
			},
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		Status:  true,
		Message: "Successfully created post",
		Data:    post,
	})

}

// UpdatePost godoc
// @Summary      Update post
// @Description  Update post and return it
// @Tags         Post
// @Accept       json
// @Headers      Content-Type application/json
// @Param	 id path int true "Post ID"
// @Produce      json
// @Security	 ApiKeyAuth
// @Success      200  {array}   models.Post
// @Failure      400  {object}  utils.Response
// @Failure      404  {object}  utils.Response
// @Failure      500  {object}  utils.Response
// @Router       /posts/{id} [put]
func UpdatePost(c *fiber.Ctx) error {

	// get and parse request body
	postRequest := new(models.UpdatePostRequest)
	if err := c.BodyParser(postRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			Status:  false,
			Message: "Invalid request payload",
			Data:    []interface{}{},
		})
	}

	// validate user input
	validationErrors := utils.GlobalValidator.Validate(postRequest)
	if len(validationErrors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			Status:  false,
			Message: "Validation errors",
			Data:    validationErrors,
		})
	}

	// find post by id
	var post models.Post
	if err := db.DB.First(&post, c.Params("id")).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.Response{
			Status:  false,
			Message: "Post not found",
			Data:    []interface{}{},
		})
	}

	// update post
	if err := db.DB.Model(&post).Updates(postRequest).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Response{
			Status:  false,
			Message: "Failed to update post",
			Data: map[string]interface{}{
				"error": err.Error(),
			},
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		Status:  true,
		Message: "Successfully updated post",
		Data:    post,
	})

}

// DestroyPost godoc
// @Summary      Delete post
// @Description  Delete post and return the success deletion message
// @Tags         Post
// @Accept       json
// @Headers      Content-Type application/json
// @Param	 id path int true "Post ID"
// @Produce      json
// @Security	 ApiKeyAuth
// @Success      200  {array}   models.Post
// @Failure      400  {object}  utils.Response
// @Failure      404  {object}  utils.Response
// @Failure      500  {object}  utils.Response
// @Router       /posts/{id} [delete]
func DestroyPost(c *fiber.Ctx) error {

	// find post by id
	var post models.Post
	if err := db.DB.First(&post, c.Params("id")).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.Response{
			Status:  false,
			Message: "Post not found",
			Data:    []interface{}{},
		})
	}

	// delete post
	if err := db.DB.Delete(&post).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Response{
			Status:  false,
			Message: "Failed to delete post",
			Data: map[string]interface{}{
				"error": err.Error(),
			},
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		Status:  true,
		Message: "Successfully deleted post",
		Data:    map[string]interface{}{},
	})
}
