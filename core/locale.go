package core

import (
	"encoding/json"
	"os"

	"github.com/agaUHO/aga/system"
)

func GetTextMessage(keyMessage ...string) string {
	jsonFile, err := os.ReadFile("./locale/" + system.Language + ".json")
	if err != nil {
		return "/locale/es.json not found"
	}
	var data map[string]interface{}
	err = json.Unmarshal(jsonFile, &data)
	if err != nil {
		return "format error in locale file"
	}
	var msg string
	for i, key := range keyMessage {
		value, ok := data[key]
		result, _ := value.(string)
		if !ok {
			result += key
		}
		if i > 0 {
			msg += ". " + result
		} else {
			msg += result
		}
	}
	return msg
}
