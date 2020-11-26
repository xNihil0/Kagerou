package database

import (
	"gorm.io/gorm"
	"mtdn.io/Kagerou/internal"
)

type User struct {
	gorm.Model
	TelegramID       int32  `json:"telegram_id" gorm:"uniqueIndex"`
	U2ID             int32  `json:"u2_id" gorm:"column:u2_id"`
	VerificationCode string `json:"verification_code"`
	Verified         bool   `json:"verified"`
}

func (inputUser *InputUser) CreateUser() error {
	user := &User{
		Model:            gorm.Model{},
		TelegramID:       inputUser.TelegramID,
		U2ID:             inputUser.U2ID,
		VerificationCode: internal.GenerateVerificationCode(),
		Verified:         false,
	}
	if err := UserDB.Create(user).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func (user *InputUser) GetUser() (*User, error) {
	var res User
	if err := UserDB.Where(&User{TelegramID: user.TelegramID}).First(&res).Error; err != nil {
		return nil, err
	} else {
		return &res, nil
	}
}

func (user *User) UpdateUser() error {
	var u User
	if err := UserDB.First(&u, user.ID).Error; err != nil {
		return err
	} else {
		u.TelegramID = user.TelegramID
		u.U2ID = user.U2ID
		u.VerificationCode = user.VerificationCode
		u.Verified = user.Verified
	}
	if err := UserDB.Save(&u).Error; err != nil {
		return err
	}
	return nil
}

func (user *User) RemoveUser() error {
	if err := UserDB.Delete(user).Error; err != nil {
		return err
	} else {
		return nil
	}
}
