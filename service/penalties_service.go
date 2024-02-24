package service

import (
	"context"

	"github.com/be/perpustakaan/model/webrequest"
	"github.com/be/perpustakaan/model/webresponse"
)

type PenaltiesService interface {
	PayPenalties(ctx context.Context, request webrequest.UpdatePenaltiesRequest, id int) webresponse.PenaltiesResponse
	// UpdateUser(ctx context.Context, request webrequest.UpdateUserRequest, id int) bool
}
