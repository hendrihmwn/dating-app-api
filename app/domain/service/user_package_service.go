package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/hendrihmwn/dating-app-api/app/domain/entity"
	"github.com/hendrihmwn/dating-app-api/app/domain/repository"
	"time"
)

type ListUserPackageParams struct {
	UserId int64
}

type ListUserPackageResponse struct {
	Data []entity.UserPackage
}

type CreateUserPackageParams struct {
	UserId    int64
	PackageId int64
}

type CreateUserPackageResponse struct {
	ID         int64
	UserId     int64
	PackageId  int64
	PackageAct string
	ExpiredAt  time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type UserPackageService interface {
	List(ctx context.Context, params ListUserPackageParams) (*ListUserPackageResponse, error)
	Create(ctx context.Context, params CreateUserPackageParams) (*CreateUserPackageResponse, error)
}

type UserPackageServiceImpl struct {
	userPackageRepository repository.UserPackageRepository
	packageRepository     repository.PackageRepository
	orderRepository       repository.OrderRepository
}

func (p *UserPackageServiceImpl) List(ctx context.Context, params ListUserPackageParams) (*ListUserPackageResponse, error) {
	packages, err := p.userPackageRepository.List(ctx, params.UserId)
	if err != nil {
		return nil, err
	}

	return &ListUserPackageResponse{
		Data: packages,
	}, nil
}

func (p *UserPackageServiceImpl) Create(ctx context.Context, params CreateUserPackageParams) (*CreateUserPackageResponse, error) {
	usrPkg, err := p.userPackageRepository.Get(ctx, params.UserId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if usrPkg != nil && usrPkg.PackageID == params.PackageId {
		return nil, errors.New("package already purchase")
	}

	pkg, err := p.packageRepository.Get(ctx, params.PackageId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if pkg == nil {
		return nil, errors.New("package not found")
	}

	_, err = p.orderRepository.Create(ctx, &entity.Order{
		UserID:       params.UserId,
		PackageName:  pkg.Name,
		PackagePrice: pkg.Price,
		Status:       1,
	})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	now := time.Now()
	expired := now.AddDate(0, pkg.ValidMonths, 0)

	userPackage, err := p.userPackageRepository.Create(ctx, &entity.UserPackage{
		UserID:     params.UserId,
		PackageID:  pkg.ID,
		PackageAct: pkg.Act,
		ExpiredAt:  &expired,
	})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &CreateUserPackageResponse{
		ID:         userPackage.ID,
		UserId:     userPackage.UserID,
		PackageId:  userPackage.PackageID,
		PackageAct: userPackage.PackageAct,
		ExpiredAt:  *userPackage.ExpiredAt,
		CreatedAt:  userPackage.CreatedAt,
		UpdatedAt:  userPackage.UpdatedAt,
	}, nil
}

func NewUserPackageService(userPackageRepository repository.UserPackageRepository, packageRepository repository.PackageRepository,
	orderRepository repository.OrderRepository) UserPackageService {
	return &UserPackageServiceImpl{
		userPackageRepository: userPackageRepository,
		packageRepository:     packageRepository,
		orderRepository:       orderRepository,
	}
}
