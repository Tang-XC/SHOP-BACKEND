package common

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"'`
	Data interface{} `json:"data"`
}

var (
	CookieTokenName = `token`
	CookieLoginUser = `loginUser`
)

func NewResponse(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(code, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func SuccessResponse(c *gin.Context, data interface{}) {
	NewResponse(c, http.StatusOK, "success", data)
}

func FailedResponse(c *gin.Context, code int, err error) {
	if code == 0 {
		code = http.StatusInternalServerError
	}
	if code == http.StatusUnauthorized && c.Request != nil {
		if val, err := c.Cookie(CookieTokenName); err == nil && val != "" {
			c.SetCookie(CookieTokenName, "", -1, "/", "", true, true)
			c.SetCookie(CookieLoginUser, "", -1, "/", "", true, false)
		}
	}
	var msg string
	if err != nil {
		msg = err.Error()
		var url string
		if c.Request != nil {
			url = c.Request.URL.String()
		}
		logrus.Warnf("url: %s, user: %s, error: %v", url, msg)
	}
	NewResponse(c, code, msg, nil)
}
