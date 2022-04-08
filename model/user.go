package model

import (
	"gin_cli/global"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	global.BaseModel
	UUID     uuid.UUID `json:"uuid" gorm:"comment:用户uuid"`                                                        // 用户UUID
	Username string    `json:"username" gorm:"用户登录名"`                                                             // 用户名
	Password string    `json:"-" gorm:"登录密码"`                                                                     // 用户登录密码
	NickName string    `json:"nickname" gorm:"comment:用户名称;default:系统用户"`                                         // 用户昵称
	Avatar   string    `json:"avatar" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	Phone    string    `json:"phone"  gorm:"comment:用户手机号"`                                                       // 用户手机号
	Email    string    `json:"email"  gorm:"comment:用户邮箱"`                                                        // 用户邮箱
	RoleID   string    `json:"roleId" gorm:"comment:角色id;default:666;"`                                           // 用户角色id
}

// 加密
func (user *User) SetPssword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// 验证密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
