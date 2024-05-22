package repository

import (
	"context"
	"github.com/hendrihmwn/dating-app-api/app/domain/entity"
)

type PackageRepository interface {
	Get(ctx context.Context, id int64) (*entity.Package, error)
	List(ctx context.Context) ([]entity.Package, error)
}
