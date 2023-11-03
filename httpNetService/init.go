package httpNetService

// InitHttpAPP 初始化HTTP服务应用
// 传入参数：http地址, http头文件参数
// 返回参数：http服务应用
func InitHttpAPP(url string, params map[string]string) *HttpAPP {
	h := HttpAPP{Url: url, Params: params}
	return &h
}
