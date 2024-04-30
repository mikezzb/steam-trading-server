package services

import (
	"github.com/mikezzb/steam-trading-server/db"
	"github.com/mikezzb/steam-trading-server/util"
	"github.com/mikezzb/steam-trading-shared/database/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// beego validation form

type SignupForm struct {
	Username string `json:"username" valid:"Required; MinSize(4); MaxSize(20)"`
	Password string `json:"password" valid:"Required; MinSize(8); MaxSize(30)"`

	Email string `json:"email" valid:"Email; MaxSize(50)"`
}

type LoginForm struct {
	Email    string `json:"email" valid:"Email; MaxSize(50)"`
	Password string `json:"password" valid:"Required; MinSize(8); MaxSize(30)"`
}

type User struct {
	ID       primitive.ObjectID
	Username string
	Password string

	Email string

	// Auto set by CheckAuth
	User *model.User
}

func (u *User) GetHashedPassword() string {
	pwd, _ := util.HashPassword(u.Password)
	return pwd
}

// CheckAuth checks if the user exists and the password is correct
func (u *User) CheckAuth() (bool, error) {
	userRepo := db.Repos.GetUserRepository()

	user, err := userRepo.GetUserByEmail(u.Email)
	if err != nil || user == nil {
		return false, err
	}

	// compare password
	err = util.ComparePassword(user.Password, u.Password)
	if err != nil {
		return true, err
	}

	u.User = user
	return true, nil
}

func (u *User) toUserModel() *model.User {
	return &model.User{
		Username: u.Username,
		Password: u.GetHashedPassword(),
		Email:    u.Email,
	}
}

// CreateUser creates a user, returns the JWT token
func (u *User) CreateUser() (string, error) {
	userRepo := db.Repos.GetUserRepository()

	uid, err := userRepo.InsertUser(u.toUserModel())
	if err != nil {
		return "", err
	}

	return util.GenerateToken(uid.Hex(), "")
}

func (u *User) LoadUser() error {
	userRepo := db.Repos.GetUserRepository()

	user, err := userRepo.GetUserById(u.ID)
	if err != nil {
		return err
	}
	u.User = user
	return nil
}

// return all user data except password
func (u *User) GetViewableUser() *model.User {
	user := *u.User
	user.Password = ""
	return &user
}
