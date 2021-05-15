package http

import (
	"fmt"

	_ "github.com/amartery/trinity-task/docs"
	"github.com/amartery/trinity-task/pkg/app"
	"github.com/amartery/trinity-task/pkg/app/models"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type AppHandlers struct {
	use app.Usecase
}

func NewAppHandlers(usecase app.Usecase) *AppHandlers {
	return &AppHandlers{
		use: usecase,
	}
}

func (h *AppHandlers) ConfigureRoutes(router fiber.Router) {
	router.Use(cors.New())
	router.Get("/api/v1/get-active-today", h.handlerActive)
	router.Post("/api/v1/add-user", h.handlerAddUser)
	router.Post("/api/v1/add-comment", h.handlerAddComment)
	router.Post("/api/v1/add-like", h.handlerAddLike)
	router.Get("/swagger/*", swagger.Handler)

}

// @Summary Get Today Activities
// @Tags api
// @Description getting users who have liked or commented today
// @ID get-active
// @Produce  json
// @Success 200 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Failure default {object} models.ErrorResponse
// @Router /api/v1/get-active-today [get]
func (h *AppHandlers) handlerActive(c *fiber.Ctx) error {
	result, err := h.use.GetTodayActivityFull()
	if err != nil {
		return c.Status(500).JSON(
			&fiber.Map{
				"success": false,
				"message": err,
			})
	}

	return c.Status(200).JSON(
		&fiber.Map{
			"success":            true,
			"today_active_users": result,
		})
}

// @Summary Add user
// @Tags debug
// @Description adding user for debug
// @ID add-user
// @Accept  json
// @Produce  json
// @Param input body models.User true "user"
// @Success 200 {object} models.ErrorResponse
// @Failure 400,500 {object} models.ErrorResponse
// @Failure default {object} models.ErrorResponse
// @Router /api/v1/add-user [post]
func (h *AppHandlers) handlerAddUser(c *fiber.Ctx) error {
	req := &models.User{}
	err := c.BodyParser(req)
	if err != nil {
		return c.Status(400).JSON(
			&fiber.Map{
				"success": false,
				"message": err,
			})
	}
	fmt.Println("user from req:", req)

	err = h.use.AddUser(req)
	if err != nil {
		return c.Status(500).JSON(
			&fiber.Map{
				"success": false,
				"message": err,
			})
	}

	return c.Status(200).JSON(
		&fiber.Map{
			"success": true,
			"message": "user created",
		})
}

// @Summary Add comment
// @Tags debug
// @Description adding comment for debug
// @ID add-comment
// @Accept  json
// @Produce  json
// @Param input body models.CommentRequest true "comment" default(models.CommentForSwagger)
// @Success 200 {object} models.ErrorResponse
// @Failure 400,500 {object} models.ErrorResponse
// @Failure default {object} models.ErrorResponse
// @Router /api/v1/add-comment [post]
func (h *AppHandlers) handlerAddComment(c *fiber.Ctx) error {
	req := &models.CommentRequest{}
	err := c.BodyParser(req)
	if err != nil {
		return c.Status(400).JSON(
			&fiber.Map{
				"success": false,
				"message": err,
			})
	}
	fmt.Println("comment from req:", req)

	err = h.use.AddComment(req)
	if err != nil {
		return c.Status(500).JSON(
			&fiber.Map{
				"success": false,
				"message": err,
			})
	}

	return c.Status(200).JSON(
		&fiber.Map{
			"success": true,
			"message": "comment created",
		})
}

// @Summary Add like
// @Tags debug
// @Description adding like for debug
// @ID add-like
// @Accept  json
// @Produce  json
// @Param input body models.Like true "like"
// @Success 200 {object} models.ErrorResponse
// @Failure 400,500 {object} models.ErrorResponse
// @Failure default {object} models.ErrorResponse
// @Router /api/v1/add-like [post]
func (h *AppHandlers) handlerAddLike(c *fiber.Ctx) error {
	req := &models.Like{}
	err := c.BodyParser(req)
	if err != nil {
		return c.Status(400).JSON(
			&fiber.Map{
				"success": false,
				"message": err,
			})
	}
	fmt.Println("like from req:", req)

	err = h.use.AddLike(req)
	if err != nil {
		return c.Status(500).JSON(
			&fiber.Map{
				"success": false,
				"message": err,
			})
	}

	return c.Status(200).JSON(
		&fiber.Map{
			"success": true,
			"message": "like created",
		})
}
