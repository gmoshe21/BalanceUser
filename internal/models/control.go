package models

type UpdateAccounts struct {
	Id    int `json:"ID" validate:"required"`
	Money float64 `json:"Money" validate:"required"`
}

type Transfer struct {
	Sender    int `json:"Sender" validate:"required"`
	Recipient int `json:"Recipient" validate:"required"`
	Money     float64 `json:"Money" validate:"required"`

}