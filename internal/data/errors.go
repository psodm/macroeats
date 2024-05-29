package data

import (
	"errors"

	"github.com/lib/pq"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrDuplicateRow   = errors.New("record already exists")
)

func isDuplicate(err error) bool {
	pqErr, ok := err.(*pq.Error)
	if ok {
		if pqErr.Code == "23505" {
			return true
		}
	}
	return false
}
