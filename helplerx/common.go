package helplerx

import (
	"github.com/gin-gonic/gin"
)

func GetRealIP(ctx *gin.Context) string {
	ip := ctx.Request.Header.Get("X-Real-IP")

	if ip == "" {
		ip = ctx.Request.Header.Get("X-Forwarded-For")
	}

	if ip == "" {
		ip = ctx.Request.RemoteAddr
	}

	return ip
}

func FindStrSliceElem(elem string, set []string) (int, bool) {
	for k, v := range set {
		if v == elem {
			return k, true
		}
	}

	return -1, false
}
