package services

import (
	"costestimator/cmd/model"
	"math"
)

const requestDivider = 1500

type RequestCalculationService struct {
	Request model.RequestCalculation
}

func (rcs *RequestCalculationService) CalculateRequests() int {

	words := rcs.Request.EinsteinInput + rcs.Request.EinsteintOutput

	divWords := float64(words) / float64(requestDivider)
	roundedCount := math.Ceil(divWords)

	return int(roundedCount) * rcs.Request.Model.Multiplier
}
