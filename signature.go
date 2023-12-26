package openapi

import "github.com/buzzxu/boys/common/signature/aksk"

// Signature 签名
func Signature(appKey, appSecret, data string, timestamp int64) string {
	return aksk.SHA1(appKey, appSecret, data, timestamp)
}
