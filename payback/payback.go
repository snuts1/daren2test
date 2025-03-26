package payback

import (
	"time"
)

type Gathering struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	NumGatherers int       `json:"numGatherers"`
	Names        string    `json:"text"`
	CreatedAt    time.Time `json:"createdAt"`
}

type Payment struct {
	ID             int       `json:"id"`
	GatheringID    int       `json:"GatheringID"`
	Desc           string    `json:"desc"`
	Amount         float32   `json:"amount"`
	CreditorID     int       `json:"creditorID`
	DebtorIDString string    `json:"debtorIDString"`
	CreatedAt      time.Time `json:"createdAt"`
}

type PaybackService interface {
	CreateGathering(g *Gathering) (int, error)
	GetAllGatherings() ([]*Gathering, error)
	CreatePayment(p *Payment) (int, error)
	GetAllPayments() ([]*Payment, error)
	DeletePayment(id int) error
}
