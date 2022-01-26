package models

import (
	"errors"
	"fmt"
	"github.com/EDDYCJY/go-gin-example/pkg/request"
	"github.com/jinzhu/gorm"
)

type User struct {
	Model
	Username    string `gorm:"comment:'用户名';size:50;not null;index;default:''" json:"username"`
	Password    string `gorm:"comment:'密码';size:256;not null;default:''" json:"password"`
	RoleIdS     string `gorm:"comment:'用户角色id';not null;default:0" json:"role_ids"`
	Desc        string `gorm:"comment:'描述信息';size:256;not null;default:''" json:"desc"`
	IsActive    bool   `gorm:"comment:'是否激活';not null;default:true" json:"is_active"`
	Virtual     bool   `gorm:"comment:'是否为虚拟用户';not null;default:false" json:"virtual"`        // 是否为虚拟用户默认为 false
	ImageUrl    string `gorm:"comment:'头像链接地址';size:256;not null;default:''" json:"image_url"` // 头像链接地址
	Phone       string `gorm:"comment:'手机号';not null;index;default:0" json:"phone"`
	IsCommunity string `gorm:"comment:'是否为社区用户';not null;default:'no_community'" json:"is_community"` //默认为 非社区成员
	Email       string `gorm:"comment:'邮箱';not null;default:''" json:"email"`
	StaffCode   string `gorm:"comment:'工号';not null;default:''" json:"staff_code"`
}

// BeforeUpdate update hook
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	fmt.Println("运行了 User BeforeUpdate")
	// TODO 如何获取请求中传递的 user_id
	if u.Id == 1 {
		return errors.New("admin user not allowed to update")
	}
	return
}

func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Println("运行了 User BeforeDelete")
	if u.Id == 1 {
		return errors.New("admin user not allowed to delete")
	}
	return
}

//generateInitPassword 生成虚拟用户的初始化密码
func (u User) generateInitPassword() string {
	return "abc123"
}

//type User struct {
//	Username string `gorm:"comment:'用户名';size:50;not null;index;default:''" json:"username"`
//	Password string `gorm:"comment:'密码';size:256;not null;default:''" json:"password"`
//}

// CheckAuth token校验
func CheckAuth(username, password string) (bool, error) {
	//var user User
	//db.Find(&user, "username = ? AND is_delete = ?", username, 0)
	//record := db.First(&user)
	//fmt.Print(record)

	// 新增一个用户
	//user := User{Username: "admin1",Password:"admin1"}

	//result := db.Create(&user) // 通过数据的指针来创建
	//fmt.Println(result.RowsAffected)
	//fmt.Println(result.RowsAffected)
	//var user User
	//result := db.First(&user)
	//row := result.RowsAffected
	//fmt.Println("row count",row)
	//err := db.Select(&user, "select id,username,password from user where username=?","admin1")
	//if err != nil {
	//	fmt.Println("exec failed, ", err)
	//}

	//if err != nil && err != gorm.ErrRecordNotFound {
	//	return false,err
	//}
	var user User
	//err := db.Select("id").Where(User{Username: username, Password: password}).First(&user).Error
	err := db.Where("username = ? and password = ?", username, password).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.Id > 0 {
		return true, nil
	}

	return false, nil
}

// GetUser 获取用户
func GetUser(req *request.UserRequestStruct) (interface{}, error) {
	var err error
	// 这里务必定义成切片,切片可以自动扩容
	var user []User

	// db.Debug 可以查看生成的原生sql语句
	db.Find(&user)
	// 查询列表
	// 这里不能 return result, result 为 *gorm.DB 对象
	//return result, err
	return user, err
}

// CreateUser 创建用户
func CreateUser(username, password string) (uint, error) {
	var user User
	err := db.Find(&user).Where("username = ?", username).First(&user).Error
	if err == nil && err != gorm.ErrRecordNotFound {
		//panic("用户名已存在")
		//TODO 增加全局异常捕获
		fmt.Println("用户名已存在")
		return 0, nil
	}

	createUser := User{Username: username, Password: password}
	// 新增一个用户
	db.Create(&createUser)

	return createUser.Id, nil
}

// UpdateUser 修改用户名
func UpdateUser(userId int, username string) (bool, error) {

	result := db.Model(&User{}).Where("id = ?", userId).Update("username", username)
	if result.Error != nil {
		panic("更新用户失败")
	} else {
		return true, nil
	}
}

// DeleteUSer 删除指定用户
func DeleteUser(userId int) (bool, error) {
	var user User
	result := db.Where("id = ?", userId).Delete(&user)
	if result.Error != nil {
		panic("删除用户成功")
	} else {
		return true, nil
	}

}
