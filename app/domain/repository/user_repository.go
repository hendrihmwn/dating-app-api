package repository

import (
	"context"
	"github.com/hendrihmwn/dating-app-api/app/domain/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	List(ctx context.Context, args *ListUsersArgs) ([]entity.User, error)
}

type ListUsersArgs ListArgs[UserFilterArgs]

type UserFilterArgs struct {
	ExcludeUserId []int64 `db:"user_ids"`
}
