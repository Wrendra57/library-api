package domain

import "time"

type Penalties struct {
	Penalty_id     int
	Loan_id        int
	Penalty_amount int
	Reason         string
	Payment_status string
	Due_date       time.Time
	Admin_id       int
	Created_at     time.Time
	Updated_at     time.Time
}
