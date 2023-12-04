package Result

// MyError 定义自定义错误类型
type MyError struct {
	Code    int
	Message string
}

// 实现 error 接口的 Error 方法
func (e *MyError) Error() string {
	return e.Message
}

// ConvertToMyError 尝试将错误转换为 MyError 对象
func ConvertToMyError(err error) *MyError {
	if myErr, ok := err.(*MyError); ok {
		return myErr
	} else {
		return &MyError{Code: 10005, Message: err.Error()}
	}
}
