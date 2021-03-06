package service

import (
	"github.com/gold-kou/go-housework/app/server/middleware"
	"github.com/gold-kou/go-housework/app/server/repository"
	"github.com/jinzhu/gorm"
)

// DeleteUserServiceInterface is a service interface of deleteUser
type DeleteUserServiceInterface interface {
	Execute(*middleware.Auth) error
}

// DeleteUser struct
type DeleteUser struct {
	tx       *gorm.DB
	userRepo repository.UserRepositoryInterface
}

// NewDeleteUser constructor
func NewDeleteUser(tx *gorm.DB, userRepo repository.UserRepositoryInterface) *DeleteUser {
	return &DeleteUser{
		tx:       tx,
		userRepo: userRepo,
	}
}

// Execute delete user
func (u *DeleteUser) Execute(auth *middleware.Auth) error {
	// data
	userName := auth.UserName

	// delete user
	err := u.userRepo.DeleteUserWhereUsername(userName)
	if err != nil {
		return err
	}
	return nil
}
