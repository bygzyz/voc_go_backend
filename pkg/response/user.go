package response

// 用户信息响应, 字段含义见models.SysUser
type UserListResponseStruct struct {
	BaseData
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Desc        string `json:"desc"`
	IsActive    bool   `json:"is_active"`
	Virtual     bool   `json:"virtual"`
	ImageUrl    string `json:"image_url"`
	Phone       string `json:"phone"`
	IsCommunity string `json:"is_community"`
}
