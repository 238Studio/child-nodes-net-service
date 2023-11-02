package service

import (
	"bytes"
	"net/http"

	_const "github.com/UniversalRobotDriveTeam/child-nodes-assist/const"
	"github.com/UniversalRobotDriveTeam/child-nodes-assist/util"
	jsoniter "github.com/json-iterator/go"
)

// Send 发送请求
// 传入参数：url、请求头、请求体、method
// 返回参数：map[sting]interface{}（返回结果，interface{}做接口断言）,错误信息
func Send(url string, params map[string]string, req interface{}, method string) (map[string]interface{}, error) {
	//json格式化
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	marshalReq, err := json.Marshal(req)
	if err != nil {
		return nil, util.NewError(_const.TrivialException, _const.JsonMarshalError, false, err)
	}

	//构造请求
	request, err := http.NewRequest(method, url, bytes.NewBuffer(marshalReq))
	if err != nil {
		return nil, util.NewError(_const.TrivialException, _const.HttpNewRequestError, false, err)
	}

	//设置请求头
	for k, v := range params {
		request.Header.Set(k, v)
	}

	//发送请求
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, util.NewError(_const.CommonException, _const.HttpDoError, false, err)
	}

	//关闭请求
	defer response.Body.Close()

	//读取返回值
	var result map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, util.NewError(_const.CommonException, _const.JsonUnmarshalError, false, err)
	}

	return result, nil
}
