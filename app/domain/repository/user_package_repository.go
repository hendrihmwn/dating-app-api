package repository

import (
	"context"
	"github.com/hendrihmwn/dating-app-api/app/domain/entity"
)

type UserPackageRepository interface {
	Create(ctx context.Context, userPackage *entity.UserPackage) (*entity.UserPackage, error)
	List(ctx context.Context, userId int64) ([]entity.UserPackage, error)
	Get(ctx context.Context, userId int64) (*entity.UserPackage, error)
}
