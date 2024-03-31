package appx

import (
	"testing"
)

func TestDefaultErrorCoder_GetMsg(t *testing.T) {
	coder := NewDefaultErrorCoder()

	tests := []struct {
		code     int
		expected string
	}{
		{400, "Bad Request"},
		{500, "Internal Server Error"},
		{999, "Unknown Error"}, // 负面测试用例：非预期的错误代码
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			msg := coder.GetMsg(tt.code)
			if msg != tt.expected {
				t.Errorf("GetMsg(%d) = %s; want %s", tt.code, msg, tt.expected)
			}
		})
	}
}
