package interfaces

import "github.com/techagentng/medium/cmd/model"

type User interface {
	GetUserByEmail(email string) (bool, model.User)
	SaveUser(user *model.User) (string, error)
	GetUserById(id string) (model.User, error)
}
