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

// New godoc
func New() *UserModel {
	var user = new(UserModel)
	return user
}

// TableName godoc
func (UserModel) TableName() string {
  return "users"
}

// SetPassword godoc
func (user *UserModel) SetPassword(password string) {
	user.Password = common.Hash(password)
}

// CreateUser godoc
func CreateUser(userModel *UserModel) error {
	db := db.GetDB()
	err := db.Create(userModel).Error
	return err
}

// GetUserByEmail godoc
func GetUserByEmail(email string) (*UserModel, error) {
	db := db.GetDB()
	user := UserModel{}
	err := db.Where("email=?", email).First(&user).Error
	return &user, err
}
