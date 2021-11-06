package main_db_provider

import (
	"fmt"
	"go.uber.org/zap"
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
		//просто заглушка, в реально жизни сюда уже pgConnect приходит
		dbConnection: fmt.Sprintf("%s,%s,%s", dbConfig.Username, dbConfig.Host, dbConfig.Password),
	}
}

func (repo *LoanRepository) CreateRequestLoan() {
	repo.logger.Infof("LoanRepository.CreateRequestLoan")
}

func (repo *LoanRepository) CheckRequestLoan() {
	repo.logger.Infof("LoanRepository.CheckRequestLoan")
}
