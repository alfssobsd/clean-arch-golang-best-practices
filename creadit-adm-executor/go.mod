module clean-arch-golang-best-practices/credit-adm-executor

go 1.17

require (
	clean-arch-golang-best-practices/credit-library v0.0.0
	clean-arch-golang-best-practices/credit-shared-module v0.0.0
	github.com/labstack/echo-contrib v0.11.0
	github.com/labstack/echo/v4 v4.6.1
	github.com/spf13/viper v1.9.0
	go.uber.org/zap v1.19.1
)

replace clean-arch-golang-best-practices/credit-shared-module => ../credit-shared-module
replace clean-arch-golang-best-practices/credit-library => ../credit-library