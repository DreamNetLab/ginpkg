package helplerx

import (
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"
)

func TestGetRealIP(t *testing.T) {
	// 设置测试用例
	tests := []struct {
		name          string
		xRealIP       string
		xForwardedFor string
		remoteAddr    string
		expectedIP    string
	}{
		{"X-Real-IP", "192.168.1.1", "", "", "192.168.1.1"},
		{"X-Forwarded-For", "", "192.168.1.2", "", "192.168.1.2"},
		{"RemoteAddr", "", "", "192.168.1.3:12345", "192.168.1.3:12345"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建请求
			req := httptest.NewRequest("GET", "/", nil)
			if tt.xRealIP != "" {
				req.Header.Add("X-Real-IP", tt.xRealIP)
			}
			if tt.xForwardedFor != "" {
				req.Header.Add("X-Forwarded-For", tt.xForwardedFor)
			}

			// 创建响应记录器
			w := httptest.NewRecorder()

			// 创建gin上下文
			c, _ := gin.CreateTestContext(w)
			c.Request = req

			// 模拟远程地址
			if tt.remoteAddr != "" {
				c.Request.RemoteAddr = tt.remoteAddr
			}

			// 调用函数并验证结果
			ip := GetRealIP(c)
			if ip != tt.expectedIP {
				t.Errorf("expected IP %s, got %s", tt.expectedIP, ip)
			}
		})
	}
}

func TestFindStrSliceElem(t *testing.T) {
	// 设置测试用例
	tests := []struct {
		name     string
		elem     string
		set      []string
		expected int
		found    bool
	}{
		{"Found", "b", []string{"a", "b", "c"}, 1, true},
		{"NotFound", "d", []string{"a", "b", "c"}, -1, false},
		{"EmptySet", "a", []string{}, -1, false},
		{"FirstElement", "a", []string{"a", "b", "c"}, 0, true},
		{"LastElement", "c", []string{"a", "b", "c"}, 2, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index, found := FindStrSliceElem(tt.elem, tt.set)
			if index != tt.expected || found != tt.found {
				t.Errorf("expected index %d and found %t, got index %d and found %t", tt.expected, tt.found, index, found)
			}
		})
	}
}
