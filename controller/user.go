package controller

import (
	"evolve/database"
	"evolve/model"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type JsonResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GetUsers(ctx *fiber.Ctx) error {
	var users []model.User

	page, _ := strconv.Atoi(ctx.Query("page"))
	size, _ := strconv.Atoi(ctx.Query("size"))

	from := ctx.Query("from")
	to := ctx.Query("to")

	if page == 0 {
		page = 1
	}
	if size == 0 {
		size = 10
	}

	p := model.Pagination{
		Page: page,
		Size: size,
	}
	chain := database.DB.Scopes(model.Paginate(p)).Order("id ASC")
	if from != "" {
		if to != "" {
			chain = chain.Where("created_at > ? AND created_at <=  ?", from, to)
		}
	}

	if err := chain.Find(&users).Error; err != nil {
		return ctx.JSON(JsonResponse{
			Success: false,
			Message: "No users found",
			Data:    err,
		})
	}

	return ctx.JSON(fiber.Map{
		"page":         p.Page,
		"result_count": p.Size,
		"success":      true,
		"message":      "Users retrieved successfully",
		"data":         users,
	})
}
func GetUserByEmail(ctx *fiber.Ctx) error {
	email := ctx.Params("email")
	var user model.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return ctx.JSON(JsonResponse{
			Success: false,
			Message: "No user found with email",
			Data:    err,
		})
	}
	return ctx.JSON(JsonResponse{
		Success: true,
		Message: "User retrieved successfully",
		Data:    user,
	})
}
