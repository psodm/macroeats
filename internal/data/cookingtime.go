package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidCookingTimeFormat = errors.New("invalid cooking time format")

type CookingTime int64

func (c CookingTime) MarshalJSON() ([]byte, error) {
	var jsonValue string
	if c < 60 {
		jsonValue = fmt.Sprintf("%d min", c)
	} else {
		jsonValue = fmt.Sprintf("%d hr, %d min", c/60, c%60)
	}
	quotedCookingTime := strconv.Quote(jsonValue)
	return []byte(quotedCookingTime), nil
}

func (c *CookingTime) UnmarshalJSON(jsonValue []byte) error {
	var totalMinutes int64
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidCookingTimeFormat
	}
	parts := strings.Split(unquotedJSONValue, " ")
	if len(parts) != 2 && len(parts) != 4 {
		fmt.Println("I'm here")
		return ErrInvalidCookingTimeFormat
	}
	if len(parts) == 2 {
		if !strings.Contains(parts[1], "min") {
			return ErrInvalidCookingTimeFormat
		}
		min, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			return err
		}
		totalMinutes += min
	} else {
		if !strings.Contains(parts[1], "hr") && !strings.Contains(parts[3], "min") {
			return ErrInvalidCookingTimeFormat
		}
		hr, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			return err
		}
		totalMinutes += hr * 60
		min, err := strconv.ParseInt(parts[3], 10, 64)
		if err != nil {
			return err
		}
		totalMinutes += min
	}
	*c = CookingTime(totalMinutes)
	return nil
}
