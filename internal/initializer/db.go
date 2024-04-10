package initializer

import (
	db "github.com/api/interface/internal/database"
	appError "github.com/api/interface/internal/error"
)

func ConnectDB() {
	err := db.Connect()
	appError.FatalError(err)
}
