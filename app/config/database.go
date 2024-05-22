package config

import (
	"fmt"
	"github.com/hendrihmwn/dating-app-api/app/utils"
)

func GetDBConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		utils.GetEnv("DATABASE_USER", "root"),
		utils.GetEnv("DATABASE_PASSWORD", "12345678"),
		utils.GetEnv("DATABASE_HOST", "localhost"),
		utils.GetEnv("DATABASE_PORT", "5432"),
		utils.GetEnv("DATABASE_NAME", "dating-app"),
		utils.GetEnv("DATABASE_SSL_MODE", "disable"),
	)
}
