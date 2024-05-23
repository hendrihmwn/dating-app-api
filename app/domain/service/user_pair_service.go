package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/hendrihmwn/dating-app-api/app/domain/entity"
	"github.com/hendrihmwn/dating-app-api/app/domain/repository"
	"time"
)

type CreateUserPairParams struct {
	UserId     int64
	PairUserId int64
	Status     int
}

type CreateUserPairResponse struct {
	ID         int64
	UserId     int64
	PairUserId int64
	Status     int
	StatusLang string
	MaxLimit   *string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type UserPairService interface {
	Create(ctx context.Context, params CreateUserPairParams) (*CreateUserPairResponse, error)
}

type UserPairServiceImpl struct {
	userPairRepository    repository.UserPairRepository
	userPackageRepository repository.UserPackageRepository
}

func (p *UserPairServiceImpl) Create(ctx context.Context, params CreateUserPairParams) (*CreateUserPairResponse, error) {

	userPackages, err := p.userPackageRepository.List(ctx, params.UserId)
	if err != nil {
		return nil, err
	}

	unlimitedSwipe := false
	if len(userPackages) > 0 {
		for _, userP := range userPackages {
			if userP.PackageAct == "unlimited-swipe" {
				unlimitedSwipe = true
			}
		}
	}
	countBefore, err := p.userPairRepository.CountToday(ctx, params.UserId)
	if err != nil {
		return nil, err
	}

	if !unlimitedSwipe {
		if countBefore >= 10 {
			return nil, errors.New("max limit 10 has reach")
		}
	}

	userPair, err := p.userPairRepository.Create(ctx, &entity.UserPair{
		UserID:     params.UserId,
		PairUserID: params.PairUserId,
		Status:     params.Status,
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	response := &CreateUserPairResponse{
		ID:         userPair.ID,
		UserId:     userPair.UserID,
		PairUserId: userPair.PairUserID,
		Status:     userPair.Status,
		StatusLang: userPair.StatusLang,
		CreatedAt:  userPair.CreatedAt,
		UpdatedAt:  userPair.UpdatedAt,
	}

	countAfter, err := p.userPairRepository.CountToday(ctx, params.UserId)
	if err != nil {
		return nil, err
	}

	if !unlimitedSwipe {
		if countAfter >= 10 {
			maxLimit := "max limit 10 has reach"
			response.MaxLimit = &maxLimit
		}
	}

	return response, nil
}

func NewUserPairService(userPairRepository repository.UserPairRepository, userPackageRepository repository.UserPackageRepository) UserPairService {
	return &UserPairServiceImpl{
		userPackageRepository: userPackageRepository,
		userPairRepository:    userPairRepository,
	}
}
