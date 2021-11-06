module clean-arch-golang-best-practices/credit-executor

go 1.17

require (
	github.com/labstack/echo/v4 v4.4.0
	github.com/spf13/viper v1.8.1
	go.uber.org/zap v1.19.0
	clean-arch-golang-best-practices/credit-library v0.0.0
)

replace clean-arch-golang-best-practices/credit-library => ../credit-library