package data

import (
	"fmt"
	"strconv"
)

type CookingTime int32

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
