package types

// APIResponse api response
type APIResponse struct {
	Msg     interface{} `json:"msg"`
	Success bool        `json:"success"`
}
