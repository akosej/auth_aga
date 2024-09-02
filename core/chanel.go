package core

import (
	"reflect"

	"github.com/agaUHO/aga/system"
)

// InitializeChannelListening  -- Initialize channel listening
// NOTE En esta función se añadirán todos los canales que se creen
func InitializeChannelListening() {
	for {
		select {
		case log := <-system.Log:
			requestValue := reflect.ValueOf(log)
			SendLogs(
				requestValue.FieldByName("Module").String(),
				requestValue.FieldByName("Action").String(),
				requestValue.FieldByName("Login").String(),
				requestValue.FieldByName("Uid").String(),
				requestValue.FieldByName("Description").String(),
				requestValue.FieldByName("App").String(),
			)
		}
	}
}
