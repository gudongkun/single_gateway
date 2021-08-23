package middlewares

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/gudongkun/single_common/jaeger_log"
	wrapperTrace "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"io/ioutil"
)

type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w CustomResponseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func AutoJaeger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 初始化span
		rootCtx, rootSpan, _ := wrapperTrace.StartSpanFromContext(
			c.Request.Context(),
			opentracing.GlobalTracer(),
			c.Request.Method+":"+c.FullPath(),
		)
		defer rootSpan.Finish()

		c.Set("rootCtx", rootCtx)
		body, _ := ioutil.ReadAll(c.Request.Body)


		blw := &CustomResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Request = c.Request.WithContext(rootCtx)
		c.Writer = blw

		c.Next()
		jaeger_log.Info(
			rootCtx,
			"common_info",
			log.String("method", c.Request.Method),
			log.String("path", c.Request.RequestURI),
			log.String("body", string(body)),
			log.String("return", blw.body.String()),
		)
	}
}
