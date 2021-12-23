package models

// UserDB is used to interact with users database.
//
// For pretty much all single user queries:
// If the user is found, we'll return a nill error
// If the user isn't found, we'll return ErrNotFound
// If there is another error, we'll return an error with more information about
// what went wrong. This may not be an error generated by the models package.
//
// For single user queries, any error but ErrorNotFound should probably result in
// 500 error.
type UserDB interface {
	// Methods for quering for single users
	ByEmail(email string) (*User, error)

	// Methods for altering users
	Create(user *User) error
	Update(user *User) error
	Delete(email string) error
}
