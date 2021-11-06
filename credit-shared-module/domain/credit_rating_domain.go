package domain

import (
	"go.uber.org/zap"
	"math"
	"time"
)

type CreditRatingDomain struct {
	logger *zap.SugaredLogger
}

type ICreditRatingDomain interface {
	CalculateCreditRating(age int, annualIncomeMicros int64) int
}

func NewCreditRatingDomain(logger *zap.SugaredLogger) *CreditRatingDomain {
	return &CreditRatingDomain{logger: logger}
}

func (crd CreditRatingDomain) CalculateCreditRating(dateOfBirth time.Time, annualIncomeMicros int64) int {
	today := time.Now()
	age := math.Floor(today.Sub(dateOfBirth).Hours() / 24 / 365)
	return int(math.Sqrt(age) - (float64(annualIncomeMicros)/12)*0.3)
}
