package common

import "errors"

var (
    ErrUserNotFound       = errors.New("user not found")
    ErrUnauthorized      = errors.New("unauthorized")
    ErrConditionFailed   = errors.New("condition evaluation failed")
)
