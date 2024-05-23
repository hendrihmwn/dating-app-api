package service_test

import (
	"context"
	"errors"
	"github.com/hendrihmwn/dating-app-api/app/domain/entity"
	"github.com/hendrihmwn/dating-app-api/app/domain/service"
	mocks "github.com/hendrihmwn/dating-app-api/mocks/domain/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

type MockUserService struct {
	userRepository        *mocks.UserRepository
	userPackageRepository *mocks.UserPackageRepository
	userPairRepository    *mocks.UserPairRepository
}

func TestUserServiceImpl_List(t *testing.T) {
	type input struct {
		ctx    context.Context
		params *service.UserListParams
	}

	type output struct {
		result *service.UserListResponse
		err    error
	}

	type testcase struct {
		name string
		in   *input
		out  *output
		on   func(*MockUserService, *input, *output)
	}

	userListParam := &service.UserListParams{
		Limit:  5,
		UserId: 1,
	}

	for _, tt := range []testcase{
		{
			name: "INTERNAL_SERVER_ERROR",
			in: &input{
				ctx:    context.Background(),
				params: userListParam,
			},
			out: &output{
				result: nil,
				err:    errors.New(""),
			},
			on: func(m *MockUserService, i *input, o *output) {
				m.userPackageRepository.On("List", i.ctx, int64(1)).Return(nil, o.err)
			},
		},
		{
			name: "INTERNAL_SERVER_ERROR",
			in: &input{
				ctx:    context.Background(),
				params: userListParam,
			},
			out: &output{
				result: nil,
				err:    errors.New(""),
			},
			on: func(m *MockUserService, i *input, o *output) {
				m.userPackageRepository.On("List", i.ctx, int64(1)).Return([]entity.UserPackage{{
					UserID:     1,
					PackageAct: "unlimited-swipe",
				}}, nil)
				m.userPairRepository.On("CountToday", i.ctx, int64(1)).Return(0, o.err)
			},
		},
		{
			name: "MAX_LIMIT",
			in: &input{
				ctx:    context.Background(),
				params: userListParam,
			},
			out: &output{
				result: nil,
				err:    errors.New("max limit 10 has reach"),
			},
			on: func(m *MockUserService, i *input, o *output) {
				m.userPackageRepository.On("List", i.ctx, int64(1)).Return([]entity.UserPackage{{}}, nil)
				m.userPairRepository.On("CountToday", i.ctx, int64(1)).Return(11, nil)
			},
		},
		{
			name: "INTERNAL_SERVER_ERROR",
			in: &input{
				ctx:    context.Background(),
				params: userListParam,
			},
			out: &output{
				result: nil,
				err:    errors.New(""),
			},
			on: func(m *MockUserService, i *input, o *output) {
				m.userPackageRepository.On("List", i.ctx, int64(1)).Return([]entity.UserPackage{{}}, nil)
				m.userPairRepository.On("CountToday", i.ctx, int64(1)).Return(5, nil)
				m.userRepository.On("List", i.ctx, mock.AnythingOfType("*repository.ListUsersArgs")).Return(nil, o.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx:    context.Background(),
				params: userListParam,
			},
			out: &output{
				result: &service.UserListResponse{
					Data: []entity.User{
						{
							ID:    2,
							Email: "testmail@yopmail.com",
						},
					},
				},
				err: nil,
			},
			on: func(m *MockUserService, i *input, o *output) {
				m.userPackageRepository.On("List", i.ctx, int64(1)).Return([]entity.UserPackage{{}}, nil)
				m.userPairRepository.On("CountToday", i.ctx, int64(1)).Return(5, nil)
				m.userRepository.On("List", i.ctx, mock.AnythingOfType("*repository.ListUsersArgs")).Return([]entity.User{
					{
						ID:    2,
						Email: "testmail@yopmail.com",
					},
				}, nil)
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockUserService{
				userRepository:        &mocks.UserRepository{},
				userPackageRepository: &mocks.UserPackageRepository{},
				userPairRepository:    &mocks.UserPairRepository{},
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := service.NewUserService(
				m.userRepository, m.userPairRepository, m.userPackageRepository,
			)
			result, err := subject.List(tt.in.ctx, *tt.in.params)

			if tt.out.err != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.out.err, err)
			}

			if tt.out.result != nil {
				assert.NotNil(t, result)
				assert.Equal(t, tt.out.result, result)
			}
		})
	}
}

func TestUserServiceImpl_Login(t *testing.T) {
	type input struct {
		ctx    context.Context
		params *service.UserLoginParams
	}

	type output struct {
		result *service.UserLoginResponse
		err    error
	}

	type testcase struct {
		name string
		in   *input
		out  *output
		on   func(*MockUserService, *input, *output)
	}

	userLoginParams := &service.UserLoginParams{
		Email:    "testmail@yopmail.com",
		Password: "123456",
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(userLoginParams.Password), bcrypt.DefaultCost)
	userLoginRes, _ := service.CreateSession(&entity.User{ID: 1, Name: "Test", Password: string(hash)})

	for _, tt := range []testcase{
		{
			name: "INTERNAL_SERVER_ERROR",
			in: &input{
				ctx:    context.Background(),
				params: userLoginParams,
			},
			out: &output{
				result: nil,
				err:    errors.New(""),
			},
			on: func(m *MockUserService, i *input, o *output) {
				m.userRepository.On("GetByEmail", i.ctx, userLoginParams.Email).Return(nil, o.err)
			},
		},
		{
			name: "EMAIL_NOT_FOUND",
			in: &input{
				ctx:    context.Background(),
				params: userLoginParams,
			},
			out: &output{
				result: nil,
				err:    errors.New("Email not found"),
			},
			on: func(m *MockUserService, i *input, o *output) {
				m.userRepository.On("GetByEmail", i.ctx, userLoginParams.Email).Return(nil, nil)
			},
		},
		{
			name: "WRONG_PASSWORD",
			in: &input{
				ctx:    context.Background(),
				params: userLoginParams,
			},
			out: &output{
				result: nil,
				err:    errors.New("Wrong password"),
			},
			on: func(m *MockUserService, i *input, o *output) {
				m.userRepository.On("GetByEmail", i.ctx, userLoginParams.Email).Return(&entity.User{ID: 1, Name: "Test", Password: "sss"}, nil)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx:    context.Background(),
				params: userLoginParams,
			},
			out: &output{
				result: userLoginRes,
				err:    nil,
			},
			on: func(m *MockUserService, i *input, o *output) {
				m.userRepository.On("GetByEmail", i.ctx, userLoginParams.Email).Return(&entity.User{ID: 1, Name: "Test", Password: string(hash)}, nil)
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockUserService{
				userRepository:        &mocks.UserRepository{},
				userPackageRepository: &mocks.UserPackageRepository{},
				userPairRepository:    &mocks.UserPairRepository{},
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := service.NewUserService(
				m.userRepository, m.userPairRepository, m.userPackageRepository,
			)
			result, err := subject.Login(tt.in.ctx, *tt.in.params)

			if tt.out.err != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.out.err, err)
			}

			if tt.out.result != nil {
				assert.NotNil(t, result)
			}
		})
	}
}

