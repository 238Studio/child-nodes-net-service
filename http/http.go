package http

import (
	"bytes"
	"errors"
	"io"
	"net/http"

	"github.com/238Studio/child-nodes-error-manager/errpack"
	jsoniter "github.com/json-iterator/go"
)

// Send 发送请求
// 传入参数：url、请求头、请求体、method
// 返回参数：map[sting]interface{}（返回结果，interface{}做接口断言）,响应头,错误信息
func send(url string, params map[string]string, req interface{}, method string) (result map[string]interface{}, header http.Header, err error) {
	//捕获panic
	defer func() {
		if er := recover(); er != nil {
			//panic错误，定级为fatal
			//返回值赋值
			err = errpack.NewError(errpack.FatalException, errpack.Network, errors.New(er.(string)))
			result = nil
		}
	}()

	//json格式化
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	marshalReq, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errpack.NewError(errpack.TrivialException, errpack.Network, err)
	}

	//构造请求
	request, err := http.NewRequest(method, url, bytes.NewBuffer(marshalReq))
	if err != nil {
		return nil, nil, errpack.NewError(errpack.TrivialException, errpack.Network, err)
	}

	//设置请求头
	for k, v := range params {
		request.Header.Set(k, v)
	}

	//FIXME:设置超时
	//发送请求
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, nil, errpack.NewError(errpack.CommonException, errpack.Network, err)
	}

	//关闭请求
	defer response.Body.Close()

	//获取响应头
	header = response.Header

	//检查响应体是否为空值
	if response.Body == nil {
		return nil, header, nil
	}

	//读取返回值
	result = make(map[string]interface{})
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, nil, errpack.NewError(errpack.CommonException, errpack.Network, err)
	}

	return result, header, nil
}

// SendAndReturnByte 发送请求并直接返回[]byte。只使用get方法
// 传入参数：url、请求头、请求体
// 返回参数：[]byte（返回结果）,响应体,错误信息
func sendAndReturnByte(url string, params map[string]string, req interface{}) (result []byte, header http.Header, err error) {
	//捕获panic
	defer func() {
		if er := recover(); er != nil {
			//panic错误，定级为fatal
			//返回值赋值
			err = errpack.NewError(errpack.FatalException, errpack.Network, errors.New(er.(string)))
			result = nil
		}
	}()

	//json格式化
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	marshalReq, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errpack.NewError(errpack.TrivialException, errpack.Network, err)
	}

	//构造请求
	request, err := http.NewRequest("GET", url, bytes.NewBuffer(marshalReq))
	if err != nil {
		return nil, nil, errpack.NewError(errpack.TrivialException, errpack.Network, err)
	}

	//设置请求头
	for k, v := range params {
		request.Header.Set(k, v)
	}

	//FIXME:设置超时
	//发送请求
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, nil, errpack.NewError(errpack.CommonException, errpack.Network, err)
	}

	//关闭请求
	defer response.Body.Close()

	//获取响应头
	header = response.Header

	//获取[]byte
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, nil, errpack.NewError(errpack.CommonException, errpack.Network, err)
	}

	return body, header, nil
}
