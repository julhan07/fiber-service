package utils

import "github.com/gofiber/fiber/v2"

type payload struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseSuccess(c *fiber.Ctx, code *int, msg *string, data *interface{}) {
	c.JSONP(&payload{
		Code:    *code,
		Message: *msg,
		Data:    *data,
	})
}

func ResponseError(c *fiber.Ctx, code int, msg string) error {
	return c.JSONP(&payload{
		Code:    code,
		Message: msg,
	})
}
