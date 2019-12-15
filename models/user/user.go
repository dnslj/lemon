package user

import (
	"lemon/models"
)

type UserModel struct {
	models.BaseModel
	Mobile   string `json:"mobile" gorm:"column:mobile;not null" binding:"required" validate:"min=1,max=11" form:"mobile"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128" form:"password"`
	NickName string `json:"nickname" gorm:"column:nickname;not null"`
}

type UserList struct {
	UserList UserModel `json:"user_list"`
}

func (user *UserModel) TableName() string {
	return "t_users"
}

func (user *UserModel) Create() error {
	return models.DB.Default.Create(&user).Error
}

func (user *UserModel) Update() error {
	return models.DB.Default.Save(&user).Error
}

func GetUserById(UserId uint64) (*UserModel, error) {
	user := &UserModel{}
	d := models.DB.Default.Where("id = ?", UserId).First(&user)
	return user, d.Error
}

func GetUserByMobile(mobile string) (*UserModel, error) {
	u := &UserModel{}
	d := models.DB.Default.Where("mobile = ?", mobile).First(&u)
	return u, d.Error
}

func GetUserList() ([]*UserModel, error) {
	var users []*UserModel
	d := models.DB.Default.Find(&users)
	return users, d.Error
}

func UpdateUserById(UserId uint64, data interface{}) error {
	result := models.DB.Default.Model(&UserModel{}).Where("id=?", UserId).Updates(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
