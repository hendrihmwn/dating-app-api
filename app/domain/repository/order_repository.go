package repository

import (
	"context"
	"github.com/hendrihmwn/dating-app-api/app/domain/entity"
)

type OrderRepository interface {
	Create(ctx context.Context, order *entity.Order) (*entity.Order, error)
}
