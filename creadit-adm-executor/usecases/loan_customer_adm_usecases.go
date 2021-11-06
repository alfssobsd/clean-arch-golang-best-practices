package usecases

import (
	"clean-arch-golang-best-practices/credit-shared-module/dataproviders/main_db_provider"
	"clean-arch-golang-best-practices/credit-shared-module/domain"
	"go.uber.org/zap"
)

type LoanCustomerAdmUseCase struct {
	logger             *zap.SugaredLogger
	creditRatingDomain *domain.CreditRatingDomain
	loanRepo           *main_db_provider.LoanRepository
}

type ILoanCustomerAdmUseCase interface {
	CalculateRatingByRequest(requestId int)
}

func NewLoanCustomerAdmUseCase(logger *zap.SugaredLogger, creditRatingDomain *domain.CreditRatingDomain, loanRepo *main_db_provider.LoanRepository) *LoanCustomerAdmUseCase {
	uc := LoanCustomerAdmUseCase{
		logger:             logger,
		creditRatingDomain: creditRatingDomain,
		loanRepo:           loanRepo,
	}
	return &uc
}

func (uc *LoanCustomerAdmUseCase) CalculateRatingByRequest(requestId int) LoanCustomerRatingOutDto {
	uc.logger.Infof("LoanCustomerAdmUseCase.CalculateRatingByRequest")
	loanRequestModel := uc.loanRepo.GetRequestLoanByID(requestId)
	rating := uc.creditRatingDomain.CalculateCreditRating(loanRequestModel.BorrowerDateOfBirth, loanRequestModel.AnnualIncomeMicros)

	return LoanCustomerRatingOutDto{Rating: rating, RequestId: requestId}
}
