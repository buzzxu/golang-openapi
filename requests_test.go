package openapi

import (
	"testing"
)

func TestDoIt(t *testing.T) {
	appkey := "xwjd"
	secret := "XWJD010230441222212312313V"
	result := make(map[string]interface{})
	error := DoIt("https://api.mgr.xwjd.xingchenga.xyz/open/test", appkey, secret, "", &result, nil)
	if error != nil {
		t.Error(error)
	}
	t.Log(result["data"])
}
