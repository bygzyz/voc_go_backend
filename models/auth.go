package models

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//
//// CheckAuth checks if authentication information exists
//func CheckAuth(username, password string) (bool, error) {
//	var user User
//	err := initialize.db.Select("id").Where(User{Username: username, Password: password}).First(&user).Error
//	if err != nil && err != gorm.ErrRecordNotFound {
//		return false, err
//	}
//
//	if user.ID > 0 {
//		return true, nil
//	}
//
//	return false, nil
//}
