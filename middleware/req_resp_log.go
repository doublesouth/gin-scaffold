package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"io/ioutil"
	"net/http"
	"time"
	"github.com/doublesouth/gin-scaffold/common"
)

// middleware for log request and response data
func RequestAndResponseLogger(duration time.Duration) gin.HandlerFunc {
	setupLogging(duration)
	return func(c *gin.Context) {
		method := c.Request.Method
		requestId, _ := c.Get(common.ContextRequestId)
		if method == http.MethodGet || method == http.MethodDelete {
			glog.Infof("request: uri => %s, method => %s, request_id => %s, access_token => %s, body => nil",
				c.Request.RequestURI,
				method,
				requestId,
				c.Request.Header.Get(common.RequestHeaderAccessToken),
			)
		} else if method == http.MethodPost || method == http.MethodPut {
			body, err := ioutil.ReadAll(c.Request.Body)
			if err != nil {
				glog.Infof("ioutil.ReadAll() cause err: %s", err)
				return
			}
			glog.Infof("request: uri => %s, method => %s, request_id => %s, client_ip => %s, access_token => %s, body => %s",
				c.Request.RequestURI,
				method,
				requestId,
				c.ClientIP(),
				c.Request.Header.Get(common.RequestHeaderAccessToken),
				string(body),
			)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		} else {
			// TODO
		}

		t := time.Now()
		c.Next()
		cost := time.Since(t)

		glog.Infof("response: uri => %s, method => %s, request_id => %s, client_ip => %s, access_token => %s, status => %d, cost => %v",
			c.Request.RequestURI,
			method,
			requestId,
			c.ClientIP(),
			c.Request.Header.Get(common.RequestHeaderAccessToken),
			c.Writer.Status(),
			cost,
		)
	}
}

func setupLogging(duration time.Duration) {
	go func() {
		for range time.Tick(duration) {
			glog.Flush()
		}
	}()
}
