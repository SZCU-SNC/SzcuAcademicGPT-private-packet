package Result

type APIResponse struct {
	StatusCode int         `json:"status_code"`
	Success    bool        `json:"success"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Error      string      `json:"error,omitempty"`
}

const (
	ErrExpiredIdentity   = 10001 // 身份过期
	ErrFrequentCall      = 10002 // 调用太频繁
	ErrExecutionTimeout  = 10003 // 执行等待超时
	ErrPermissionDenied  = 10004 // 权限错误
	ErrCodeError         = 10005 // 代码错误
	ErrInvalidParameters = 10006 // 参数错误
)

func (r *APIResponse) OK(data interface{}) *APIResponse {
	r.Success = true
	r.Data = data
	r.StatusCode = 200
	return r
}

func (r *APIResponse) Err(errorMessage string, statusCode int) *APIResponse {
	r.Success = false
	r.Error = getErrorMessage(statusCode)
	r.StatusCode = statusCode
	r.Message = errorMessage
	return r
}

func getErrorMessage(statusCode int) string {
	switch statusCode {
	case ErrExpiredIdentity:
		return "Err Expired Identity"
	case ErrFrequentCall:
		return "Err Frequent Call"
	case ErrExecutionTimeout:
		return "Err Execution Timeout"
	case ErrPermissionDenied:
		return "Err Permission Denied"
	case ErrCodeError:
		return "Err Code Error"
	case ErrInvalidParameters:
		return "Err Invalid Parameters"
	default:
		return ""
	}
}
