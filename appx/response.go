package appx

import "net/http"

const (
	StatusSuccess       = http.StatusOK
	StatusBadRequest    = http.StatusBadRequest
	StatusUnAuth        = http.StatusUnauthorized
	StatusForbidden     = http.StatusForbidden
	StatusInternalError = http.StatusInternalServerError
)

var defaultErrMsg = map[int]string{
	400: "invalid params",
	401: "not authorized",
	403: "forbidden",
	500: "internal error",
}

type stdResp struct {
	// 状态码: 200成功，400失败,500错误
	Code int `json:"code" binding:"required"`
	// 返回信息
	Msg string `json:"msg" binding:"required"`
	// 返回data
	Data any `json:"data"`
}

func (g *Gin) respond(httpStatus, code int, msg string, data any) {
	if msg == "" {
		if g.CodeMsger != nil {
			msg = g.CodeMsger.GetMsg(code)
		} else {
			msg = defaultErrMsg[code]
			if msg == "" {
				msg = "unknown error"
			}
		}
	}

	g.Ctx.JSON(httpStatus, stdResp{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func (g *Gin) RespSuccess(data any) {
	g.respond(StatusSuccess, 200, "success", data)
}

func (g *Gin) RespInvalidParams(errors any) {
	g.respond(StatusBadRequest, 400, "invalid params", errors)
}

func (g *Gin) RespBadRequest(code int) {
	g.respond(StatusBadRequest, code, "", map[string]any{})
}

func (g *Gin) RespUnAuth() {
	g.respond(StatusUnAuth, 401, "not authorized", map[string]any{})
}

func (g *Gin) RespForbidden() {
	g.respond(StatusForbidden, 403, "forbidden", map[string]any{})
}

func (g *Gin) RespError(code int) {
	g.respond(StatusInternalError, code, "", map[string]any{})
}
