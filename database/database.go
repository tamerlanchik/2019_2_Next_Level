package database

import "errors"

type User struct {
	Name       string
	Sirname    string
	MiddleName string
	Email      string
	Password   string `json:"-"`
}

var sessionBook map[string]string

var usersList map[string]User

func GetUserEmailBySession(sessionId string) (string, error) {
	email, err := sessionBook[sessionId]
	if !err {
		return email, errors.New("No element")
	}
	return email, nil
}

func RegisterNewSession(sessionId, email string) {
	sessionBook[sessionId] = email
}

func GetUserByEmail(email string) (User, error) {
	user, err := usersList[email]
	if !err {
		return User{}, errors.New("No such a user")
	}
	return user, nil
}

func Init() {
	sessionBook = map[string]string{
		"12345": "a@mail.ru",
	}
	usersList = map[string]User{
		"a@mail.ru": User{
			Name:       "Ian",
			Sirname:    "Ivanov",
			MiddleName: "tamerlanchik",
			Email:      "a@mail.ru",
			Password:   "pass",
		},
	}
}
