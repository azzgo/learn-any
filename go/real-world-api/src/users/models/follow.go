package models

import (
	"real-world-api/src/db"

	"github.com/jinzhu/gorm"
)

// FollowModel godoc
type FollowModel struct {
	gorm.Model
	UserID         uint `gorm:"not null"`
	FollowedUserID uint `gorm:"not null"`
}

// TableName godoc
func (FollowModel) TableName() string {
	return "follow"
}

// GetTargetFollowingIDs godoc
func GetTargetFollowingIDs(username string, userID uint) ([]uint, error) {
	db := db.GetDB()
	ids := make([]uint, 0)
	var userModel = new(UserModel)
	var err error

	if username != "" {
		userModel.Username = username
	} else {
		userModel.ID = userID
	}

	var followModels = make([]*FollowModel, 0)
	err = db.Find(userModel).Related(&followModels, "followed_user_id").Error
	if err != nil {
		return ids, err
	}

	for _, followModel := range(followModels) {
		ids = append(ids, followModel.FollowedUserID)
	}

	return ids, nil
}

// GetTargetFollowedIDs godoc
func GetTargetFollowedIDs(username string, userID uint) ([]uint, error) {
	db := db.GetDB()
	ids := make([]uint, 0)
	var err error
	var userModel = new(UserModel)

	if username != "" {
		userModel.Username = username
	} else {
		userModel.ID = userID
	}

	var followModels = make([]*FollowModel, 0)
	err = db.Find(userModel).Related(&followModels, "user_id").Error
	if err != nil {
		return ids, err
	}

	for _, followModel := range(followModels) {
		ids = append(ids, followModel.FollowedUserID)
	}

	return ids, nil
}

// FollowUser godoc
func FollowUser(currentUser *UserModel, username string, ) error {
	db := db.GetDB()
	var err error
	userModel, err := GetUserByUsername(username)
	if err != nil {
		return err
	}

	// Make Sure Record not in database
	var count int = 0;
	if db.Table("follow").Where("user_id=? AND followed_user_id=?", userModel.ID, currentUser.ID).Count(&count); count == 0 {
		return nil
	}

	followModel := new(FollowModel)
	followModel.UserID = userModel.ID 
	followModel.FollowedUserID = currentUser.ID

	err = db.Save(followModel).Error

	if err != nil {
		return err
	}

	return nil
}

// UnFollowUser godoc
func UnFollowUser(CurrentUser *UserModel, username string, ) error {
	db := db.GetDB()
	var err error
	userModel, err := GetUserByUsername(username)
	if err != nil {
		return err
	}

	followModel := new(FollowModel)
	followModel.UserID = userModel.ID 
	followModel.FollowedUserID = CurrentUser.ID

	err = db.Delete(followModel).Error

	if err != nil {
		return err
	}

	return nil
}
