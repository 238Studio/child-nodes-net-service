package httpNetService

import (
	"bytes"
	_const "github.com/UniversalRobotDriveTeam/child-nodes-assist/const"
	"github.com/UniversalRobotDriveTeam/child-nodes-assist/util"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

// 发送请求
// 传入参数：httpAPP、请求体、method
// 返回参数：map[sting]interface{},错误信息
func send(h HttpAPP, req interface{}, method string) (map[string]interface{}, error) {
	//获取httpAPP的url和params
	url := h.Url
	Header := h.Params

	//json格式化
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	marshalReq, err := json.Marshal(req)
	if err != nil {
		//TODO:约定的错误处理
		return nil, util.NewError(_const.TrivialException, _const.JsonMarshalError, false, err)
	}

	//构造请求
	request, err := http.NewRequest(method, url, bytes.NewBuffer(marshalReq))
	if err != nil {
		return nil, util.NewError(_const.TrivialException, _const.HttpNewRequestError, false, err)
	}

	//设置请求头
	for k, v := range Header {
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
		//TODO:约定的错误处理
		return nil, err
	}

	return result, nil
}
