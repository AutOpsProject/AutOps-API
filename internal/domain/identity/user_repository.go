package identity

import "github.com/AutOpsProject/AutOps-API/internal/domain/common"

type UserRepository interface {
	Create(user *User) error
	Update(user *User) error
	Delete(userId common.Identifier) error

	FindById(id common.Identifier, offset int, limit int) (*User, error)
	FindAll(offset int, limit int) ([]*User, error)
	FindByUsername(username string, offset int, limit int) (*User, error)
	FindByEmail(email string, offset int, limit int) (*User, error)
}
