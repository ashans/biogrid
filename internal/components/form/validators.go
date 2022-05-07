package form

import (
	"errors"
	"strconv"
)

func seqValidator(s string) error {
	if len(s) == 0 {
		return errors.New("should not be empty")
	}
	if len(s) > 6 {
		return errors.New("should not exceed 6")
	}
	return nil
}

func intValidator(s string) error {
	_, err := strconv.Atoi(s)
	if err != nil {
		return errors.New("should only be integer")
	}
	return nil
}
