package platform

import (
	"errors"
)

//AccessEnum used by Assets
type AccessEnum string

const (

	//PUBLIC_READ defines constant for the `public-read`
	PUBLIC_READ AccessEnum = "public-read"

	//PRIVATE defines constant for the `private`
	PRIVATE AccessEnum = "private"
)

//IsValid return error if enum is invalid
func (ac AccessEnum) IsValid() error {
	switch ac {
	case PUBLIC_READ, PRIVATE:
		return nil
	}
	return errors.New("Invalid AccessEnum type")
}
