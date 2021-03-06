package mailpicker

import "2019_2_Next_Level/internal/model"

type Repository interface {
	UserExists(login string) bool
	AddEmail(*model.Email) error
}
