package services

import (
	"github.com/mikezzb/steam-trading-server/db"
	"github.com/mikezzb/steam-trading-shared/database/model"
)

type User struct {
	Username string
	Password string

	User *model.User
}

func (u *User) CheckAuth() (bool, error) {
	userRepo := db.Repos.GetUserRepository()

	user, err := userRepo.GetUser(u.Username, u.Password)

	if err != nil || user == nil {
		return false, err
	}

	u.User = user
	return true, nil
}
