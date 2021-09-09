package errmsg

import "errors"

func EXEC_ERROR(err error) error {
	return errors.New("EXEC_ERROR : "+ err.Error())
}