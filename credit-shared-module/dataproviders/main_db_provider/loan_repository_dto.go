package main_db_provider

import "time"

type LoanRequestModel struct {
	LoanRequestId       int
	BorrowerDateOfBirth time.Time
	AnnualIncomeMicros  int64
}
