package model

import (
	"github.com/pkg/errors"
)

var ErrInvalidRequest = errors.New("request data is invalid")
var ErrNotFound = errors.New("not found")
