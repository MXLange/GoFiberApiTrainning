package initializer

import (
	appError "github.com/api/interface/internal/error"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	appError.FatalError(err)
}
