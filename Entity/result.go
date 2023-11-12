package Entity

type APIResponse struct {
	StatusCode int         `json:"status_code"`
	Success    bool        `json:"success"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Error      string      `json:"error,omitempty"`
}

func (r *APIResponse) OK(data interface{}) *APIResponse {
	r.Success = true
	r.Data = data
	r.StatusCode = 200
	return r
}

func (r *APIResponse) Err(errorMessage string, statusCode int) *APIResponse {
	r.Success = false
	r.Error = errorMessage
	r.StatusCode = statusCode
	return r
}
