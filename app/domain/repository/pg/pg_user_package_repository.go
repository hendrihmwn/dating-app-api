package pg

import (
	"context"
	"fmt"
	"github.com/hendrihmwn/dating-app-api/app/domain/entity"
	"github.com/hendrihmwn/dating-app-api/app/domain/repository"
	"github.com/jmoiron/sqlx"
)

type UserPackageRepositoryImpl struct {
	db *sqlx.DB
}

func (p *UserPackageRepositoryImpl) Create(ctx context.Context, userPackage *entity.UserPackage) (*entity.UserPackage, error) {
	rows, err := p.db.NamedQueryContext(ctx, CREATE_USER_PACKAGE, userPackage)

	if err != nil {
		return nil, err
	}

	var res *entity.UserPackage
	for rows.Next() {
		res = new(entity.UserPackage)
		if err := rows.StructScan(res); err != nil {
			return nil, err
		}

	}
	return res, nil
}

func (p *UserPackageRepositoryImpl) List(ctx context.Context, userId int64) ([]entity.UserPackage, error) {

	rows, err := p.db.QueryxContext(ctx, LIST_USER_PACKAGE, userId)
	if err != nil {
		return nil, err
	}

	userPackages := []entity.UserPackage{}

	for rows.Next() {
		userPackage := entity.UserPackage{}

		if err := rows.StructScan(&userPackage); err != nil {
			return nil, err
		}

		userPackages = append(userPackages, userPackage)
	}

	return userPackages, nil
}

func (p *UserPackageRepositoryImpl) Get(ctx context.Context, id int64) (*entity.UserPackage, error) {
	rows, err := p.db.QueryxContext(ctx, GET_USER_PACKAGE, id)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var res *entity.UserPackage
	for rows.Next() {
		res = new(entity.UserPackage)
		if err := rows.StructScan(res); err != nil {
			return nil, err
		}

	}
	return res, nil
}

func NewUserPackageRepository(db *sqlx.DB) repository.UserPackageRepository {
	return &UserPackageRepositoryImpl{
		db: db,
	}
}
