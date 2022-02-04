package request

// 获取用户列表结构体
type UserRequestStruct struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Desc        string `json:"desc"`
	IsActive    bool   `json:"is_active"`
	Virtual     bool   `json:"virtual"`
	ImageUrl    string `json:"image_url"`
	Phone       string `json:"phone"`
	IsCommunity string `json:"is_community"`
}

// CreateUserRequestStruct 创建用户结构体
type CreateUserRequestStruct struct {
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Desc        string `json:"desc"`
	IsActive    bool   `json:"is_active"`
	Virtual     bool   `json:"virtual"`
	ImageUrl    string `json:"image_url"`
	Phone       string `json:"phone"`
	IsCommunity string `json:"is_community"`
}

// 更新用户结构体
type UpdateUserRequestStruct struct {
	UserId      int    `json:"user_id" validate:"required"`
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Desc        string `json:"desc"`
	IsActive    bool   `json:"is_active"`
	Virtual     bool   `json:"virtual"`
	ImageUrl    string `json:"image_url"`
	Phone       string `json:"phone"`
	IsCommunity string `json:"is_community"`
}

// 删除用户结构体
type DeleteUserRequestStruct struct {
	UserId int `json:"user_id" validate:"required"`
}
