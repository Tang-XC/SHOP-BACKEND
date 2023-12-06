package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

func WrapFunc(f interface{}, args ...interface{}) gin.HandlerFunc {
	//使用反射reflect获取函数f的value
	fn := reflect.ValueOf(f)

	//判断传入的函数f的参数数量和args传入的数量是否相同
	if fn.Type().NumIn() != len(args) {
		panic(fmt.Sprintf("invalid input parameters of function %v", fn.Type()))
	}
	//判断传入的函数f的返回参数是否为0，这里规定函数不可返回空
	if fn.Type().NumOut() == 0 {
		panic(fmt.Sprintf("invalid output parameters of function %v,at least one", fn.Type()))
	}

	//反射所有参数args传过来的值
	inputs := make([]reflect.Value, len(args))
	for k, v := range args {
		inputs[k] = reflect.ValueOf(v)
	}

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				FailedResponse(c, http.StatusInternalServerError, fmt.Errorf("%v", err))
			}
		}()

		outputs := fn.Call(inputs)
		if len(outputs) > 1 {
			err, ok := outputs[len(outputs)-1].Interface().(error)
			if ok && err != nil {
				FailedResponse(c, http.StatusInternalServerError, err)
				return
			}
		}
		c.JSON(http.StatusOK, outputs[0].Interface())
	}
}
