package types

// APIResponse api response
type APIResponse struct {
	Msg     interface{} `json:"msg"`
	Success bool        `json:"success"`
}

// APIErrResponse api response
type APIErrResponse struct {
	Msg     interface{} `json:"msg"`
	Success bool        `json:"success"`
	Err     interface{} `json:"err"`
}
