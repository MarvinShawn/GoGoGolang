package model

import (
	"../pkg/auth"
	"../pkg/constvar"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type UserModel struct {
	BaseModel
	Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}

func (c *UserModel) TableName() string {

	return "tb_users"
}

// create creates a new user account
func (u *UserModel) Create() error {

	return DB.Self.Create(&u).Error

}

// 通过id 删除 user
func DeleteUser(id uint64) error {
	user := UserModel{}
	user.BaseModel.Id = id
	return DB.Self.Delete(&user).Error
}

// 更新
func (u *UserModel) Update() error {

	return DB.Self.Save(u).Error

}

func GetUserById(id uint64, user *UserModel) error {

	d := DB.Self.First(&user, id)
	return d.Error

}

// 查询一个user
func GetUser(username string) (*UserModel, error) {

	u := &UserModel{}
	d := DB.Self.Where("username = ?", username).First(&u)
	return u, d.Error

}

func ListUser(username string, offset, limit int) ([]*UserModel, uint64, error) {

	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	users := make([]*UserModel, 0)
	var count uint64

	where := fmt.Sprintf("username like '%%%s%%'", username)

	if err := DB.Self.Model(&UserModel{}).Where(where).Count(&count).Error; err != nil {

		return users, count, err

	}

	if err := DB.Self.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {

		return users, count, err

	}

	return users, count, nil

}

func (u *UserModel) Compare(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)
	return
}

func (u *UserModel) Encrypt() (err error) {

	u.Password, err = auth.Encrypt(u.Password)
	return
}

func (u *UserModel) Validate() error {

	validate := validator.New()

	return validate.Struct(u)

}