func TestUserServiceImpl_Register(t *testing.T) {
	type input struct {
		ctx    context.Context
		params *service.UserRegisterParams
	}

	type output struct {
		result *service.UserLoginResponse
		err    error
	}

	type testcase struct {
		name string
		in   *input
		out  *output
		on   func(*MockUserService, *input, *output)
	}

	userRegisterParams := &service.UserRegisterParams{
		Name:     "TESTT",
		Email:    "testmail@yopmail.com",
		Password: "123456",
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(userRegisterParams.Password), bcrypt.DefaultCost)
	userLoginRes, _ := service.CreateSession(&entity.User{ID: 1, Name: "Test", Password: string(hash)})

	for _, tt := range []testcase{
		{
			name: "INTERNAL_SERVER_ERROR",
			in: &input{
				ctx:    context.Background(),
				params: userRegisterParams,
			},
			out: &output{
				result: nil,
				err:    errors.New(""),
			},
			on: func(m *MockUserService, i *input, o *output) {
				m.userRepository.On("GetByEmail", i.ctx, userRegisterParams.Email).Return(nil, o.err)
			},
		},
		{
			name: "EMAIL_ALREADY_REGISTERED",
			in: &input{
				ctx:    context.Background(),
				params: userRegisterParams,
			},
			out: &output{
				result: nil,
				err:    errors.New("Email already registered"),
			},
			on: func(m *MockUserService, i *input, o *output) {
				m.userRepository.On("GetByEmail", i.ctx, userRegisterParams.Email).Return(&entity.User{ID: 1, Name: "Test", Password: "sss"}, nil)
			},
		},
		{
			name: "INTERNAL_SERVER_ERROR",
			in: &input{
				ctx:    context.Background(),
				params: userRegisterParams,
			},
			out: &output{
				result: nil,
				err:    errors.New(""),
			},
			on: func(m *MockUserService, i *input, o *output) {
				m.userRepository.On("GetByEmail", i.ctx, userRegisterParams.Email).Return(nil, nil)
				m.userRepository.On("Create", i.ctx, mock.AnythingOfType("*entity.User")).Return(nil, o.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx:    context.Background(),
				params: userRegisterParams,
			},
			out: &output{
				result: userLoginRes,
				err:    nil,
			},
			on: func(m *MockUserService, i *input, o *output) {
				m.userRepository.On("GetByEmail", i.ctx, userRegisterParams.Email).Return(nil, nil)
				m.userRepository.On("Create", i.ctx, mock.AnythingOfType("*entity.User")).Return(&entity.User{
					Name:     userRegisterParams.Name,
					Email:    userRegisterParams.Email,
					Password: string(hash),
				}, nil)
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockUserService{
				userRepository:        &mocks.UserRepository{},
				userPackageRepository: &mocks.UserPackageRepository{},
				userPairRepository:    &mocks.UserPairRepository{},
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := service.NewUserService(
				m.userRepository, m.userPairRepository, m.userPackageRepository,
			)
			result, err := subject.Register(tt.in.ctx, *tt.in.params)

			if tt.out.err != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.out.err, err)
			}

			if tt.out.result != nil {
				assert.NotNil(t, result)
			}
		})
	}
}
