package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/rn-consider/compuBackend/config"
	"github.com/rn-consider/compuBackend/utils"
	"go.uber.org/zap"
	"io"
	"strings"
	"time"
)

// bodyLogWriter 结构体用于记录响应体
type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// Write 方法实现了 gin.ResponseWriter 接口的 Write 方法，用于同时记录响应数据并写入底层的 ResponseWriter
func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// WriteString 方法类似于 Write，用于写入字符串
func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

// Logger 是 Gin 中间件函数，用于记录请求和响应的详细信息到日志
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 如果请求 URI 以 "/admin" 开头，跳过不处理
		if strings.HasPrefix(c.Request.RequestURI, "/admin") {
			c.Next()
			return
		}

		// 创建 bodyLogWriter 实例来捕获响应数据
		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter

		// 记录请求开始时间
		startTime := time.Now()

		var request map[string]interface{}

		// 从请求中获取原始 JSON 数据并解析为 request 变量
		b, _ := c.Copy().GetRawData()
		_ = json.Unmarshal(b, &request)
		s, _ := json.Marshal(request)

		// 替换请求的 Body 为不可关闭的 io.NopCloser，以确保后续请求处理可以正常访问 Body 数据
		c.Request.Body = io.NopCloser(bytes.NewReader(b))

		// 执行后续请求处理
		c.Next()

		// 获取捕获的响应数据（从 bodyLogWriter 的 body 缓冲区中获取）
		responseBody := bodyLogWriter.body.String()

		// 记录请求结束时间
		endTime := time.Now()

		// 使用 Zap 日志库记录请求和响应的详细信息
		config.GVA_LOG.Info("请求响应",
			zap.String("request_uri", c.Request.RequestURI),
			zap.String("request_method", c.Request.Method),
			zap.String("client_ip", c.ClientIP()),
			zap.String("request_time", utils.TimeFormat(startTime)),
			zap.String("response_time", utils.TimeFormat(endTime)),
			zap.String("request", string(s)),
			zap.String("response", responseBody),
			zap.String("cost_time", endTime.Sub(startTime).String()),
		)
	}
}
