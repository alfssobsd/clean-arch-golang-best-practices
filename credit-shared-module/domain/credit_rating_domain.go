package domain

import (
	"clean-arch-golang-best-practices/credit-library/loggerhelper"
	"context"
	"math"
	"time"
)

type CreditRatingDomain struct {
	logger *loggerhelper.CustomLogger
}

type ICreditRatingDomain interface {
	CalculateCreditRating(ctx context.Context, dateOfBirth time.Time, annualIncomeMicros int64) int
}

func NewCreditRatingDomain(logger *loggerhelper.CustomLogger) ICreditRatingDomain {
	return &CreditRatingDomain{logger: logger}
}

func (crd CreditRatingDomain) CalculateCreditRating(ctx context.Context, dateOfBirth time.Time, annualIncomeMicros int64) int {
	crd.logger.InfofWithTracing(ctx, "CreditRatingDomain.CalculateCreditRating")
	today := time.Now()
	age := math.Floor(today.Sub(dateOfBirth).Hours() / 24 / 365)
	return int(math.Sqrt(age) - (float64(annualIncomeMicros)/12)*0.3)
}
