package models

// UserService is a set of methods to manipulate and work with the user model
type UserService interface {
	UserDB
}

// NewUserService ...
func NewUserService() (*userService, error) {
	uo, err := newUserOrm()
	if err != nil {
		return nil, err
	}
	uv := newUserValidator(uo)
	return &userService{
		UserDB: uv,
	}, nil
}

var _ UserService = &userService{}

type userService struct {
	UserDB
}
