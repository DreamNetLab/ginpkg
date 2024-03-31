package helplerx

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// 测试初始化翻译器
func TestInitTrans(t *testing.T) {
	err := initTrans("zh")
	assert.Nil(t, err, "应该成功初始化中文翻译器")

	err = initTrans("en")
	assert.Nil(t, err, "应该成功初始化英文翻译器")

	err = initTrans("fr")
	assert.NotNil(t, err, "不支持的语言应该返回错误")
}

// 测试ErrorsInUri函数
func TestErrorsInUri(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.GET("/test/:name", func(c *gin.Context) {
		errors := ErrorsInUri(c, &struct {
			Name string `uri:"name" binding:"required,max=4"`
		}{})
		if errors != nil {
			c.JSON(http.StatusBadRequest, errors)
			return
		}
		c.String(http.StatusOK, "success")
	})

	t.Run("成功的URI参数校验", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test/gin", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("失败的URI参数校验", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test/23144", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

// 测试ErrorsInParams函数
func TestErrorsInParams(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.POST("/test", func(c *gin.Context) {
		errors := ErrorsInParams(c, &struct {
			Name string `form:"name" binding:"required"`
		}{})
		if errors != nil {
			c.JSON(http.StatusBadRequest, errors)
			return
		}
		c.String(http.StatusOK, "success")
	})

	t.Run("成功的表单参数校验", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/test", strings.NewReader("name=gin"))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("失败的表单参数校验", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/test", strings.NewReader(""))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
