package validation

import (
	"errors"
	"strconv"
)

func ValidateInputNotNil(data ...interface{}) error {
	for _, value := range data {
		switch result := value.(type) {
		case string:
			if len(result) == 0 {
				return errors.New("Data Input Cannot Empty")
			}
		case int:
			if result == 0 {
				return errors.New("Data Input Cannot 0")
			}
		}
	}
	return nil
}

func ValidateInputNumber(data interface{}) error {
	switch result := data.(type) {
	case string:
		if _, err := strconv.Atoi(result); err != nil {
			return errors.New("ID Cannot Contain Characters")
		}
	default:
		return errors.New("ID input not an INT data type")
	}
	return nil
}
