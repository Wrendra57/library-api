package webresponse

import "time"

type PenaltiesResponse struct {
	Penalty_id     int       `json:"penalty_id"`
	Loan_id        int       `json:"loan_id"`
	Penalty_amount int       `json:"penalty_amount"`
	Reason         string    `json:"reason"`
	Payment_status string    `json:"payment_status"`
	Due_date       time.Time `json:"due_date"`
	Admin_id       int       `json:"admin_id"`
	Created_at     time.Time `json:"created_at"`
	Updated_at     time.Time `json:"updated_at"`
}
