package ginadaptor

import (
	"bytes"
	"io"
	"log/slog"

	"github.com/gin-gonic/gin"

	"cms-be/internal/pkg/observability/logx"
	"cms-be/internal/pkg/stringutil"
)

func InjectLogx(logger *logx.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request = c.Request.WithContext(logx.ContextWithLogger(c.Request.Context(), logger.Clone()))

		c.Next()
	}
}

func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody, _ := c.GetRawData()
		c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))

		writer := &bodyLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = writer

		c.Next()

		ctx := c.Request.Context()

		logx.Info(ctx, "request",
			slog.String("URL", c.Request.URL.String()),
			slog.String("body", stringutil.FromBytes(requestBody)))
		logx.Info(ctx, "response",
			slog.String("body", writer.body.String()))
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
