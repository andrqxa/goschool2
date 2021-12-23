package models

import (
	"regexp"
	"strings"
)

type userValFunc func(*User) error

func runUserValFuncs(user *User, fns ...userValFunc) error {
	for _, fn := range fns {
		if err := fn(user); err != nil {
			return err
		}
	}
	return nil
}

var _ UserDB = &userValidator{}

func newUserValidator(udb UserDB) *userValidator {
	return &userValidator{
		UserDB:     udb,
		emailRegex: regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,16}$`),
		nameRegex:  regexp.MustCompile(`^[a-zA-Z\.\s]+$`),
	}
}

type userValidator struct {
	UserDB
	emailRegex *regexp.Regexp
	nameRegex  *regexp.Regexp
}

// ByEmail will normalize the email address before calling ByEmail on the userDB
// field.
func (uv *userValidator) ByEmail(email string) (*User, error) {
	user := User{
		Email: email,
	}
	if err := runUserValFuncs(&user, uv.normalizeEmail); err != nil {
		return nil, err
	}
	return uv.UserDB.ByEmail(user.Email)
}

// Create will create the provided user and backfill data like ID, CreatedAt,
// and UpdatedAt fields
func (uv *userValidator) Create(user *User) error {
	err := runUserValFuncs(user,
		uv.normalizeEmail,
		uv.requireEmail,
		uv.emailFormat,
		uv.emaiIslAvail,
		uv.normalizeName,
		uv.requireEmail,
		uv.nameFormat,
	)
	if err != nil {
		return err
	}
	return uv.UserDB.Create(user)
}

// Update will hash a remember token if it is provided.
func (uv *userValidator) Update(user *User) error {
	err := runUserValFuncs(user,
		uv.normalizeEmail,
		uv.requireEmail,
		uv.emailFormat,
		uv.emaiIslAvail,
		uv.normalizeName,
		uv.requireEmail,
		uv.nameFormat,
	)
	if err != nil {
		return err
	}
	return uv.UserDB.Update(user)
}

// Delete will delete the user with the provided ID
func (uv *userValidator) Delete(email string) error {
	return uv.UserDB.Delete(email)
}

func (uv *userValidator) normalizeEmail(user *User) error {
	user.Email = strings.ToLower(user.Email)
	user.Email = strings.TrimSpace(user.Email)
	return nil
}

func (uv *userValidator) requireEmail(user *User) error {
	if user.Email == "" {
		return ErrEmailRequired
	}
	return nil
}

func (uv *userValidator) emailFormat(user *User) error {
	if user.Email == "" {
		return nil
	}
	if !uv.emailRegex.MatchString(user.Email) {
		return ErrEmailInvalid
	}
	return nil
}

func (uv *userValidator) emaiIslAvail(user *User) error {
	existing, err := uv.ByEmail(user.Email)
	if err == ErrNotFound {
		// Email address is not taken
		return nil
	}
	if err != nil {
		return err
	}

	// We found a user w/ this email address...
	// If the found user has the same ID as this user, it is an update and this
	// is the same user.
	if user.Email != existing.Email {
		return ErrEmailTaken
	}
	return nil
}

func (uv *userValidator) normalizeName(user *User) error {
	user.Name = strings.ToLower(user.Name)
	user.Name = strings.TrimSpace(user.Name)
	return nil
}

func (uv *userValidator) requireName(user *User) error {
	if user.Name == "" {
		return ErrNameRequired
	}
	return nil
}

func (uv *userValidator) nameFormat(user *User) error {
	if user.Name == "" {
		return nil
	}
	if !uv.emailRegex.MatchString(user.Name) {
		return ErrNameInvalid
	}
	return nil
}
