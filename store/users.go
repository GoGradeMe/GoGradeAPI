package store

import (
	"errors"

	r "github.com/dancannon/gorethink"
	m "github.com/gogrademe/apiserver/model"
)

var (

	//ErrUserAlreadyExists err for duplicate user
	ErrUserAlreadyExists = errors.New("User with email already exists.")
	//ErrUserForPersonExists err for duplicate person
	ErrUserForPersonExists = errors.New("User for person already exists.")
)

// UserStore used to interact with db users
type UserStore struct {
	DefaultStore
}

// NewUserStore returns a UserStore.
func NewUserStore() UserStore {
	return UserStore{DefaultStore: NewDefaultStore("users")}
}

func userExist(email string) bool {
	row, _ := r.Table("users").Filter(r.Row.Field("email").Eq(email)).Run(sess)

	return !row.IsNil()
}

func userForPersonExist(personID string) bool {
	row, _ := r.Table("users").Filter(r.Row.Field("personId").Eq(personID)).Run(sess)

	return !row.IsNil()
}

// Store saves a user into the db
func (us *UserStore) Store(u *m.User) error {
	if userExist(u.Email) {
		return ErrUserAlreadyExists
	}
	if userForPersonExist(u.PersonID) {
		return ErrUserForPersonExists
	}
	res, err := r.Table("users").Insert(u).RunWrite(sess)
	if err != nil {
		return err
	}
	if u.ID == "" && len(res.GeneratedKeys) == 1 {
		u.ID = res.GeneratedKeys[0]
	}

	return nil
}

// FindByEmail finds a single user that matches an email.
func (us *UserStore) FindByEmail(email string) (m.User, error) {
	var u m.User

	res, err := r.Table("users").Filter(r.Row.Field("email").Eq(email)).Run(sess)
	if err != nil {
		return u, err
	}

	err = res.One(&u)
	return u, nil

}

// GetUserByID get a user by a ID.
func GetUserByID(id string) (m.User, error) {
	u := m.User{}

	res, err := r.Table("users").Get(id).Run(sess)
	if err != nil {
		return u, err
	}

	err = res.All(&u)
	if err != nil {
		return u, err
	}

	return u, nil
}
