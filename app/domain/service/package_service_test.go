package service_test

import (
	"context"
	"errors"
	"github.com/hendrihmwn/dating-app-api/app/domain/entity"
	"github.com/hendrihmwn/dating-app-api/app/domain/service"
	mocks "github.com/hendrihmwn/dating-app-api/mocks/domain/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockPackageService struct {
	packageRepository *mocks.PackageRepository
}

func TestPackageServiceImpl_Get(t *testing.T) {
	type input struct {
		ctx    context.Context
		params *service.GetPackageParams
	}

	type output struct {
		result *service.GetPackageResponse
		err    error
	}

	type testcase struct {
		name string
		in   *input
		out  *output
		on   func(*MockPackageService, *input, *output)
	}

	getPackageParams := &service.GetPackageParams{
		ID: int64(1),
	}

	for _, tt := range []testcase{
		{
			name: "INTERNAL_SERVER_ERROR",
			in: &input{
				ctx:    context.Background(),
				params: getPackageParams,
			},
			out: &output{
				result: nil,
				err:    errors.New(""),
			},
			on: func(m *MockPackageService, i *input, o *output) {
				m.packageRepository.On("Get", i.ctx, i.params.ID).Return(nil, o.err)
			},
		},
		{
			name: "INTERNAL_SERVER_ERROR",
			in: &input{
				ctx:    context.Background(),
				params: getPackageParams,
			},
			out: &output{
				result: nil,
				err:    errors.New("data not found"),
			},
			on: func(m *MockPackageService, i *input, o *output) {
				m.packageRepository.On("Get", i.ctx, i.params.ID).Return(nil, nil)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx:    context.Background(),
				params: getPackageParams,
			},
			out: &output{
				result: &service.GetPackageResponse{
					ID:          1,
					Name:        "Unlimited Swipe",
					Act:         "unlimited-swipe",
					Price:       100000,
					ValidMonths: 12,
				},
			},
			on: func(m *MockPackageService, i *input, o *output) {
				m.packageRepository.On("Get", i.ctx, i.params.ID).Return(&entity.Package{
					ID:          1,
					Name:        "Unlimited Swipe",
					Act:         "unlimited-swipe",
					Price:       100000,
					ValidMonths: 12,
				}, nil)
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockPackageService{
				packageRepository: &mocks.PackageRepository{},
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := service.NewPackageService(
				m.packageRepository,
			)
			result, err := subject.Get(tt.in.ctx, *tt.in.params)

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

func TestPackageServiceImpl_List(t *testing.T) {
	type input struct {
		ctx context.Context
	}

	type output struct {
		result *service.PackageListResponse
		err    error
	}

	type testcase struct {
		name string
		in   *input
		out  *output
		on   func(*MockPackageService, *input, *output)
	}

	for _, tt := range []testcase{
		{
			name: "INTERNAL_SERVER_ERROR",
			in: &input{
				ctx: context.Background(),
			},
			out: &output{
				result: nil,
				err:    errors.New(""),
			},
			on: func(m *MockPackageService, i *input, o *output) {
				m.packageRepository.On("List", i.ctx).Return(nil, o.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
			},
			out: &output{
				result: &service.PackageListResponse{
					Data: []entity.Package{{
						ID:          1,
						Name:        "Unlimited Swipe",
						Act:         "unlimited-swipe",
						Price:       100000,
						ValidMonths: 12,
					}},
				},
			},
			on: func(m *MockPackageService, i *input, o *output) {
				m.packageRepository.On("List", i.ctx).Return([]entity.Package{{
					ID:          1,
					Name:        "Unlimited Swipe",
					Act:         "unlimited-swipe",
					Price:       100000,
					ValidMonths: 12,
				}}, nil)
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockPackageService{
				packageRepository: &mocks.PackageRepository{},
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := service.NewPackageService(
				m.packageRepository,
			)
			result, err := subject.List(tt.in.ctx)

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
