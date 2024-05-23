package pg_test

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/hendrihmwn/dating-app-api/app/domain/entity"
	"github.com/hendrihmwn/dating-app-api/app/domain/repository/pg"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	"time"
)

type MockUserPackage struct {
	db sqlmock.Sqlmock
}

func TestUserPackageRepositoryImpl_List(t *testing.T) {
	type params struct {
		userId int64
	}

	type input struct {
		ctx    context.Context
		params params
	}

	type output struct {
		result []entity.UserPackage
		err    error
	}

	type testcase struct {
		name string
		in   *input
		out  *output
		on   func(*MockUserPackage, *input, *output)
	}
	now := time.Now()
	expired := now.AddDate(0, 12, 0)
	user := []entity.UserPackage{{ID: 1, UserID: 1, PackageID: 1, PackageAct: "unlimited-swipe", ExpiredAt: &expired}}
	inputParams := params{userId: user[0].UserID}

	const QUERY = `
		SELECT 	user_packages.id,
			user_packages.user_id,
			user_packages.package_id,
			user_packages.package_act,
			user_packages.expired_at,
			user_packages.created_at,
			user_packages.updated_at,
			user_packages.deleted_at
	FROM user_packages
	WHERE user_packages.deleted_at IS NULL AND user_packages.expired_at > current_date AND user_packages.user_id = ?
	`
	tests := []testcase{
		{
			name: "OK",
			in: &input{
				ctx:    context.Background(),
				params: inputParams,
			},
			out: &output{
				result: nil,
			},
			on: func(ppr *MockUserPackage, i *input, o *output) {
				rows := sqlmock.NewRows([]string{
					"id",
					"user_id",
					"package_id",
					"package_act",
					"expired_at",
				}).AddRow(
					user[0].ID,
					user[0].UserID,
					user[0].PackageID,
					user[0].PackageAct,
					user[0].ExpiredAt,
				)
				ppr.db.ExpectQuery(regexp.QuoteMeta(QUERY)).WithArgs(i.params.userId).WillReturnRows(rows)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mockDB := mockDatabase()
			defer db.Close()
			m := &MockUserPackage{
				db: mockDB,
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			res, err := pg.NewUserPackageRepository(db).List(tt.in.ctx, tt.in.params.userId)

			if tt.out.err != nil {
				assert.NotNil(t, err)
				assert.Equal(t, tt.out.err, err)
			}

			if tt.out.result != nil {
				assert.NotNil(t, res)
				assert.Equal(t, tt.out.result, res)
			}
		})
	}
}

func TestUserPackageRepositoryImpl_Create(t *testing.T) {
	type params struct {
		userPackage *entity.UserPackage
	}

	type input struct {
		ctx    context.Context
		params params
	}

	type output struct {
		result *entity.UserPackage
		err    error
	}

	type testcase struct {
		name string
		in   *input
		out  *output
		on   func(*MockUserPackage, *input, *output)
	}
	now := time.Now()
	expired := now.AddDate(0, 12, 0)
	user := entity.UserPackage{ID: 1, UserID: 1, PackageID: 1, PackageAct: "unlimited-swipe", ExpiredAt: &expired}
	inputParams := params{userPackage: &user}

	const QUERY = `
		INSERT INTO user_packages (user_id, package_id, package_act, expired_at)
	VALUES (?, ?, ?, ?)
	RETURNING id, user_id, package_id, package_act, expired_at, created_at, updated_at, deleted_at;
	`
	tests := []testcase{
		{
			name: "OK",
			in: &input{
				ctx:    context.Background(),
				params: inputParams,
			},
			out: &output{
				result: nil,
			},
			on: func(ppr *MockUserPackage, i *input, o *output) {
				ppr.db.ExpectExec(regexp.QuoteMeta(QUERY)).WithArgs(
					i.params.userPackage.UserID,
					i.params.userPackage.PackageID,
					i.params.userPackage.PackageAct,
					i.params.userPackage.ExpiredAt,
				).WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mockDB := mockDatabase()
			defer db.Close()
			m := &MockUserPackage{
				db: mockDB,
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			res, err := pg.NewUserPackageRepository(db).Create(tt.in.ctx, tt.in.params.userPackage)

			if tt.out.err != nil {
				assert.NotNil(t, err)
				assert.Equal(t, tt.out.err, err)
			}

			if tt.out.result != nil {
				assert.NotNil(t, res)
				assert.Equal(t, tt.out.result, res)
			}
		})
	}
}

func TestUserPackageRepositoryImpl_Get(t *testing.T) {
	type params struct {
		id int64
	}

	type input struct {
		ctx    context.Context
		params params
	}

	type output struct {
		result *entity.UserPackage
		err    error
	}

	type testcase struct {
		name string
		in   *input
		out  *output
		on   func(*MockUserPackage, *input, *output)
	}
	now := time.Now()
	expired := now.AddDate(0, 12, 0)
	user := entity.UserPackage{ID: 1, UserID: 1, PackageID: 1, PackageAct: "unlimited-swipe", ExpiredAt: &expired}
	inputParams := params{id: user.UserID}

	const QUERY = `
		SELECT 	user_packages.id,
			user_packages.user_id,
			user_packages.package_id,
			user_packages.package_act,
			user_packages.expired_at,
			user_packages.created_at,
			user_packages.updated_at,
			user_packages.deleted_at
	FROM user_packages
	WHERE user_packages.user_id = $1 AND user_packages.deleted_at IS NULL AND user_packages.expired_at > current_date
	LIMIT 1;
`
	tests := []testcase{
		{
			name: "OK",
			in: &input{
				ctx:    context.Background(),
				params: inputParams,
			},
			out: &output{
				result: &user,
			},
			on: func(ppr *MockUserPackage, i *input, o *output) {
				rows := sqlmock.NewRows([]string{
					"id",
					"user_id",
					"package_id",
					"package_act",
					"expired_at",
				}).AddRow(
					user.ID,
					user.UserID,
					user.PackageID,
					user.PackageAct,
					user.ExpiredAt,
				)
				ppr.db.ExpectQuery(regexp.QuoteMeta(QUERY)).WithArgs(i.params.id).WillReturnRows(rows)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mockDB := mockDatabase()
			defer db.Close()
			m := &MockUserPackage{
				db: mockDB,
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			res, err := pg.NewUserPackageRepository(db).Get(tt.in.ctx, tt.in.params.id)

			if tt.out.err != nil {
				assert.NotNil(t, err)
				assert.Equal(t, tt.out.err, err)
			}

			if tt.out.result != nil {
				assert.NotNil(t, res)
				assert.Equal(t, tt.out.result, res)
			}
		})
	}
}
