package domain

import "time"

type Rak struct {
	Rak_id     int
	Name       string
	Rows_rak   int
	Col        int
	Created_at time.Time
	Updated_at time.Time
}
