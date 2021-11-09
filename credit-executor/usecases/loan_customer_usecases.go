package usecases

import (
	"clean-arch-golang-best-practices/credit-executor/dataproviders/agify_api_gateway"
	"clean-arch-golang-best-practices/credit-executor/dataproviders/redis_repository"
	"clean-arch-golang-best-practices/credit-library/loggerhelper"
	"clean-arch-golang-best-practices/credit-shared-module/dataproviders/main_db_provider"
	"clean-arch-golang-best-practices/credit-shared-module/utils/heavyprocessor"
	"context"
	"math/rand"
)

type LoanCustomerUseCase struct {
	logger          *loggerhelper.CustomLogger
	heavyProcessor  heavyprocessor.IHeavyProcessor
	loanRepo        main_db_provider.ILoanRepository
	agifyApiGateway agify_api_gateway.IAgifyApiGateway
	cacheRepo       redis_repository.ISessionCacheRepository
}

type ILoanCustomerUseCase interface {
	CreateRequestForLoan(ctx context.Context)
	CheckLoanRequestStatus(ctx context.Context)
}

func NewLoanCustomerUseCase(logger *loggerhelper.CustomLogger, heavyProcessor heavyprocessor.IHeavyProcessor,
	agifyApiGateway agify_api_gateway.IAgifyApiGateway, loanRepo main_db_provider.ILoanRepository, cacheRepo redis_repository.ISessionCacheRepository) ILoanCustomerUseCase {
	uc := LoanCustomerUseCase{
		logger:          logger,
		heavyProcessor:  heavyProcessor,
		loanRepo:        loanRepo,
		agifyApiGateway: agifyApiGateway,
		cacheRepo:       cacheRepo,
	}
	return &uc
}

func (uc *LoanCustomerUseCase) CreateRequestForLoan(ctx context.Context) {
	uc.logger.SugarWithTracing(ctx).Infof("LoanCustomerUseCase.CreateRequestForLoan")
	_ = uc.heavyProcessor.ExecuteProcessor(ctx, rand.Int())
	uc.loanRepo.CreateRequestLoan(ctx)
	apiDto, _ := uc.agifyApiGateway.PredicateAgeOfName(ctx, "vasily", "RU")
	uc.logger.SugarWithTracing(ctx).Infof("API %v", apiDto)
	err := uc.cacheRepo.CreateSession(ctx, "SomeKey", "value")
	if err != nil {
		return
	}
}

func (uc *LoanCustomerUseCase) CheckLoanRequestStatus(ctx context.Context) {
	uc.logger.SugarWithTracing(ctx).Infof("LoanCustomerUseCase.CheckLoanRequestStatus")
	_ = uc.heavyProcessor.ExecuteProcessor(ctx, rand.Int())
	uc.loanRepo.CheckRequestLoan(ctx)
	val, err := uc.cacheRepo.GetSession(ctx, "SomeKey")
	if err != nil {
		return
	}
	uc.logger.SugarWithTracing(ctx).Infof("LoanCustomerUseCase.CreateRequestForLoan session %s", val)

}
