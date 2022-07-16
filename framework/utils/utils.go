package utils

import (
	"encoding/json"
)

func IsJson(str string) error {
	var js struct{}

	if err := json.Unmarshal([]byte(str), &js); err != nil {
		return err
	}

	return nil
}
