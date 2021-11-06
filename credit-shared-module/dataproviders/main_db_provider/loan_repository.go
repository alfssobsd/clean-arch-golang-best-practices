package main_db_provider

import (
	"fmt"
	"go.uber.org/zap"
	"time"
)

type LoanRepository struct {
	logger       *zap.SugaredLogger
	dbConnection string
}

type ILoanRepository interface {
	CreateRequestLoan()
	CheckRequestLoan()
}

func NewLoanRepository(logger *zap.SugaredLogger, dbConfig DatabaseConfiguration) *LoanRepository {
	return &LoanRepository{
		logger: logger,
		//просто заглушка, в реальной жизни сюда уже pgConnect приходит
		dbConnection: fmt.Sprintf("%s,%s,%s", dbConfig.Username, dbConfig.Host, dbConfig.Password),
	}
}

func (repo *LoanRepository) CreateRequestLoan() {
	repo.logger.Infof("LoanRepository.CreateRequestLoan")
}

func (repo *LoanRepository) GetRequestLoanByID(id int) LoanRequestModel {
	repo.logger.Infof("LoanRepository.GetRequestLoanByID")
	layoutISO := "2006-01-02"
	date := "1988-03-01"
	BorrowerDateOfBirth, _ := time.Parse(layoutISO, date)
	return LoanRequestModel{LoanRequestId: id, AnnualIncomeMicros: 1000000000, BorrowerDateOfBirth: BorrowerDateOfBirth}
}

func (repo *LoanRepository) CheckRequestLoan() {
	repo.logger.Infof("LoanRepository.CheckRequestLoan")
}
