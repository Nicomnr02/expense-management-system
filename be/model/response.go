package model

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
	Pagination
}

func Success(ctx *fiber.Ctx, code int, data interface{}) error {
	return ctx.Status(code).JSON(response{
		Code:    code,
		Success: true,
		Message: "success",
		Data:    data,
	})
}

func SuccessPage(ctx *fiber.Ctx, data interface{}, pagination Pagination) error {
	return ctx.Status(http.StatusOK).JSON(responsePage{
		response: response{
			Code:    http.StatusOK,
			Success: true,
			Message: "success",
			Data:    data,
		},
		Pagination: Pagination{
			Page:  pagination.Page,
			Limit: pagination.Limit,
			Total: pagination.Total,
		},
	})
}

func Error(ctx *fiber.Ctx, err error, data interface{}) error {
	fberror := err.(*fiber.Error)
	return ctx.Status(fberror.Code).JSON(response{
		Code:    fberror.Code,
		Success: false,
		Message: fberror.Message,
		Data:    data,
	})
}
