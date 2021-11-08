package main_db_provider

import (
	"clean-arch-golang-best-practices/credit-library/loggerhelper"
	"context"
	"fmt"
	"time"
)

type LoanRepository struct {
	logger       *loggerhelper.CustomLogger
	dbConnection string
}

type ILoanRepository interface {
	CreateRequestLoan(ctx context.Context)
	CheckRequestLoan(ctx context.Context)
	GetRequestLoanByID(ctx context.Context, id int) LoanRequestModel
}

func NewLoanRepository(logger *loggerhelper.CustomLogger, dbConfig DatabaseConfiguration) ILoanRepository {
	return &LoanRepository{
		logger: logger,
		//просто заглушка, в реальной жизни сюда уже pgConnect приходит
		dbConnection: fmt.Sprintf("%s,%s,%s", dbConfig.Username, dbConfig.Host, dbConfig.Password),
	}
}

func (repo *LoanRepository) CreateRequestLoan(ctx context.Context) {
	repo.logger.InfofWithTracing(ctx, "LoanRepository.CreateRequestLoan")
}

func (repo *LoanRepository) GetRequestLoanByID(ctx context.Context, id int) LoanRequestModel {
	repo.logger.InfofWithTracing(ctx, "LoanRepository.GetRequestLoanByID")
	layoutISO := "2006-01-02"
	date := "1988-03-01"
	BorrowerDateOfBirth, _ := time.Parse(layoutISO, date)
	return LoanRequestModel{LoanRequestId: id, AnnualIncomeMicros: 1000000000, BorrowerDateOfBirth: BorrowerDateOfBirth}
}

func (repo *LoanRepository) CheckRequestLoan(ctx context.Context) {
	repo.logger.InfofWithTracing(ctx, "LoanRepository.CheckRequestLoan")
}
