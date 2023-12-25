package domain

import "time"

type Publisher struct {
	Publisher_id int
	Name         string
	Created_at   time.Time
	Updated_at   time.Time
}
