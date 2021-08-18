package models

import "errors"

var ErrNotFoundUser = errors.New("user not found")
var ErrCannotCreate = errors.New("cannot create user")
var ErrExists = errors.New("user is exist")
var ErrInternal = errors.New("internal error")
var ErrInvalidUser = errors.New("invalid username or password")