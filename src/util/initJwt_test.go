// @File: initJwt_test
// @Date: 2020/12/12 9:16
// @Author: 安红豆
// @Description: token生成的测试类
package util

import (
	"fmt"
	"github.com/Ormissia/ormissia_go/src/model"
	"testing"
)

func TestReleaseToken(t *testing.T) {
	user := model.User{
		Username: "ormissia",
	}
	user.ID = 123

	tokenStr, err := ReleaseToken(user)
	if err == nil {
		fmt.Println("Successful ReleaseToken :" + tokenStr)
	} else {
		t.Errorf("ReleaseToken Failed :%s", err)
	}
}
