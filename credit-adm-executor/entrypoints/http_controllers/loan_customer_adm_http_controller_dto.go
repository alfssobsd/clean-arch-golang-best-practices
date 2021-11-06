package http_controllers

import (
	"clean-arch-golang-best-practices/credit-adm-executor/usecases"
	"time"
)

type CalculateRatingHttpResponse struct {
	RequestId int
	Rating    int
	CreatedAt time.Time
}

// SetFromLoanCustomerRatingOutDto Конвертировани из одного DTO в другое для использование между слоями
func (dto CalculateRatingHttpResponse) SetFromLoanCustomerRatingOutDto(inDto usecases.LoanCustomerRatingOutDto) {
	dto.RequestId = inDto.RequestId
	dto.Rating = inDto.RequestId
	dto.CreatedAt = time.Now()
}
