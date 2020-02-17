package models

import (
	"real-world-api/src/common"
	"real-world-api/src/db"

	"github.com/jinzhu/gorm"
)

// UserModel godoc
type UserModel struct {
	gorm.Model
	Email    string `gorm:"PRIMARY_KEY;UNIQUE_INDEX"`
	Password string
	Username string
	Bio      string
	Image    string
}

// TableName godoc
func (UserModel) TableName() string {
  return "users"
}

// SetPassword godoc
func (u *UserModel) SetPassword(password string) {
	u.Password = common.Hash(password)
}

// CreateUser godoc
func CreateUser(email string, username string, password string) (*UserModel, error) {
	user := new(UserModel)
	user.Email = email
	user.Username = username
	user.SetPassword(password)
	db := db.GetDB()
	err := db.Create(user).Error
	return user, err
}

// GetUserByEmail godoc
func GetUserByEmail(email string) (*UserModel, error) {
	db := db.GetDB()
	user := UserModel{}
	err := db.Where("email=?", email).First(&user).Error
	return &user, err
}
