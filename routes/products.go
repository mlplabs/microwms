package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	l "github.com/mikelpsv/mod_logging"
	app "github.com/mikelpsv/mod_micro_app"
	"github.com/mlplabs/microwms-core/models"
	"net/http"
	"strconv"
)

type WrapHttpHandlers struct {
	Storage *models.Storage
}

func RegisterProductsHandlers(routeItems app.Routes, wHandlers *WrapHttpHandlers) app.Routes {
	routeItems = append(routeItems, app.Route{
		Name:          "GetProducts",
		Method:        "GET",
		Pattern:       "/products",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.GetProducts,
	})
	routeItems = append(routeItems, app.Route{
		Name:          "GetManufacturers",
		Method:        "GET",
		Pattern:       "/manufacturers",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.GetManufacturers,
	})
	routeItems = append(routeItems, app.Route{
		Name:          "GetManufacturers",
		Method:        "GET",
		Pattern:       "/warehouses",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.GetWarehouses,
	})
	routeItems = append(routeItems, app.Route{
		Name:          "GetProductByParam",
		Method:        "GET",
		Pattern:       "/products/{param}:{value}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.GetProductByParam,
	})
	return routeItems
}

func (wh *WrapHttpHandlers) GetProducts(w http.ResponseWriter, r *http.Request) {
	prodService := wh.Storage.GetProductService()
	products, err := prodService.GetProducts()
	if err != nil {
		l.Warning.Printf("data fetch error, %v", err)
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
	}

	app.ResponseJSON(w, http.StatusOK, products)
}

func (wh *WrapHttpHandlers) GetManufacturers(w http.ResponseWriter, r *http.Request) {
	prodService := wh.Storage.GetProductService()
	mnfs, err := prodService.GetManufacturers()
	if err != nil {
		l.Warning.Printf("data fetch error, %v", err)
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
	}

	app.ResponseJSON(w, http.StatusOK, mnfs)
}

func (wh *WrapHttpHandlers) GetProductByParam(w http.ResponseWriter, r *http.Request) {
	var err error
	var prod *models.Product
	vars := mux.Vars(r)
	val := vars["value"]
	param := vars["param"]

	prodService := wh.Storage.GetProductService()

	switch param {
	case "id":
		valId := 0

		valId, err = strconv.Atoi(vars["value"])
		if err != nil {
			app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("invalid query params"))
		}
		prod, err = prodService.FindProductById(int64(valId))

	case "barcode":
		prod, err = prodService.FindProductsByBarcode(val)
	}

	app.ResponseJSON(w, http.StatusOK, prod)
}

func (wh *WrapHttpHandlers) GetWarehouses(w http.ResponseWriter, r *http.Request) {
	whsService := wh.Storage.GetWhsService()
	whss, err := whsService.GetWarehouses()
	if err != nil {
		l.Warning.Printf("data fetch error, %v", err)
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
	}

	app.ResponseJSON(w, http.StatusOK, whss)
}
