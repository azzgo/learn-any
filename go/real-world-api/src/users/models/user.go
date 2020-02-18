package models

import (
	"real-world-api/src/common"
	"real-world-api/src/db"

	"github.com/jinzhu/gorm"
)

// UserModel godoc
type UserModel struct {
	gorm.Model
	Email    string `gorm:"UNIQUE_INDEX;NOT NULL"`
	Username string `gorm:"UNIQUE_INDEX;NOT NULL"`
	Password string `gorm:"NOT NULL"`
	Bio      string
	Image    string
	Following bool `gorm:"DEFAULT:0"`
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
	user := new(UserModel)
	err := db.Where("email=?", email).First(user).Error
	return user, err
}

// GetUserByID godoc
func GetUserByID(id uint) (*UserModel, error) {
	db := db.GetDB()
	user := new(UserModel)
	err := db.Where("id=?", id).First(user).Error
	return user, err
}

// GetUserByUsername godoc
func GetUserByUsername(username string) (*UserModel, error)  {
	db := db.GetDB()
	user := new(UserModel)
	err := db.Where("username=?", username).First(user).Error
	return user, err
}