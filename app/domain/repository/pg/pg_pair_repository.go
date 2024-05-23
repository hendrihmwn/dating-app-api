package pg

import (
	"context"
	"github.com/hendrihmwn/dating-app-api/app/domain/entity"
	"github.com/hendrihmwn/dating-app-api/app/domain/repository"
	"github.com/jmoiron/sqlx"
)

type PairRepositoryImpl struct {
	db *sqlx.DB
}

func (p *PairRepositoryImpl) Create(ctx context.Context, userPair *entity.UserPair) (*entity.UserPair, error) {
	rows, err := p.db.NamedQueryContext(ctx, CREATE_USER_PAIR, userPair)

	if err != nil {
		return nil, err
	}

	var res *entity.UserPair
	for rows.Next() {
		res = new(entity.UserPair)
		if err := rows.StructScan(res); err != nil {
			return nil, err
		}
		if res.Status == 0 {
			res.StatusLang = "pass"
		} else {
			res.StatusLang = "like"
		}
	}
	return res, nil
}

func (p *PairRepositoryImpl) CountToday(ctx context.Context, userId int64) (int, error) {

	rows, err := p.db.QueryxContext(ctx, COUNT_USER_PAIR_TODAY, userId)
	if err != nil {
		return 0, err
	}

	var res int

	for rows.Next() {
		err = rows.Scan(&res)
		if err != nil {
			return 0, err
		}
	}

	return res, nil
}

func NewUserPairRepository(db *sqlx.DB) repository.UserPairRepository {
	return &PairRepositoryImpl{
		db: db,
	}
}
