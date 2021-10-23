package exception

import "errors"

var (
	NotFoundError = errors.New("record is not found")
	AlreadyExistError = errors.New("record is already exist")
	UnauthorizedError = errors.New("data is not valid")
)
