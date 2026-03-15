package domain

import (
	"database/sql"
	"time"
)

type (
	Account struct {
		Name        string       `json:"name"`
		Email       string       `json:"email"`
		Password    string       `json:"password"`
		ActivatedAt sql.NullTime `json:"activatedAt"`
		CreatedAt   time.Time    `json:"createdAt"`
		UpdatedAt   time.Time    `json:"updatedAt"`
	}
)
