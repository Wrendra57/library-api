package domain

import "time"

type Author struct {
	Author_id  int
	Name       string
	Created_at time.Time
	Updated_at time.Time
}
