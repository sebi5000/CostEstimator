package model

import "errors"

type LicenseCalculation struct {
	Count    int
	Discount int
	Type     LicenseType
	Module   string
	addOns   []LicenseCalculation
}

type LicenseType struct {
	Type      string
	listPrice int
}

func (lt LicenseType) GetListPrice() int {

	switch lt.Type {
	case "enterprise":
		lt.listPrice = 165
	case "unlimited":
		lt.listPrice = 330
	case "einstein1":
		lt.listPrice = 500
	case "addon":
		lt.listPrice = 70
	}

	return lt.listPrice
}

func (lc LicenseCalculation) CalculatePrice() int {
	if lc.Discount <= 0 ||
		lc.Discount > 100 {
		return lc.Count * lc.Type.GetListPrice()
	} else {
		return lc.Count * lc.Type.GetListPrice() * (1 - lc.Discount/100)
	}
}

func (lc *LicenseCalculation) GetAddOns() []LicenseCalculation {
	return lc.addOns
}

func (lc *LicenseCalculation) AddAddOn(name string, addOn LicenseCalculation) error {

	var err error = nil

	if lc.Module == "sales" &&
		lc.Type.Type == "enterprise" &&
		(name == "einstein_for_sales" ||
			name == "einstein_conversational_intelligence" ||
			name == "einstein_relationship_insights") {
		lc.addOns = append(lc.addOns, addOn)
	} else if lc.Module == "sales" &&
		lc.Type.Type == "unlimited" &&
		name == "einstein_for_saless" {
		lc.addOns = append(lc.addOns, addOn)
	} else if lc.Module == "service" &&
		lc.Type.Type == "enterprise" &&
		(name == "einstein_for_service" ||
			name == "einstein_bots") {
		lc.addOns = append(lc.addOns, addOn)
	} else {
		err = errors.New("invalid add-on can't be added to your license configuration")
	}

	return err
}
