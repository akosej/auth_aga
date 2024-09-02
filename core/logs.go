package core

import (
	"github.com/agaUHO/aga/database"
	"github.com/agaUHO/aga/models"
)

func SendLogs(module, action, login, uid, description, app string) {
	database.DB.Create(&models.Logs{
		Module:      module,
		Action:      action,
		Login:       login,
		Uid:         uid,
		Description: description,
		App:         app,
	})
}
