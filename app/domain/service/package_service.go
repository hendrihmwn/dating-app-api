package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/hendrihmwn/dating-app-api/app/domain/entity"
	"github.com/hendrihmwn/dating-app-api/app/domain/repository"
	"time"
)

type PackageListResponse struct {
	Data []entity.Package
}

type GetPackageParams struct {
	ID int64
}

type GetPackageResponse struct {
	ID          int64
	Name        string
	Act         string
	Price       int64
	ValidMonths int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type PackageService interface {
	List(ctx context.Context) (*PackageListResponse, error)
	Get(ctx context.Context, params GetPackageParams) (*GetPackageResponse, error)
}

type PackageServiceImpl struct {
	packageRepository repository.PackageRepository
}

func (p *PackageServiceImpl) List(ctx context.Context) (*PackageListResponse, error) {
	packages, err := p.packageRepository.List(ctx)
	if err != nil {
		return nil, err
	}

	return &PackageListResponse{
		Data: packages,
	}, nil
}

func (p *PackageServiceImpl) Get(ctx context.Context, params GetPackageParams) (*GetPackageResponse, error) {
	packages, err := p.packageRepository.Get(ctx, params.ID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if packages == nil {
		return nil, errors.New("data not found")
	}

	return &GetPackageResponse{
		ID:          packages.ID,
		Name:        packages.Name,
		Act:         packages.Act,
		Price:       packages.Price,
		ValidMonths: packages.ValidMonths,
		CreatedAt:   packages.CreatedAt,
		UpdatedAt:   packages.UpdatedAt,
	}, nil
}

func NewPackageService(packageRepository repository.PackageRepository) PackageService {
	return &PackageServiceImpl{
		packageRepository: packageRepository,
	}
}
