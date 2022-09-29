package http

import(
	"BalanceUser/internal/control"

	"github.com/gofiber/fiber/v2"
)

func MapAPIRoutes(group fiber.Router, h control.Handlers) {
	group.Post("/debiting_money", h.DebitingMoney())
	group.Post("/crediting_money", h.CreditingMoney())
	group.Post("/transfer", h.Transfer())
	group.Get("/get_balance", h.GetBalance())
}