package repository

import (
	"BalanceUser/internal/control"
	"BalanceUser/internal/models"
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type controlRepo struct {
	db     *sqlx.DB
}

func NewControlRepository(db *sqlx.DB) control.Repository {
	return &controlRepo{db: db}
}

func (c *controlRepo) DebitingMoney(ctx context.Context, params models.UpdateAccounts) error {
	_, err := c.db.ExecContext(ctx, queryDebitingMoney, params.Id, params.Money)
	if err != nil {
		errors.Wrapf(err, "controlRepo.DebitingMoney.ExecContext(Sender : %s, Recipient: %s)",
			params.Id,
			params.Money,
		)
	}
	return nil
}

func (c *controlRepo) CreditingMoney(ctx context.Context, params models.UpdateAccounts) error {
	_, err := c.db.ExecContext(ctx, queryCreditingMoney, params.Id, params.Money)
	if err != nil {
		errors.Wrapf(err, "controlRepo.CreditingMoney.ExecContext(Sender : %s, Recipient: %s)",
			params.Id,
			params.Money,
		)
	}
	return nil
}

func (c *controlRepo) GetBalance(ctx context.Context, id int) (result []byte, err error) {
	err = c.db.GetContext(ctx, &result, queryGetBalance, id)
	if err != nil {
		return nil, errors.Wrap(err, "controlRepo.GetBalance.GetContext()")
	}
	return result, nil
}

