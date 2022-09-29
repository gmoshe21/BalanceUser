package http

import (
	"BalanceUser/config"
	"BalanceUser/internal/control"
	"BalanceUser/internal/models"
	"BalanceUser/pkg/reqvalidator"
	"context"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

type controlHandlers struct {
	cfg       *config.Config
	controlUC control.UseCase
}

func NewControlHandlers(cfg *config.Config, controlUC control.UseCase) control.Handlers {
	return &controlHandlers{cfg: cfg, controlUC: controlUC}
}

func (ctrl *controlHandlers) DebitingMoney() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params models.UpdateAccounts

		if err := reqvalidator.ReadRequest(c, &params); err != nil {
			log.Println("controlHandlers.DebitingMoney.ReadRequest", err)
			return err
		}

		err := ctrl.controlUC.DebitingMoney(context.Background(), params)
		if err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

func (ctrl *controlHandlers) CreditingMoney() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params models.UpdateAccounts

		if err := reqvalidator.ReadRequest(c, &params); err != nil {
			log.Println("controlHandlers.CreditingMoney.ReadRequest", err)
			return err
		}

		err := ctrl.controlUC.CreditingMoney(context.Background(), params)
		if err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

func (ctrl *controlHandlers) Transfer() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params models.Transfer

		if err := reqvalidator.ReadRequest(c, &params); err != nil {
			log.Println("controlHandlers.Transfer.ReadRequest", err)
			return err
		}

		err := ctrl.controlUC.Transfer(context.Background(), params)
		if err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

func (ctrl *controlHandlers) GetBalance() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(string(c.Context().QueryArgs().Peek("Id")))
		if err != nil {
			return errors.Wrap(err, "controlHandlers.GetBalance.Atoi")
		}

		result, err := ctrl.controlUC.GetBalance(context.Background(), id)
		if err != nil {
			return err
		}

		return c.Send(result)
	}
}