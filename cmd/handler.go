package main

import (
	htmx "costestimator/cmd/HTMX"
	"costestimator/cmd/model"
	"costestimator/cmd/model/status"
	"costestimator/cmd/services"
	"costestimator/cmd/views"
	licensecalculator "costestimator/cmd/views/components/license_calculator"
	requestcalculator "costestimator/cmd/views/components/request_calculator"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(views.Index()).ServeHTTP(w, r)
}

func requestCalculationHandler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(views.RequestCalculation()).ServeHTTP(w, r)
}

func calculatePriceHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	var salesCalc model.LicenseCalculation
	salesCalc.Module = "sales"
	var serviceCalc model.LicenseCalculation
	serviceCalc.Module = "service"
	statusFeedback := status.Success("")
	htmxService := htmx.NewService(w)

	exp_discount, err := strconv.Atoi(r.Form.Get("expected_discount"))

	if err != nil {
		statusFeedback = status.Danger(err.Error())
		htmxService.AddEvent(htmx.Event{Name: "onFeedback", Param: statusFeedback.Error()})
		return
	}

	//SALES CLOUD

	salesCalc.Count, err = strconv.Atoi(r.Form.Get("sales_user"))

	if err != nil {
		statusFeedback = status.Danger(err.Error())
		htmxService.AddEvent(htmx.Event{Name: "onFeedback", Param: statusFeedback.Error()})
		return
	}

	salesLicenseType := &model.LicenseType{Type: r.Form.Get("sales_license")}
	salesCalc.Type = *salesLicenseType
	salesCalc.Discount = exp_discount

	if r.Form.Get("einstein_for_sales") == "on" {
		salesCalc.AddAddOn("einstein_for_sales", model.LicenseCalculation{Module: "sales", Count: salesCalc.Count, Type: model.LicenseType{Type: "addon"}, Discount: exp_discount})
	}

	if r.Form.Get("einstein_conversational_intelligence") == "on" {
		salesCalc.AddAddOn("einstein_conversational_intelligence", model.LicenseCalculation{Module: "sales", Count: salesCalc.Count, Type: model.LicenseType{Type: "addon"}, Discount: exp_discount})
	}

	if r.Form.Get("einstein_relationship_insights") == "on" {
		salesCalc.AddAddOn("einstein_relationship_insights", model.LicenseCalculation{Module: "sales", Count: salesCalc.Count, Type: model.LicenseType{Type: "addon"}, Discount: exp_discount})
	}

	//SERVICE CLOUD

	serviceCalc.Count, err = strconv.Atoi(r.Form.Get("service_user"))

	if err != nil {
		statusFeedback = status.Danger(err.Error())
		htmxService.AddEvent(htmx.Event{Name: "onFeedback", Param: statusFeedback.Error()})
		return
	}

	serviceLicenseType := &model.LicenseType{Type: r.Form.Get("service_license")}
	serviceCalc.Type = *serviceLicenseType
	serviceCalc.Discount = exp_discount

	if r.Form.Get("einstein_for_service") == "on" {
		serviceCalc.AddAddOn("einstein_for_service", model.LicenseCalculation{Module: "service", Count: serviceCalc.Count, Type: model.LicenseType{Type: "addon"}, Discount: exp_discount})

	}

	if r.Form.Get("einstein_bots") == "on" {
		serviceCalc.AddAddOn("einstein_bots", model.LicenseCalculation{Module: "sales", Count: serviceCalc.Count, Type: model.LicenseType{Type: "addon"}, Discount: exp_discount})
	}

	var licenseCalcs []model.LicenseCalculation
	licenseCalcs = append(licenseCalcs, salesCalc, serviceCalc)
	licenseService := services.NewLicenseCalcService(licenseCalcs)

	templ.Handler(licensecalculator.CalculationResult(licenseService)).ServeHTTP(w, r)
}

func clearPriceHandler(w http.ResponseWriter, r *http.Request) {
	//Simple Clear Handler - don't need code
}

func calculateRequestHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	inputCount := len(r.Form.Get("einstein_input"))
	outputCount := len(r.Form.Get("einstein_output"))
	choosenModel := r.Form.Get("model")

	if outputCount == 0 {
		outputCount = inputCount
	}

	var requestCalc model.RequestCalculation
	requestCalc.EinsteinInput = inputCount
	requestCalc.EinsteintOutput = outputCount
	requestCalc.Model = *model.NewLLM(choosenModel)

	var requestService services.RequestCalculationService
	requestService.Request = requestCalc

	templ.Handler(requestcalculator.RequestResult(requestService)).ServeHTTP(w, r)
}

func clearRequestHandler(w http.ResponseWriter, r *http.Request) {
	//Simple Clear Handler - don't need code
}
