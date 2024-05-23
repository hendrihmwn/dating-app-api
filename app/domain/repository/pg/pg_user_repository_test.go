package pg_test

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/hendrihmwn/dating-app-api/app/domain/entity"
	"github.com/hendrihmwn/dating-app-api/app/domain/repository"
	"github.com/hendrihmwn/dating-app-api/app/domain/repository/pg"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"log"
	"regexp"
	"testing"
)

type MockUserRepository struct {
	db sqlmock.Sqlmock
}

func mockDatabase() (*sqlx.DB, sqlmock.Sqlmock) {
	db, dbMock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	return sqlxDB, dbMock
}

func TestUserRepositoryImpl_List(t *testing.T) {
	type params struct {
		args *repository.ListUsersArgs
	}

	type input struct {
		ctx    context.Context
		params params
	}

	type output struct {
		result []entity.User
		err    error
	}

	type testcase struct {
		name string
		in   *input
		out  *output
		on   func(*MockUserRepository, *input, *output)
	}

	user := []entity.User{{ID: 1, Name: "Testt", GenderLang: "men"}}
	inputParams := params{args: &repository.ListUsersArgs{}}

	const QUERY = `
		SELECT DISTINCT users.id,
			users.name,
			users.email,
			user_profiles.image,
			user_profiles.birthdate,
			user_profiles.location,
			user_profiles.gender,
			users.created_at,
			users.updated_at,
			users.deleted_at,
			user_packages.package_act as label
	FROM users
	JOIN user_profiles ON user_profiles.user_id = users.id
	LEFT JOIN user_packages ON user_packages.user_id = users.id AND user_packages.package_act = 'verified-label' AND user_packages.expired_at > current_date
	WHERE users.deleted_at IS NULL
	`
	tests := []testcase{
		{
			name: "INTERNAL_SERVER_ERROR",
			in: &input{
				ctx:    context.Background(),
				params: inputParams,
			},
			out: &output{
				err: errors.New(""),
			},
			on: func(ppr *MockUserRepository, i *input, o *output) {
				query := QUERY
				i.params.args.Filter = &repository.UserFilterArgs{}
				ppr.db.ExpectPrepare(regexp.QuoteMeta(query)).WillReturnError(errors.New(""))
			},
		},
		{
			name: "OK",
			in: &input{
				ctx:    context.Background(),
				params: inputParams,
			},
			out: &output{
				result: user,
			},
			on: func(ppr *MockUserRepository, i *input, o *output) {
				ppr.db.ExpectPrepare(regexp.QuoteMeta(QUERY)).WillBeClosed()
				rows := sqlmock.NewRows([]string{
					"id",
					"name",
				}).AddRow(
					user[0].ID,
					user[0].Name,
				)
				ppr.db.ExpectQuery(regexp.QuoteMeta(QUERY)).WillReturnRows(rows)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mockDB := mockDatabase()
			defer db.Close()
			m := &MockUserRepository{
				db: mockDB,
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			res, err := pg.NewUserRepository(db).List(tt.in.ctx, tt.in.params.args)

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

func TestUserRepositoryImpl_Create(t *testing.T) {
	type params struct {
		userEntity *entity.User
	}

	type input struct {
		ctx    context.Context
		params params
	}

	type output struct {
		result *entity.User
		err    error
	}

	type testcase struct {
		name string
		in   *input
		out  *output
		on   func(*MockUserRepository, *input, *output)
	}

	//user := entity.User{ID: 1, Email: "test@yopmail.com", Name: "Testt"}
	inputParams := params{userEntity: &entity.User{Name: "Testt", Password: "123456", Email: "test@yopmail.com"}}

	const QUERY = `
		INSERT INTO users (name, email, password)
	VALUES (?, ?, ?)
	RETURNING id, name, email, password, created_at, updated_at, deleted_at;
	`
	const CREATE_PROFILE = `
	INSERT INTO user_profiles (user_id, image, birthdate, gender, location)
	VALUES (?, ?, ?, ?, ?)
	RETURNING id, user_id, image, birthdate, gender, location, created_at, updated_at, deleted_at;
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
			on: func(ppr *MockUserRepository, i *input, o *output) {
				query := QUERY
				i.params.userEntity = &entity.User{}
				ppr.db.ExpectBegin()
				ppr.db.ExpectExec(regexp.QuoteMeta(query)).WithArgs(
					i.params.userEntity.Name,
					i.params.userEntity.Email,
					i.params.userEntity.Password,
				).WillReturnResult(sqlmock.NewResult(1, 1))
				ppr.db.ExpectExec(regexp.QuoteMeta(CREATE_PROFILE)).WithArgs(
					1,
				).WillReturnResult(sqlmock.NewResult(1, 1))
				ppr.db.ExpectCommit()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mockDB := mockDatabase()
			defer db.Close()
			m := &MockUserRepository{
				db: mockDB,
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			res, err := pg.NewUserRepository(db).Create(tt.in.ctx, tt.in.params.userEntity)

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

func TestUserRepositoryImpl_GetByEmail(t *testing.T) {
	type params struct {
		email string
	}

	type input struct {
		ctx    context.Context
		params params
	}

	type output struct {
		result *entity.User
		err    error
	}

	type testcase struct {
		name string
		in   *input
		out  *output
		on   func(*MockUserRepository, *input, *output)
	}

	user := entity.User{ID: 1, Email: "test@yopmail.com", Name: "Testt", Password: "123456"}
	inputParams := params{email: user.Email}

	const QUERY = `
		SELECT 	users.id,
			users.name,
			users.email,
			users.password,
			users.created_at,
			users.updated_at,
			users.deleted_at
	FROM users
	WHERE users.email = $1
	AND users.deleted_at IS NULL
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
				result: nil,
			},
			on: func(ppr *MockUserRepository, i *input, o *output) {
				rows := sqlmock.NewRows([]string{"id", "email", "password"}).AddRow(user.ID, user.Email, user.Password)
				ppr.db.ExpectQuery(regexp.QuoteMeta(QUERY)).WithArgs(i.params.email).WillReturnRows(rows)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mockDB := mockDatabase()
			defer db.Close()
			m := &MockUserRepository{
				db: mockDB,
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			res, err := pg.NewUserRepository(db).GetByEmail(tt.in.ctx, tt.in.params.email)

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
