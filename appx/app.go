package appx

import "github.com/gin-gonic/gin"

type Gin struct {
	Ctx       *gin.Context
	CodeMsger ErrCoder
}

type ErrCoder interface {
	GetMsg(code int) string
}

type DefaultErrorCoder struct{}

func (dec DefaultErrorCoder) GetMsg(code int) string {
	switch code {
	case 400:
		return "Bad Request"
	case 401:
		return "Not Authorized"
	case 500:
		return "Internal Server Error"
	default:
		return "Unknown Error"
	}
}

func NewDefaultErrorCoder() ErrCoder {
	return DefaultErrorCoder{}
}
