package http_controllers

import (
	"clean-arch-golang-best-practices/credit-executor/usecases"
	"clean-arch-golang-best-practices/credit-executor/utils/appconfig"
	"clean-arch-golang-best-practices/credit-library/loggerhelper"
	"github.com/labstack/echo/v4"
	"net/http"
)

type LoanCustomerHttpController struct {
	logger              *loggerhelper.CustomLogger
	echoServer          *echo.Echo
	appConfig           *appconfig.AppConfiguration
	loanCustomerUseCase usecases.ILoanCustomerUseCase
}

type ILoanCustomerHttpController interface {
	makeRoutes()
}

func NewLoanCustomerHttpController(logger *loggerhelper.CustomLogger, echoServer *echo.Echo, loanCustomerUseCase usecases.ILoanCustomerUseCase) ILoanCustomerHttpController {
	controller := LoanCustomerHttpController{
		logger:              logger,
		echoServer:          echoServer,
		loanCustomerUseCase: loanCustomerUseCase,
	}
	controller.makeRoutes()
	return &controller
}

func (c *LoanCustomerHttpController) makeRoutes() {
	v1 := c.echoServer.Group("/api/v1/loan-customer")

	v1.POST("/create-request-for-loan", c.createRequestForLoan)
	v1.GET("/check-loan-request-status", c.checkLoanRequestStatus)
}

func (c *LoanCustomerHttpController) createRequestForLoan(ctx echo.Context) error {
	c.logger.SugarWithTracing(ctx.Request().Context()).Infof("Create request for loan")
	c.loanCustomerUseCase.CreateRequestForLoan(ctx.Request().Context())

	return ctx.JSON(http.StatusCreated, "CREATED")
}

func (c *LoanCustomerHttpController) checkLoanRequestStatus(ctx echo.Context) error {
	c.logger.SugarWithTracing(ctx.Request().Context()).Infof("Check loan request status")
	c.loanCustomerUseCase.CheckLoanRequestStatus(ctx.Request().Context())

	return ctx.JSON(http.StatusOK, "IN PROGRESS")
}
