package control

import (
	"github.com/gofiber/fiber/v2"
)

type Handlers interface {
	DebitingMoney() fiber.Handler
	CreditingMoney() fiber.Handler
	Transfer() fiber.Handler
	GetBalance() fiber.Handler
}