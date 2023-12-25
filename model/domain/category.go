package domain

import "time"

type Category struct {
	Category_id int
	Category    string
	Created_at  time.Time
	Updated_at  time.Time
}
