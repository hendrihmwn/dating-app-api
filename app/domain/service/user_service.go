package service

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/hendrihmwn/dating-app-api/app/config"
	"github.com/hendrihmwn/dating-app-api/app/domain/entity"
	"github.com/hendrihmwn/dating-app-api/app/domain/repository"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type UserServiceImpl struct {
	userRepository        repository.UserRepository
	userPairRepository    repository.UserPairRepository
	userPackageRepository repository.UserPackageRepository
}

type UserListParams struct {
	Limit  int
	UserId int64
}
type UserListResponse struct {
	Data []entity.User
}

type UserLoginParams struct {
	Email    string
	Password string
}
type UserLoginResponse struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiredAt    time.Time `json:"expired_at"`
}

type UserRegisterParams struct {
	Name     string
	Email    string
	Password string
}

type UserService interface {
	List(ctx context.Context, params UserListParams) (*UserListResponse, error)
	Login(ctx context.Context, params UserLoginParams) (*UserLoginResponse, error)
	Register(ctx context.Context, params UserRegisterParams) (*UserLoginResponse, error)
}

func (u *UserServiceImpl) List(ctx context.Context, params UserListParams) (*UserListResponse, error) {
	userPackages, err := u.userPackageRepository.List(ctx, params.UserId)
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
	countToday, err := u.userPairRepository.CountToday(ctx, params.UserId)

	if err != nil {
		return nil, err
	}

	if !unlimitedSwipe {
		if countToday > 10 {
			return nil, errors.New("max limit 10 has reach")
		}
	}

	users, err := u.userRepository.List(ctx, &repository.ListUsersArgs{
		Limit:  params.Limit,
		Offset: 0,
		Filter: &repository.UserFilterArgs{
			ExcludeUserId: []int64{params.UserId},
		}})
	if err != nil {
		return nil, err
	}

	return &UserListResponse{Data: users}, nil
}

type UserClaims struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

type RefreshTokenClaims struct {
	AccessToken string `json:"access_token"`
	jwt.StandardClaims
}

func (u *UserServiceImpl) Login(ctx context.Context, params UserLoginParams) (*UserLoginResponse, error) {
	user, err := u.userRepository.GetByEmail(ctx, params.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("Email not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))
	if err != nil {
		return nil, errors.New("Wrong password")
	}

	session, err := CreateSession(user)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (u *UserServiceImpl) Register(ctx context.Context, params UserRegisterParams) (*UserLoginResponse, error) {
	email := strings.ToLower(params.Email)
	user, err := u.userRepository.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if user != nil {
		return nil, errors.New("Email already registered")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user, err = u.userRepository.Create(ctx, &entity.User{
		Name:     params.Name,
		Email:    email,
		Password: string(hash),
	})
	if err != nil {
		return nil, err
	}

	session, err := CreateSession(user)
	if err != nil {
		return nil, err
	}

	return session, nil

}

func CreateSession(user *entity.User) (*UserLoginResponse, error) {
	now := time.Now()
	accessTokenExpiredAt := now.Add(config.GetAccessTokenExpirationTime())

	accessTokenClaim := &UserClaims{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessTokenExpiredAt.Unix(),
			IssuedAt:  now.Unix(),
		},
	}

	session := new(UserLoginResponse)
	session.ExpiredAt = accessTokenExpiredAt
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaim).SignedString([]byte(config.GetAccessTokenSecretKey()))
	session.AccessToken = accessToken
	if err != nil {
		return nil, err
	}

	refreshTokenExpiredAt := now.Add(config.GetRefreshTokenExpirationTime())
	refreshTokenClaim := &RefreshTokenClaims{
		AccessToken: session.AccessToken,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshTokenExpiredAt.Unix(),
			IssuedAt:  now.Unix(),
		},
	}
	session.RefreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaim).SignedString([]byte(config.GetRefreshTokenSecretKey()))
	if err != nil {
		return nil, err
	}
	return session, nil
}

func NewUserService(
	userRepository repository.UserRepository,
	userPairRepository repository.UserPairRepository,
	userPackageRepository repository.UserPackageRepository) UserService {
	return &UserServiceImpl{
		userRepository:        userRepository,
		userPairRepository:    userPairRepository,
		userPackageRepository: userPackageRepository,
	}
}
