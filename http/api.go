// Package http
// 这个包是暴露给控制结构的API层
// 控制上位机的HTTP网络传输过程
package http

import "net/http"

// Get 发送GET请求
// 传入参数：url、请求头、请求体
// 返回参数：请求结果，响应头，错误信息
func Get(url string, header map[string]string, body interface{}) (map[string]interface{}, http.Header, error) {
	return send(url, header, body, "GET")
}

// Post 发送POST请求
// 传入参数：url、请求头、请求体
// 返回参数：请求结果，响应体，错误信息
func Post(url string, header map[string]string, body interface{}) (map[string]interface{}, http.Header, error) {
	return send(url, header, body, "POST")
}

// Put 发送PUT请求
// 传入参数：url、请求头、请求体
// 返回参数：请求结果，响应体，错误信息
func Put(url string, header map[string]string, body interface{}) (map[string]interface{}, http.Header, error) {
	return send(url, header, body, "PUT")
}

// Delete 发送DELETE请求
// 传入参数：url、请求头、请求体
// 返回参数：请求结果，响应体，错误信息
func Delete(url string, header map[string]string, body interface{}) (map[string]interface{}, http.Header, error) {
	return send(url, header, body, "DELETE")
}

// Download 下载文件
// 传入参数：url、请求头、请求体
// 返回参数：请求结果([]byte)，响应体，错误信息
func Download(url string, header map[string]string, body interface{}) ([]byte, http.Header, error) {
	return sendAndReturnByte(url, header, body)
}
