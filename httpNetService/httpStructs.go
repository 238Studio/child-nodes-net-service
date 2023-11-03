package httpNetService

// HttpAPP http服务应用
type HttpAPP struct {
	Url    string            //http调用路径
	Params map[string]string //http调用参数
}
