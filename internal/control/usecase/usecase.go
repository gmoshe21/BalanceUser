package usecase

import (
	"BalanceUser/config"
	"BalanceUser/internal/control"
	"BalanceUser/internal/models"

	"encoding/binary"
    "math"
	"context"
	"errors"
)

type controlUC struct {
	cfg                    *config.Config
	controlRepo            control.Repository
}

func NewControlUseCase( cfg *config.Config, controlRepo control.Repository) control.UseCase {
	return &controlUC{cfg: cfg, controlRepo: controlRepo}
}

func (c *controlUC) DebitingMoney(ctx context.Context, params models.UpdateAccounts) error {
	check, err := c.controlRepo.GetBalance(ctx, params.Id)
	if err != nil {
		return err
	}

	if math.Float64frombits(binary.LittleEndian.Uint64(check)) < params.Money {
		return errors.New("Insufficient funds to complete the transaction")
	}

	return c.controlRepo.DebitingMoney(ctx, params)
}

func (c *controlUC) CreditingMoney(ctx context.Context, params models.UpdateAccounts) error {
	return c.controlRepo.CreditingMoney(ctx, params)
}

func (c *controlUC) Transfer(ctx context.Context, params models.Transfer) error {
	data := models.UpdateAccounts {
		Id: params.Sender,
		Money: params.Money,
	}

	err := c.controlRepo.DebitingMoney(ctx, data)
	if err != nil {
		return err
	}

	data.Id = params.Recipient
	return c.controlRepo.CreditingMoney(ctx, data)
}

func (c *controlUC) GetBalance(ctx context.Context, id int) (result []byte, err error) {
	return c.controlRepo.GetBalance(ctx, id)
}
