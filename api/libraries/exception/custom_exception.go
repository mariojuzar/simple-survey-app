package exception

import "errors"

type notFoundException struct {

}

func (notFoundException) Error() string  {
	return "Not Found"
}

func NotFoundException() error  {
	return notFoundException{}
}

func RequiredFieldException(field string) error {
	return errors.New("Required " + field)
}