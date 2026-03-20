package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	users = []User{
		{ID: "u1", UserName: "admin", PasswordHash: "", Role: RoleAdmin},
		{ID: "u2", UserName: "artist1", PasswordHash: "", Role: RoleArtist},
		{ID: "u3", UserName: "artist2", PasswordHash: "", Role: RoleArtist},
	}
	err = errors.New("Invalid User")
)

func hashPlain(plain string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashed)
}
func Authenticate(username, password string) (*User, error) {
	for _, u := range users {
		if u.UserName == username {
			if e := bcrypt.CompareHashAndPassword(
				[]byte(u.PasswordHash), []byte(password),
			); e != nil {
				return nil, err
			}
			return &u, nil
		}
	}
	return nil, err
}
