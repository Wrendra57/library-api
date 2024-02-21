package webrequest

import (
	"database/sql"
)

type UpdatePenaltiesRequest struct {
	Payment_status string       `validate:"omitempty,max=255" json:"payment_status"`
	Admin_id       int          `validate:"omitempty,min=0" json:"admin_id"`
	Reason         string       `validate:"omitempty,min=3" json:"reason"`
	Due_date       sql.NullTime ` json:"due_date" `
	Penalty_amount int          `validate:"omitempty,min=0" json:"penalty_amount"`
}
