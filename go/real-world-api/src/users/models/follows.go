package models

import (
	"real-world-api/src/db"

	"github.com/jinzhu/gorm"
)

// FollowModels godoc
type FollowModels struct {
	gorm.Model
	UserID         uint
	FollowedUserID uint
}

// TableName godoc
func (FollowModels) TableName() string {
	return "follow"
}

// GetTargetFollowedIDs godoc
func GetTargetFollowedIDs(username string) ([]uint, error) {
	db := db.GetDB()
	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
		}
	}()
	userModel, err := GetUserByUsername(username)
	var followModels = make([]*FollowModels, 0)
	err = db.Model(userModel).Related(&followModels, "user_id").Error

	ids := make([]uint, 0)

	for _, followModel := range(followModels) {
		ids = append(ids, followModel.UserID)
	}

	return ids, err
}
