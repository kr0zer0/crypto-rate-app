package customerrors

import "errors"

var ErrEmailDuplicate = errors.New("this email already exists")
