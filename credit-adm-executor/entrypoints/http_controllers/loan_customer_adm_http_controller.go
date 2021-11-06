package http_controllers

import (
	"clean-arch-golang-best-practices/credit-adm-executor/usecases"
	"clean-arch-golang-best-practices/credit-adm-executor/utils/appconfig"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
)

type LoanCustomerAdmHttpController struct {
	logger                 *zap.SugaredLogger
	echoServer             *echo.Echo
	appConfig              *appconfig.AppConfiguration
	loanCustomerAdmUseCase *usecases.LoanCustomerAdmUseCase
}

type ILoanCustomerAdmHttpController interface {
	makeRoutes()
}

func NewLoanCustomerAdmHttpController(logger *zap.SugaredLogger, echoServer *echo.Echo, loanCustomerAdmUseCase *usecases.LoanCustomerAdmUseCase) *LoanCustomerAdmHttpController {
	controller := LoanCustomerAdmHttpController{
		logger:                 logger,
		echoServer:             echoServer,
		loanCustomerAdmUseCase: loanCustomerAdmUseCase,
	}
	controller.makeRoutes()
	return &controller
}

func (c *LoanCustomerAdmHttpController) makeRoutes() {
	v1 := c.echoServer.Group("/api/v1/adm/loan-customer")

	v1.POST("/calculate-rating-by-request-of-loan", c.calculateRating)
}

func (c *LoanCustomerAdmHttpController) calculateRating(ctx echo.Context) error {
	c.logger.Infof("Calculate Rating")
	ucResult := c.loanCustomerAdmUseCase.CalculateRatingByRequest(rand.Int())

	response := CalculateRatingHttpResponse{}
	response.SetFromLoanCustomerRatingOutDto(ucResult)

	return ctx.JSON(http.StatusCreated, response)
}
