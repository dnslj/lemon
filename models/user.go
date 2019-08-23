package models

type UserModel struct {
	BaseModel
	Mobile   string `json:"mobile" gorm:"column:mobile;not null" binding:"required" validate:"min=11,max=11" form:"mobile"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128" form:"password"`
	NickName string `json:"nickname" gorm:"column:nickname;not null"`
}

func (user *UserModel) TableName() string {
	return "t_users"
}

func (user *UserModel) Create() error {
	return DB.Local.Create(&user).Error
}

func (user *UserModel) Update() error {
	return DB.Local.Save(&user).Error
}

func GetUserById(UserId uint64) (*UserModel, error) {
	user := &UserModel{}
	d := DB.Local.Where("id = ?", UserId).First(&user)
	return user, d.Error
}

func GetUserByMobile(mobile string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Local.Where("mobile = ?", mobile).First(&u)
	return u, d.Error
}
