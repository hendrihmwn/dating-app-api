package service_test

import (
	"context"
	"errors"
	"github.com/hendrihmwn/dating-app-api/app/domain/entity"
	"github.com/hendrihmwn/dating-app-api/app/domain/service"
	mocks "github.com/hendrihmwn/dating-app-api/mocks/domain/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockUserPairService struct {
	userPairRepository    *mocks.UserPairRepository
	userPackageRepository *mocks.UserPackageRepository
}

func TestUserPairServiceImpl_Create(t *testing.T) {
	type input struct {
		ctx    context.Context
		params *service.CreateUserPairParams
	}

	type output struct {
		result *service.CreateUserPairResponse
		err    error
	}

	type testcase struct {
		name string
		in   *input
		out  *output
		on   func(*MockUserPairService, *input, *output)
	}

	createUserPairParams := &service.CreateUserPairParams{
		UserId:     int64(1),
		PairUserId: int64(1),
		Status:     1,
	}

	for _, tt := range []testcase{
		{
			name: "INTERNAL_SERVER_ERROR",
			in: &input{
				ctx:    context.Background(),
				params: createUserPairParams,
			},
			out: &output{
				result: nil,
				err:    errors.New(""),
			},
			on: func(m *MockUserPairService, i *input, o *output) {
				m.userPackageRepository.On("List", i.ctx, i.params.UserId).Return(nil, o.err)
			},
		},
		{
			name: "INTERNAL_SERVER_ERROR",
			in: &input{
				ctx:    context.Background(),
				params: createUserPairParams,
			},
			out: &output{
				result: nil,
				err:    errors.New(""),
			},
			on: func(m *MockUserPairService, i *input, o *output) {
				m.userPackageRepository.On("List", i.ctx, i.params.UserId).Return([]entity.UserPackage{{
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
				params: createUserPairParams,
			},
			out: &output{
				result: nil,
				err:    errors.New("max limit 10 has reach"),
			},
			on: func(m *MockUserPairService, i *input, o *output) {
				m.userPackageRepository.On("List", i.ctx, i.params.UserId).Return([]entity.UserPackage{}, nil)
				m.userPairRepository.On("CountToday", i.ctx, int64(1)).Return(11, nil)
			},
		},
		{
			name: "INTERNAL_SERVER_ERROR",
			in: &input{
				ctx:    context.Background(),
				params: createUserPairParams,
			},
			out: &output{
				result: nil,
				err:    errors.New(""),
			},
			on: func(m *MockUserPairService, i *input, o *output) {
				m.userPackageRepository.On("List", i.ctx, i.params.UserId).Return([]entity.UserPackage{}, nil)
				m.userPairRepository.On("CountToday", i.ctx, int64(1)).Return(5, nil)
				m.userPairRepository.On("Create", i.ctx, mock.AnythingOfType("*entity.UserPair")).Return(nil, o.err)
			},
		},
		{
			name: "INTERNAL_SERVER_ERROR",
			in: &input{
				ctx:    context.Background(),
				params: createUserPairParams,
			},
			out: &output{
				result: nil,
				err:    nil,
			},
			on: func(m *MockUserPairService, i *input, o *output) {
				m.userPackageRepository.On("List", i.ctx, i.params.UserId).Return([]entity.UserPackage{}, nil)
				m.userPairRepository.On("CountToday", i.ctx, i.params.UserId).Return(5, nil)
				m.userPairRepository.On("Create", i.ctx, mock.AnythingOfType("*entity.UserPair")).Return(&entity.UserPair{
					ID:         1,
					UserID:     1,
					PairUserID: 2,
					Status:     1,
				}, nil)
				m.userPairRepository.On("CountToday", i.ctx, i.params.UserId).Return(0, o.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx:    context.Background(),
				params: createUserPairParams,
			},
			out: &output{
				result: &service.CreateUserPairResponse{
					ID:         1,
					UserId:     1,
					PairUserId: 2,
					Status:     1,
				},
				err: nil,
			},
			on: func(m *MockUserPairService, i *input, o *output) {
				m.userPackageRepository.On("List", i.ctx, i.params.UserId).Return([]entity.UserPackage{}, nil)
				m.userPairRepository.On("CountToday", i.ctx, int64(1)).Return(5, nil)
				m.userPairRepository.On("Create", i.ctx, mock.AnythingOfType("*entity.UserPair")).Return(&entity.UserPair{
					ID:         1,
					UserID:     1,
					PairUserID: 2,
					Status:     1,
				}, nil)
				m.userPairRepository.On("CountToday", i.ctx, int64(1)).Return(6, nil)
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockUserPairService{
				userPairRepository:    &mocks.UserPairRepository{},
				userPackageRepository: &mocks.UserPackageRepository{},
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := service.NewUserPairService(
				m.userPairRepository, m.userPackageRepository,
			)
			result, err := subject.Create(tt.in.ctx, *tt.in.params)

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
