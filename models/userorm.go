package models

var _ UserDB = &userOrm{}

func newUserOrm() (*userOrm, error) {
	db := make(map[string]User)
	return &userOrm{
		db: db,
	}, nil
}

type userOrm struct {
	db map[string]User
}

// ByEmail will look up a user with the given email address and return that user
// If the user is found, we'll return a nill error
// If the user isn't found, we'll return ErrNotFound
//
// As a general rule, any error but ErrorNotFound should probably result in
// 500 error.
func (uo *userOrm) ByEmail(email string) (*User, error) {
	user, ok := uo.db[email]
	if !ok {
		return nil, ErrNotFound
	}
	return &user, nil
}

// Create will create the provided user
func (uo *userOrm) Create(user *User) error {
	uo.db[user.Email] = *user
	return nil
}

// Update will update the provided user
func (uo *userOrm) Update(user *User) error {
	uo.db[user.Email] = *user
	return nil
}

// Delete will delete the user with the provided email
func (uo *userOrm) Delete(email string) error {
	_, err := uo.ByEmail(email)
	if err != nil {
		return err
	}
	delete(uo.db, email)
	return nil
}
