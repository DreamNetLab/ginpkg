package helplerx

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en2 "github.com/go-playground/validator/v10/translations/en"
	zh2 "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

var (
	mimeType string
	Trans    ut.Translator
)

func init() {
	setTransTagMimeType("label")
	if err := initTrans("zh"); err != nil {
		fmt.Println("Fail to initial validator translator!")
		panic("Fail to initial validator translator")
	}
}

func setTransTagMimeType(mt string) {
	mType := "json"
	if mt != "" {
		mType = mt
	}
	mimeType = mType
}

func getStructFieldTagName(field reflect.StructField) string {
	name := strings.SplitN(field.Tag.Get(mimeType), ",", 2)[0]
	if name == "-" {
		return ""
	}

	return name
}

func removeTopStruct(fields map[string]string) map[string]string {
	resp := make(map[string]string)
	for field, err := range fields {
		resp[field[strings.Index(field, ".")+1:]] = err
	}

	return resp
}

func initTrans(locale string) (err error) {
	//  修改gin框架中的validator引擎属性，实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册一个标签名函数
		v.RegisterTagNameFunc(getStructFieldTagName)

		zhT := zh.New()
		enT := en.New()

		//  第一个参数是备用语言环境，后续参数是应该支持的语言环境
		uni := ut.New(enT, zhT, enT)
		Trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", locale)
		}

		// 语言环境选择
		switch locale {
		case "en":
			err = en2.RegisterDefaultTranslations(v, Trans)
		case "zh":
			err = zh2.RegisterDefaultTranslations(v, Trans)
		default:
			err = en2.RegisterDefaultTranslations(v, Trans)
		}
		return
	}
	return
}

// ErrorsInUri 绑定uri校验
func ErrorsInUri(ctx *gin.Context, instance any) map[string]string {
	if err := ctx.ShouldBindUri(instance); err != nil {
		var errs validator.ValidationErrors
		if ok := errors.As(err, &errs); !ok {
			return nil
		}

		return removeTopStruct(errs.Translate(Trans))
	}

	return nil
}

// ErrorsInParams 绑定校验参数
func ErrorsInParams(ctx *gin.Context, instance any) map[string]string {
	if err := ctx.ShouldBind(instance); err != nil {
		var errs validator.ValidationErrors
		if ok := errors.As(err, &errs); !ok {
			return nil
		}

		return removeTopStruct(errs.Translate(Trans))
	}

	return nil
}
