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
	"time"
)

type MockUserPackageService struct {
	userPackageRepository *mocks.UserPackageRepository
	packageRepository     *mocks.PackageRepository
	orderRepository       *mocks.OrderRepository
}

func TestUserPackageServiceImpl_Create(t *testing.T) {
	type input struct {
		ctx    context.Context
		params *service.CreateUserPackageParams
	}

	type output struct {
		result *service.CreateUserPackageResponse
		err    error
	}

	type testcase struct {
		name string
		in   *input
		out  *output
		on   func(*MockUserPackageService, *input, *output)
	}

	createUserPackageParams := &service.CreateUserPackageParams{
		UserId:    int64(1),
		PackageId: int64(1),
	}
	now := time.Now()
	expired := now.AddDate(0, 12, 0)

	for _, tt := range []testcase{
		{
			name: "INTERNAL_SERVER_ERROR",
			in: &input{
				ctx:    context.Background(),
				params: createUserPackageParams,
			},
			out: &output{
				result: nil,
				err:    errors.New(""),
			},
			on: func(m *MockUserPackageService, i *input, o *output) {
				m.userPackageRepository.On("Get", i.ctx, i.params.UserId).Return(nil, o.err)
			},
		},
		{
			name: "PACKAGE_ALREADY_PURCHASE",
			in: &input{
				ctx:    context.Background(),
				params: createUserPackageParams,
			},
			out: &output{
				result: nil,
				err:    errors.New("package already purchase"),
			},
			on: func(m *MockUserPackageService, i *input, o *output) {
				m.userPackageRepository.On("Get", i.ctx, i.params.UserId).Return(&entity.UserPackage{PackageID: createUserPackageParams.PackageId}, nil)
			},
		},
		{
			name: "INTERNAL_SERVER_ERROR",
			in: &input{
				ctx:    context.Background(),
				params: createUserPackageParams,
			},
			out: &output{
				result: nil,
				err:    errors.New(""),
			},
			on: func(m *MockUserPackageService, i *input, o *output) {
				m.userPackageRepository.On("Get", i.ctx, i.params.UserId).Return(&entity.UserPackage{}, nil)
				m.packageRepository.On("Get", i.ctx, i.params.PackageId).Return(nil, o.err)
			},
		},
		{
			name: "PACKAGE_NOT_FOUND",
			in: &input{
				ctx:    context.Background(),
				params: createUserPackageParams,
			},
			out: &output{
				result: nil,
				err:    errors.New("package not found"),
			},
			on: func(m *MockUserPackageService, i *input, o *output) {
				m.userPackageRepository.On("Get", i.ctx, i.params.UserId).Return(&entity.UserPackage{}, nil)
				m.packageRepository.On("Get", i.ctx, i.params.PackageId).Return(nil, nil)
			},
		},
		{
			name: "INTERNAL_SERVER_ERROR",
			in: &input{
				ctx:    context.Background(),
				params: createUserPackageParams,
			},
			out: &output{
				result: nil,
				err:    errors.New("package not found"),
			},
			on: func(m *MockUserPackageService, i *input, o *output) {
				m.userPackageRepository.On("Get", i.ctx, i.params.UserId).Return(&entity.UserPackage{}, nil)
				m.packageRepository.On("Get", i.ctx, i.params.PackageId).Return(&entity.Package{
					ID:          1,
					Name:        "Unlimited Swipe",
					Act:         "unlimited-swipe",
					Price:       100000,
					ValidMonths: 12,
				}, nil)
				m.orderRepository.On("Create", i.ctx, mock.AnythingOfType("*entity.Order")).Return(nil, o.err)
			},
		},
		{
			name: "INTERNAL_SERVER_ERROR",
			in: &input{
				ctx:    context.Background(),
				params: createUserPackageParams,
			},
			out: &output{
				result: nil,
				err:    errors.New("package not found"),
			},
			on: func(m *MockUserPackageService, i *input, o *output) {
				m.userPackageRepository.On("Get", i.ctx, i.params.UserId).Return(&entity.UserPackage{}, nil)
				m.packageRepository.On("Get", i.ctx, i.params.PackageId).Return(&entity.Package{
					ID:          1,
					Name:        "Unlimited Swipe",
					Act:         "unlimited-swipe",
					Price:       100000,
					ValidMonths: 12,
				}, nil)
				m.orderRepository.On("Create", i.ctx, mock.AnythingOfType("*entity.Order")).Return(nil, nil)
				m.userPackageRepository.On("Create", i.ctx, mock.AnythingOfType("*entity.UserPackage")).Return(nil, o.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx:    context.Background(),
				params: createUserPackageParams,
			},
			out: &output{
				result: &service.CreateUserPackageResponse{
					ID:         1,
					UserId:     1,
					PackageId:  1,
					PackageAct: "unlimited-swipe",
					ExpiredAt:  expired,
				},
				err: nil,
			},
			on: func(m *MockUserPackageService, i *input, o *output) {
				m.userPackageRepository.On("Get", i.ctx, i.params.UserId).Return(&entity.UserPackage{}, nil)
				m.packageRepository.On("Get", i.ctx, i.params.PackageId).Return(&entity.Package{
					ID:          1,
					Name:        "Unlimited Swipe",
					Act:         "unlimited-swipe",
					Price:       100000,
					ValidMonths: 12,
				}, nil)
				m.orderRepository.On("Create", i.ctx, mock.AnythingOfType("*entity.Order")).Return(nil, nil)
				m.userPackageRepository.On("Create", i.ctx, mock.AnythingOfType("*entity.UserPackage")).Return(&entity.UserPackage{
					ID:         1,
					UserID:     1,
					PackageID:  1,
					ExpiredAt:  &expired,
					PackageAct: "unlimited-swipe",
				}, nil)
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockUserPackageService{
				packageRepository:     &mocks.PackageRepository{},
				userPackageRepository: &mocks.UserPackageRepository{},
				orderRepository:       &mocks.OrderRepository{},
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := service.NewUserPackageService(
				m.userPackageRepository, m.packageRepository, m.orderRepository,
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
