// Package childHTTP
// 这个包是暴露给控制结构的API层
// 控制上位机的网络传输过程
package childHTTP

import "child-nodes-net-service/http/service"

// Get 发送GET请求
// 传入参数：url、请求头、请求体
// 返回参数：请求结果，错误信息
func Get(url string, header map[string]string, req interface{}) (map[string]interface{}, error) {
	return service.Send(url, header, req, "GET")
}

// Post 发送POST请求
// 传入参数：url、请求头、请求体
// 返回参数：请求结果，错误信息
func Post(url string, header map[string]string, req interface{}) (map[string]interface{}, error) {
	return service.Send(url, header, req, "POST")
}

// Put 发送PUT请求
// 传入参数：url、请求头、请求体
// 返回参数：请求结果，错误信息
func Put(url string, header map[string]string, req interface{}) (map[string]interface{}, error) {
	return service.Send(url, header, req, "PUT")
}

// Delete 发送DELETE请求
// 传入参数：url、请求头、请求体
// 返回参数：请求结果，错误信息
func Delete(url string, header map[string]string, req interface{}) (map[string]interface{}, error) {
	return service.Send(url, header, req, "DELETE")
}
