package control

import (
	"BalanceUser/internal/models"
	"context"
)

type Repository interface {
	DebitingMoney(ctx context.Context, params models.UpdateAccounts) error
	CreditingMoney(ctx context.Context, params models.UpdateAccounts) error
	GetBalance(ctx context.Context, id int) (result []byte, err error)
}