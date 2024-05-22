package pg

import (
	"context"
	"github.com/hendrihmwn/dating-app-api/app/domain/entity"
	"github.com/hendrihmwn/dating-app-api/app/domain/repository"
	"github.com/jmoiron/sqlx"
)

type OrderRepositoryImpl struct {
	db *sqlx.DB
}

func (p *OrderRepositoryImpl) Create(ctx context.Context, order *entity.Order) (*entity.Order, error) {
	rows, err := p.db.NamedQueryContext(ctx, CREATE_ORDER, order)

	if err != nil {
		return nil, err
	}

	var res *entity.Order
	for rows.Next() {
		res = new(entity.Order)
		if err := rows.StructScan(res); err != nil {
			return nil, err
		}

	}
	return res, nil
}

func NewOrderRepository(db *sqlx.DB) repository.OrderRepository {
	return &OrderRepositoryImpl{
		db: db,
	}
}
