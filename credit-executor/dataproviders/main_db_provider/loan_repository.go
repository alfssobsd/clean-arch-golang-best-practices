package main_db_provider

import (
	"clean-arch-golang-best-practices/credit-executor/utils/appconfig"
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

func NewLoanRepository(logger *zap.SugaredLogger, dbConfig appconfig.DatabaseConfiguration) *LoanRepository {
	return &LoanRepository{
		logger:       logger,
		//просто заглушка, в реально жизни сюда уже pgConnect приходит
		dbConnection: fmt.Sprintf("%s,%s,%s", dbConfig.Username, dbConfig.Host, dbConfig.Password),
	}
}

func (repo *LoanRepository) CreateRequestLoan() {
	panic("implement me")
}

func (repo *LoanRepository) CheckRequestLoan() {
	panic("implement me")
}



