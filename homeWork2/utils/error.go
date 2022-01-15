package utils
// Error 自定义错误接口
type Error interface {
	error
	Status() int
}

// StatusError 声明错误状态结构体
type StatusError struct {
	Code int
	Err  error
}
// 返回错误状态码
func (se StatusError) Error() string {
	return se.Err.Error()
}

// Status 实现错误接口方法体
func (se StatusError) Status() int {
	return se.Code
}
