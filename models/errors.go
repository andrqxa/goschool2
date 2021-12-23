package models

import "errors"

var (
	// ErrNotFound is returned when a resource can't be found in the database
	ErrNotFound = errors.New("models: resource not found")

	// ErrEmailRequired is returned when an email address is not provided when
	// creating a user
	ErrEmailRequired = errors.New("models: email address is required")

	// ErrEmailInvalid is returned when an email address provided does not match
	// any od our requirements
	ErrEmailInvalid = errors.New("models: email address is not valid")

	// ErrEmailTaken is returned when an update or create is attempted with
	// an email address that is already in use.
	ErrEmailTaken = errors.New("models: email address is already taken")

	// ErrNameRequired is returned when a name is not provided when
	// creating a user
	ErrNameRequired = errors.New("models: name is required")

	// ErrNameInvalid is returned when a name provided does not match
	// any of our requirements
	ErrNameInvalid = errors.New("models: name is not valid")
)
