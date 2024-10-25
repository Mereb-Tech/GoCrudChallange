package domain

import (
	"errors"
)

var ErrNoRecord = errors.New("entity: No record/s were found")
var ErrDuplicateId = errors.New("entity: Duplicate Id")
