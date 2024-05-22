package config

import (
	"github.com/hendrihmwn/dating-app-api/app/utils"
	"strconv"
	"time"
)

func GetAccessTokenSecretKey() string {
	return utils.GetEnv("ACCESS_TOKEN_SECRET_KEY", "dating app solusi cari pasangan")
}

func GetAccessTokenExpirationTime() time.Duration {
	val := utils.GetEnv("ACCESS_TOKEN_EXPIRATION_TIME", "86400")
	exp, _ := strconv.ParseUint(val, 10, 32)
	return time.Second * time.Duration(exp)
}

func GetRefreshTokenSecretKey() string {
	return utils.GetEnv("REFRESH_TOKEN_SECRET_KEY", "dating app semoga dapat yang diinginkan")
}

func GetRefreshTokenExpirationTime() time.Duration {
	val := utils.GetEnv("REFRESH_TOKEN_EXPIRATION_TIME", "259200")
	exp, _ := strconv.ParseUint(val, 10, 32)
	return time.Second * time.Duration(exp)
}
