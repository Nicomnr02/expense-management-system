package dto

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type response struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type responsePage struct {
	response
	Page
}

func Success(ctx *fiber.Ctx, code int, data interface{}) error {
	return ctx.Status(code).JSON(response{
		Code:    code,
		Success: true,
		Message: "success",
		Data:    data,
	})
}

func SuccessPage(ctx *fiber.Ctx, data interface{}, page Page) error {
	return ctx.Status(http.StatusOK).JSON(responsePage{
		response: response{
			Code:    http.StatusOK,
			Success: true,
			Message: "success",
			Data:    data,
		},
		Page: Page{
			Page:  page.Page,
			Limit: page.Limit,
			Total: page.Total,
		},
	})
}

func Error(ctx *fiber.Ctx, err *fiber.Error, data interface{}) error {
	return ctx.Status(err.Code).JSON(response{
		Code:    err.Code,
		Success: false,
		Message: err.Message,
		Data:    data,
	})
}
