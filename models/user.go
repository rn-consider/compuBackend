package models

import (
	"database/sql"
	"github.com/rn-consider/compuBackend/dao"
	"time"
)

type User struct {
	ID           uint `gorm:"primaryKey"`
	Name         string
	Email        *string
	Age          uint8
	Birthday     time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

/*
   User的增删改查
*/
// CreateUser 创建User字段
func CreateUser(user *User) (err error) {
	if err = dao.DB.Create(&user).Error; err != nil {
		return err
	}
	return
}

// DeleteUser 根据主键删除User字段
func DeleteUser(ID int) (err error) {
	err = dao.DB.Where("ID =?", ID).Delete(&User{}).Error
	return
}

func UpdatedAtUser(user *User) (err error) {
	err = dao.DB.Save(user).Error
	return
}

func GetAllUser() (userList []*User, err error) {
	if err = dao.DB.Find(&userList).Error; err != nil {
		return nil, err
	}
	return
}

func GetAUser(id int) (user *User, err error) {
	if err = dao.DB.Where("id=?", id).First(user).Error; err != nil {
		return nil, err
	}
	return
}
