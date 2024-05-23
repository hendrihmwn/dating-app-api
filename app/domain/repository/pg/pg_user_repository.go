package pg

import (
	"context"
	"fmt"
	"github.com/hendrihmwn/dating-app-api/app/domain/entity"
	"github.com/hendrihmwn/dating-app-api/app/domain/repository"
	"github.com/jmoiron/sqlx"
	"strings"
)

type UserRepositoryImpl struct {
	db *sqlx.DB
}

func (p *UserRepositoryImpl) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	// Initiate the transaction
	tx := p.db.MustBegin()

	row, err := tx.NamedQuery(CREATE_USER, user)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, err
		}

		return nil, err
	}

	// Scan insert account_users result
	for row.Next() {
		if err := row.StructScan(&user); err != nil {
			// Abort transaction
			if err := tx.Rollback(); err != nil {
				return nil, err
			}

			return nil, err
		}
	}

	userProfile := entity.UserProfile{
		UserID: user.ID,
	}
	_, err = tx.NamedQuery(CREATE_PROFILE, userProfile)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, err
		}

		return nil, err
	}

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		// Abort transaction
		if err := tx.Rollback(); err != nil {
			return nil, err
		}

		return nil, err
	}

	return user, nil
}

func (p *UserRepositoryImpl) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user *entity.User

	rows, err := p.db.QueryxContext(ctx, GET_USER_BY_EMAIL, email)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user = new(entity.User)
		err := rows.StructScan(&user)
		if err != nil {
			return nil, err
		}
	}

	if user == nil {
		return nil, nil
	}

	return user, nil
}

func (p *UserRepositoryImpl) List(ctx context.Context, args *repository.ListUsersArgs) ([]entity.User, error) {
	var (
		query = LIST_USER
		qargs []interface{}
	)

	if args.Filter != nil {
		filterQuery, filterArgs, err := p.Filter(ctx, args.Filter)
		if err != nil {
			return nil, err
		}

		query = fmt.Sprint(query, "AND ", filterQuery)
		qargs = append(qargs, filterArgs...)
	}

	paginationQuery, paginationArgs := Paginate(ctx, args.Limit, args.Offset)
	query = fmt.Sprint(query, " ", paginationQuery)
	qargs = append(qargs, paginationArgs...)

	query = p.db.Rebind(query)

	stmt, err := p.db.PreparexContext(ctx, query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryxContext(ctx, qargs...)
	if err != nil {
		return nil, err
	}

	users := []entity.User{}

	for rows.Next() {
		user := entity.User{}

		if err := rows.StructScan(&user); err != nil {
			return nil, err
		}
		if user.Gender == 0 {
			user.GenderLang = "men"
		} else {
			user.GenderLang = "women"
		}
		users = append(users, user)
	}

	return users, nil
}

func (p *UserRepositoryImpl) Filter(ctx context.Context, filter *repository.UserFilterArgs) (query string, args []interface{}, err error) {
	queries := []string{}

	if len(filter.ExcludeUserId) > 0 {
		additionalFilterUser := "(users.id not in (SELECT user_pairs.pair_user_id FROM user_pairs WHERE (user_pairs.status = 1 OR (user_pairs.status = 0 AND user_pairs.created_at >= current_date AND user_pairs.created_at < current_date + interval '1 day')) AND user_pairs.user_id IN (:user_ids)))"
		queries = append(queries, fmt.Sprintf("(users.id NOT IN (:user_ids)) AND %s", additionalFilterUser))
	}

	query = strings.Join(queries, " AND ")
	query, args, err = sqlx.Named(query, filter)
	if err != nil {
		return "", nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return "", nil, err
	}

	return query, args, nil
}

func Paginate(ctx context.Context, limit int, offset int) (query string, args []interface{}) {
	query = "LIMIT ? OFFSET ?"
	args = []interface{}{limit, offset}
	return query, args
}

func NewUserRepository(db *sqlx.DB) repository.UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}
