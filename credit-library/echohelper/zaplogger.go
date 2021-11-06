package echohelper

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func EchoZapLogger(log *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			startAt := time.Now()
			err := next(c)

			if err != nil {
				c.Error(err)
			}

			request := c.Request()
			response := c.Response()

			zapFields := []zapcore.Field{
				zap.Namespace("meta"),
				zap.String("time", time.Since(startAt).String()),
				zap.String("host", request.Host),
				zap.String("real_ip", c.RealIP()),
				zap.Int("response_status", response.Status),
				zap.Int64("response_size", response.Size),
				zap.String("user_agent", request.UserAgent()),
			}

			statusCode := response.Status
			logString := fmt.Sprintf("%s %s %s status=%d", request.Method, request.RequestURI, request.Proto, statusCode)
			switch {
			case statusCode >= 500:
				log.Error(logString, zapFields...)
			case statusCode >= 400:
				log.Warn(logString, zapFields...)
			default:
				log.Info(logString, zapFields...)
			}

			return nil
		}
	}
}
