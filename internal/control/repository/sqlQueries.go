package repository

const (
	queryCreditingMoney = `UPDATE balance_user SET balanse = balance + $2 WHERE id = $1;`
	queryDebitingMoney = `UPDATE balance_user SET balanse = balance - $2 WHERE id = $1;`
	queryGetBalance = `SELECT balanse FROM balance_user WHERE id = $1`
)