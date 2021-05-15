package http

import (
	"fmt"

	"github.com/amartery/trinity-task/pkg/app"
	"github.com/amartery/trinity-task/pkg/app/models"
	"github.com/gofiber/fiber/v2"
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
	router.Get("/api/v1/get-active-today", h.handlerActive)
	router.Post("/api/v1/add-user", h.handlerAddUser)
	router.Post("/api/v1/add-comment", h.handlerAddComment)
	router.Post("/api/v1/add-like", h.handlerAddLike)
}

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

func (h *AppHandlers) handlerAddLike(c *fiber.Ctx) error {
	req := &models.Lile{}
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
