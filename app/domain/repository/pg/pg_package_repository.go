package pg

import (
	"context"
	"fmt"
	"github.com/hendrihmwn/dating-app-api/app/domain/entity"
	"github.com/hendrihmwn/dating-app-api/app/domain/repository"
	"github.com/jmoiron/sqlx"
)

type PackageRepositoryImpl struct {
	db *sqlx.DB
}

func (p *PackageRepositoryImpl) Get(ctx context.Context, id int64) (*entity.Package, error) {
	rows, err := p.db.QueryxContext(ctx, GET_PACKAGE, id)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var res *entity.Package
	for rows.Next() {
		res = new(entity.Package)
		if err := rows.StructScan(res); err != nil {
			return nil, err
		}

	}
	return res, nil
}

func (p *PackageRepositoryImpl) List(ctx context.Context) ([]entity.Package, error) {

	rows, err := p.db.QueryxContext(ctx, LIST_PACKAGE)
	if err != nil {
		return nil, err
	}

	packages := []entity.Package{}

	for rows.Next() {
		pkg := entity.Package{}

		if err := rows.StructScan(&pkg); err != nil {
			return nil, err
		}

		packages = append(packages, pkg)
	}

	return packages, nil
}

func NewPackageRepository(db *sqlx.DB) repository.PackageRepository {
	return &PackageRepositoryImpl{
		db: db,
	}
}
