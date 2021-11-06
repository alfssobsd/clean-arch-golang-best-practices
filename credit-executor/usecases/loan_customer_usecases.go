package usecases

import (
	"clean-arch-golang-best-practices/credit-executor/dataproviders/agify_api_gateway"
	"clean-arch-golang-best-practices/credit-executor/dataproviders/main_db_provider"
	"clean-arch-golang-best-practices/credit-executor/utils/heavyprocessor"
	"go.uber.org/zap"
	"math/rand"
)

type LoanCustomerUseCase struct {
	logger          *zap.SugaredLogger
	heavyProcessor  *heavyprocessor.HeavyProcessor
	loanRepo        *main_db_provider.LoanRepository
	agifyApiGateway *agify_api_gateway.AgifyApiGateway
}

type ILoanCustomerUseCase interface {
	CreateRequestForLoan()
	CheckLoanRequestStatus()
}

func NewLoanCustomerUseCase(logger *zap.SugaredLogger, heavyProcessor *heavyprocessor.HeavyProcessor,
	agifyApiGateway *agify_api_gateway.AgifyApiGateway, loanRepo *main_db_provider.LoanRepository) *LoanCustomerUseCase {
	uc := LoanCustomerUseCase{
		logger:          logger,
		heavyProcessor:  heavyProcessor,
		loanRepo:        loanRepo,
		agifyApiGateway: agifyApiGateway,
	}
	return &uc
}

func (uc *LoanCustomerUseCase) CreateRequestForLoan() {
	_ = uc.heavyProcessor.ExecuteProcessor(rand.Int())
	uc.loanRepo.CreateRequestLoan()
	apiDto, _ := uc.agifyApiGateway.PredicateAgeOfName("vasily", "RU")
	uc.logger.Infof("API %v", apiDto)
}

func (uc *LoanCustomerUseCase) CheckLoanRequestStatus() {
	_ = uc.heavyProcessor.ExecuteProcessor(rand.Int())
	uc.loanRepo.CheckRequestLoan()
}
