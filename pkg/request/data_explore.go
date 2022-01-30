package request

// DataExploreRequestStruct 数据探索请求结构体
type DataExploreRequestStruct struct {
	Start         string `json:"start"`
	End           string `json:"end" `
	Category      string `json:"category"`
	Brand         bool   `json:"brand"`
	Model         bool   `json:"model"`
	DataType      string `json:"data_type"`
	MessageSource string `json:"message_source"`
}
