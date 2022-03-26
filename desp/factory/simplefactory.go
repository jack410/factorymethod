package factory

import "github.com/jack410/factory_patterns/desp/models"

const (
	FrontUser = iota
	Aminuser
)

type UserType int

func CreateUser(t UserType) models.UserCreateFunc {
	switch t {
	case FrontUser:
		return models.NewUser()
	case Aminuser:
		return models.NewAdmin()
	default:
		return models.NewUser()
	}
}
