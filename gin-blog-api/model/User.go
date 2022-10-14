package model

import (
	"encoding/base64"
	"gin-blog/utils/errmsg"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20); not null" json:"username" validate:"required,min=6,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(20); not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色"`
}

// CheckUser 查询用户是否存在
func CheckUser(name string) (code int) {
	var users User
	db.Select("id").Where("username = ?", name).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED // 1001
	}
	return errmsg.SUCCESS
}

// CheckUpUser 更新查询
func CheckUpUser(id int, name string) (code int) {
	var user User
	db.Select("id, username").Where("username = ?", name).First(&user)
	if user.ID == uint(id) {
		return errmsg.SUCCESS
	}
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //1001
	}
	return errmsg.ERROR
}

// CreateUser 新增用户
func CreateUser(data *User) int {
	//data.Password = ScryptPw(data.Password)
	// 如果想自己定义加密方法 不按照gorm来写 就直接调用即可
	//data.Before()
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200
}

// GetUserInfo 获取用户信息
func GetUserInfo(id int) (User, int) {
	var user User
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, errmsg.ERROR_USER_NOT_EXIST
	}
	return user, errmsg.SUCCESS
}

// GetUsers 获取用户列表
func GetUsers(pageSize int, pageNum int) ([]User, int64) {
	var users []User
	var total int64
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return users, total
}

// EditUser 编辑用户
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = db.Model(&user).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteUser 删除用户
func DeleteUser(id int) int {
	var user User
	err = db.Where("id = ? ", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// BeforeSave 密码加密
//钩子函数调用 实现密码加密效果
func (u *User) BeforeSave(_ *gorm.DB) (err error) {
	u.Password = ScryptPw(u.Password)
	return
}

//func (u *User)Before(){
//	u.Password = ScryptPw(u.Password)
//	return
//}

func ScryptPw(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 6, 66, 22, 222, 111}
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	FinalPw := base64.StdEncoding.EncodeToString(HashPw)
	return FinalPw
}

// CheckLogin 登录验证
func CheckLogin(username string, password string) int {
	var user User
	db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if user.Password != ScryptPw(password) {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 0 {
		return errmsg.ERROR_USER_NOT_RIGHT
	}
	return errmsg.SUCCESS
}
