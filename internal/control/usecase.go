package control

import (
	"BalanceUser/internal/models"
	"context"
)

type UseCase interface {
	DebitingMoney(ctx context.Context, params models.UpdateAccounts) error
	CreditingMoney(ctx context.Context, params models.UpdateAccounts) error
	Transfer(ctx context.Context, params models.Transfer) error
	GetBalance(ctx context.Context, id int) (result []byte, err error)
}