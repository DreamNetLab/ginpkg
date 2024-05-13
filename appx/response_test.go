package appx

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRespSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/test/success", func(c *gin.Context) {
		g := &Gin{Ctx: c}
		g.RespSuccess(map[string]string{"hello": "world"})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test/success", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response stdResp
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, StatusSuccess, response.Code)
	assert.Equal(t, "success", response.Msg)
	assert.Equal(t, map[string]interface{}{"hello": "world"}, response.Data)
}

func TestRespInvalidParams(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/test/invalid", func(c *gin.Context) {
		g := &Gin{Ctx: c}
		g.RespInvalidParams(map[string]string{"error": "missing parameter"})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test/invalid", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	var response stdResp
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, StatusBadRequest, response.Code)
	assert.Equal(t, "invalid params", response.Msg)
	assert.Equal(t, map[string]interface{}{"error": "missing parameter"}, response.Data)
}

func TestRespBadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/test/badrequest", func(c *gin.Context) {
		g := &Gin{Ctx: c}
		g.RespBadRequest(401)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test/badrequest", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	var response stdResp
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	// 由于RespBadRequest函数内部不设置msg，所以这里的msg取决于CodeMsger接口的实现或默认错误消息
	// 这里的测试期望值可能需要根据实际情况调整
	assert.Equal(t, 401, response.Code)
	assert.NotEqual(t, "", response.Msg) // 确保有错误消息返回
}

func TestRespError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/test/error", func(c *gin.Context) {
		g := &Gin{Ctx: c}
		g.RespError(501)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test/error", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	var response stdResp
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	// 由于RespError函数内部不设置msg，所以这里的msg取决于CodeMsger接口的实现或默认错误消息
	// 这里的测试期望值可能需要根据实际情况调整
	assert.Equal(t, 501, response.Code)
	assert.NotEqual(t, "", response.Msg) // 确保有错误消息返回
}
