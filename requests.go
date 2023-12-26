package openapi

import (
	"bytes"
	"github.com/buzzxu/boys/common/httpsclient"
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"time"
)

// openApiRequestBody 表示请求体的数据结构
type requestBody struct {
	AppKey    string `json:"appkey"`
	Timestamp int64  `json:"timestamp"`
	Sign      string `json:"sign"`
	Data      string `json:"data"`
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// DoIt 发送请求
func DoIt[T any, R any](url, appKey, appSecret string, data T, result R, funcHeader func(header http.Header)) error {
	return Call(url, appKey, appSecret, data, funcHeader, func(response *http.Response) error {
		return json.NewDecoder(response.Body).Decode(result)
	})
}

// Call 发送请求
func Call[T any](url, appKey, appSecret string, data T, funcHeader func(header http.Header), funcResponse func(response *http.Response) error) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return Request(url, appKey, appSecret, string(b), funcHeader, funcResponse)
}

// Request 发送请求
func Request(url, appKey, appSecret, data string, funcHeader func(header http.Header), funcResponse func(response *http.Response) error) error {
	timestamp := time.Now().Unix()
	signature := Signature(appKey, appSecret, data, timestamp)
	body := requestBody{
		AppKey:    appKey,
		Timestamp: timestamp,
		Sign:      signature,
		Data:      data,
	}
	b, err := json.Marshal(body)
	if err != nil {
		return err
	}
	err = httpsclient.Https("POST", url, bytes.NewBuffer(b), func(header http.Header) {
		header.Set("Content-Type", "application/json")
		if funcHeader != nil {
			funcHeader(header)
		}
	}, funcResponse)
	if err != nil {
		return err
	}
	return nil
}
