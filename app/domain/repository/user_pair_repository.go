package repository

import (
	"context"
	"github.com/hendrihmwn/dating-app-api/app/domain/entity"
)

type UserPairRepository interface {
	Create(ctx context.Context, userPair *entity.UserPair) (*entity.UserPair, error)
	CountToday(ctx context.Context, userId int64) (int, error)
}
