package service

import (
	"context"

	"github.com/be/perpustakaan/model/domain"
	"github.com/be/perpustakaan/model/webrequest"
)

type PenaltiesService interface {
	PayPenalties(ctx context.Context, request webrequest.UpdatePenaltiesRequest, id int) domain.Penalties
	// UpdateUser(ctx context.Context, request webrequest.UpdateUserRequest, id int) bool
}
