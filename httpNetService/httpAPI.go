package httpNetService

// Get 发送GET请求
// 传入参数：请求体
// 返回参数：请求结果，错误信息
func (app *HttpAPP) Get(req interface{}) (map[string]interface{}, error) {
	return send(*app, req, "GET")
}

// Post 发送POST请求
// 传入参数：请求体
// 返回参数：请求结果，错误信息
func (app *HttpAPP) Post(req interface{}) (map[string]interface{}, error) {
	return send(*app, req, "POST")
}

// Put 发送PUT请求
// 传入参数：请求体
// 返回参数：请求结果，错误信息
func (app *HttpAPP) Put(req interface{}) (map[string]interface{}, error) {
	return send(*app, req, "PUT")
}

// Delete 发送DELETE请求
// 传入参数：请求体
// 返回参数：请求结果，错误信息
func (app *HttpAPP) Delete(req interface{}) (map[string]interface{}, error) {
	return send(*app, req, "DELETE")
}
