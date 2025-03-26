package sqlite

import (
	"database/sql"

	payback "github.com/snuts1/daren2test"
)

type paybackService struct {
	db *sql.DB
}

func NewPaybackService(db *sql.DB) payback.PaybackService {
	return &paybackService{
		db: db,
	}
}

//methods
