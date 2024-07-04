package services

import (
	"costestimator/cmd/model"
)

type LicenseCalcService struct {
	licenseCalculations []model.LicenseCalculation
}

func NewLicenseCalcService(licenseCalculations []model.LicenseCalculation) LicenseCalcService {
	return LicenseCalcService{licenseCalculations: licenseCalculations}
}

func (lcs LicenseCalcService) CalcOverallCosts() int {

	var price float64

	for _, lc := range lcs.licenseCalculations {
		if lc.Discount > 0 {
			price += float64(lc.Type.GetListPrice()) * float64(lc.Count) * (1 - float64(lc.Discount)/100.0)
		} else {
			price += float64(lc.Type.GetListPrice()) * float64(lc.Count)
		}

		for _, addon := range lc.GetAddOns() {
			if addon.Discount > 0 {
				price += float64(addon.Type.GetListPrice()) * float64(addon.Count) * (1 - float64(addon.Discount)/100.0)
			} else {
				price += float64(addon.Type.GetListPrice()) * float64(addon.Count)
			}
		}
	}

	return int(price)
}

func (lcs LicenseCalcService) CalcModuleCosts(module string) int {

	var price float64

	for _, lc := range lcs.licenseCalculations {
		if lc.Module == module {
			if lc.Discount > 0 {
				price += float64(lc.Type.GetListPrice()) * float64(lc.Count) * (1 - float64(lc.Discount)/100.0)
			} else {
				price += float64(lc.Type.GetListPrice()) * float64(lc.Count)
			}

			for _, addon := range lc.GetAddOns() {
				if addon.Discount > 0 {
					price += float64(addon.Type.GetListPrice()) * float64(addon.Count) * (1 - float64(addon.Discount)/100.0)
				} else {
					price += float64(addon.Type.GetListPrice()) * float64(addon.Count)
				}
			}
		}
	}

	return int(price)
}
