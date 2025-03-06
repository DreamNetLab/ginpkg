package sts

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenStsToken(t *testing.T) {
	// 正向测试用例
	t.Run("正向测试: 生成STS Token成功", func(t *testing.T) {
		resp, err := GenStsToken("LTAI5tM8wEY3V6", "ONIhGMdRORn2UapH7DcxkO",
			"sts.cn-shanghai.aliyuncs.com", "acs:ram::1289146:role/ramoss")
		fmt.Println(resp)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		assert.NotEmpty(t, resp.AccessID)
		assert.NotEmpty(t, resp.Secret)
		assert.NotEmpty(t, resp.Token)
		assert.NotEmpty(t, resp.Expire)
	})

	// 负向测试用例
	t.Run("负向测试: 无效的AccessKeyId", func(t *testing.T) {
		resp, err := GenStsToken("", "ONIhGMdRORn2UcxkO", "sts.cn-shanghai.aliyuncs.com", "acs:ram::1286:role/ramoss")
		assert.NotNil(t, err)
		assert.Nil(t, resp)
	})

	t.Run("负向测试: 无效的AccessKeySecret", func(t *testing.T) {
		resp, err := GenStsToken("LTAI5tf6", "", "sts.cn-shanghai.aliyuncs.com", "acs:ram::12895546:role/ramoss")
		assert.NotNil(t, err)
		assert.Nil(t, resp)
	})

	t.Run("负向测试: 无效的Endpoint", func(t *testing.T) {
		resp, err := GenStsToken("LTAI5tM8Hhx2f6", "ONIhGMdRORn1ZH7DcxkO", "", "acs:ram::12895546:role/ramoss")
		assert.NotNil(t, err)
		assert.Nil(t, resp)
	})

	t.Run("负向测试: 无效的Arn", func(t *testing.T) {
		resp, err := GenStsToken("LTAI5Hhx2f6", "ONIhGMdRO7DcxkO", "sts.cn-shanghai.aliyuncs.com", "")
		assert.NotNil(t, err)
		assert.Nil(t, resp)
	})
}
