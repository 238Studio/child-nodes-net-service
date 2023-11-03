package httpNetService_test

import (
	"child-nodes-http-service/httpNetService"
	"testing"
)

func TestHttpGet(t *testing.T) {
	app := httpNetService.InitHttpAPP("http://127.0.0.1:8080/get", nil)
	var test struct {
		UserName string
		PassWord string
	}
	test.UserName = "test"
	test.PassWord = "test"

	resp, err := app.Get(test)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(resp)
}

func TestHttpPost(t *testing.T) {
	app := httpNetService.InitHttpAPP("http://127.0.0.1:8080/post", nil)
	var test struct {
		UserName1 string
		PassWord1 string
	}

	test.UserName1 = "test"
	test.PassWord1 = "test"

	resp, err := app.Post(test)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(resp)
}

func TestHttpPut(t *testing.T) {
	app := httpNetService.InitHttpAPP("http://127.0.0.1:8080/put", nil)
	var test struct {
		UserName1 string
		PassWord1 string
		T         int
		O         map[string]int
		K         float32
		Y         string
	}

	test.UserName1 = "test"
	test.PassWord1 = "test"
	test.T = 1
	test.O = make(map[string]int)
	test.O["test"] = 114
	test.K = 1.1
	test.Y = "test"

	resp, err := app.Put(test)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(resp)
}

func TestHttpDelete(t *testing.T) {
	app := httpNetService.InitHttpAPP("http://127.0.0.1:8080/delete", nil)
	type test1 struct {
		UserName1 string
	}
	var test struct {
		UserName2 string
		Test      test1
	}
	test.UserName2 = "114"
	test.Test.UserName1 = "test"

	resp, err := app.Delete(test)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(resp)
}
