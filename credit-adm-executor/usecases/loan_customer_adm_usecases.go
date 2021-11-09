package usecases

import (
	"clean-arch-golang-best-practices/credit-library/loggerhelper"
	"clean-arch-golang-best-practices/credit-shared-module/dataproviders/main_db_provider"
	"clean-arch-golang-best-practices/credit-shared-module/domain"
	"context"
)

type LoanCustomerAdmUseCase struct {
	logger             *loggerhelper.CustomLogger
	creditRatingDomain domain.ICreditRatingDomain
	loanRepo           main_db_provider.ILoanRepository
}

type ILoanCustomerAdmUseCase interface {
	CalculateRatingByRequest(ctx context.Context, requestId int)
}

func NewLoanCustomerAdmUseCase(logger *loggerhelper.CustomLogger, creditRatingDomain domain.ICreditRatingDomain, loanRepo main_db_provider.ILoanRepository) *LoanCustomerAdmUseCase {
	uc := LoanCustomerAdmUseCase{
		logger:             logger,
		creditRatingDomain: creditRatingDomain,
		loanRepo:           loanRepo,
	}
	return &uc
}

func (uc *LoanCustomerAdmUseCase) CalculateRatingByRequest(ctx context.Context, requestId int) LoanCustomerRatingOutDto {
	uc.logger.SugarWithTracing(ctx).Info("LoanCustomerAdmUseCase.CalculateRatingByRequest")
	loanRequestModel := uc.loanRepo.GetRequestLoanByID(ctx, requestId)
	rating := uc.creditRatingDomain.CalculateCreditRating(ctx, loanRequestModel.BorrowerDateOfBirth, loanRequestModel.AnnualIncomeMicros)

	return LoanCustomerRatingOutDto{Rating: rating, RequestId: requestId}
}
