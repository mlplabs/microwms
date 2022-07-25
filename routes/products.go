package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	l "github.com/mikelpsv/mod_logging"
	app "github.com/mikelpsv/mod_micro_app"
	"github.com/mlplabs/microwms-core/core"
	"github.com/mlplabs/microwms-core/models"
	"io/ioutil"
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
		Name:          "GetProducts",
		Method:        "GET",
		Pattern:       "/products/{id}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.GetProductById,
	})
	routeItems = append(routeItems, app.Route{
		Name:          "GetProductByBarcode",
		Method:        "GET",
		Pattern:       "/products/barcode/{barcode}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.GetProductsByBarcode,
	})
	routeItems = append(routeItems, app.Route{
		Name:          "CreateProduct",
		Method:        "POST",
		Pattern:       "/products",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.CreateProduct,
	})
	routeItems = append(routeItems, app.Route{
		Name:          "UpdateProduct",
		Method:        "PUT",
		Pattern:       "/products",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.UpdateProduct,
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
		Name:          "CreateManufacturer",
		Method:        "POST",
		Pattern:       "/manufacturers",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.CreateManufacturer,
	})
	routeItems = append(routeItems, app.Route{
		Name:          "UpdateManufacturer",
		Method:        "PUT",
		Pattern:       "/manufacturers",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.UpdateManufacturer,
	})
	routeItems = append(routeItems, app.Route{
		Name:          "GetWarehouses",
		Method:        "GET",
		Pattern:       "/warehouses",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.GetWarehouses,
	})
	routeItems = append(routeItems, app.Route{
		Name:          "GetWarehouse",
		Method:        "GET",
		Pattern:       "/warehouses/{whs_id}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.GetWarehouse,
	})
	routeItems = append(routeItems, app.Route{
		Name:          "GetZonesOfWarehouse",
		Method:        "GET",
		Pattern:       "/warehouses/{whs_id}/zones/",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.GetZonesOfWarehouse,
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

func (wh *WrapHttpHandlers) GetProductsByBarcode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	val := vars["barcode"]

	prodService := wh.Storage.GetProductService()
	prod, err := prodService.FindProductsByBarcode(val)
	if err != nil {
		app.ResponseERROR(w, http.StatusNotFound, fmt.Errorf("product not found"))
		return
	}
	app.ResponseJSON(w, http.StatusOK, prod)
}

func (wh *WrapHttpHandlers) GetProductById(w http.ResponseWriter, r *http.Request) {
	var err error
	var prod *models.Product
	vars := mux.Vars(r)

	if v, ok := vars["id"]; !ok || v == "0" {
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("invalid path params"))
		return
	}

	prodService := wh.Storage.GetProductService()

	valId, err := strconv.Atoi(vars["id"])
	if err != nil {
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("invalid query params"))
		return
	}
	prod, err = prodService.FindProductById(int64(valId))
	if err != nil {
		app.ResponseERROR(w, http.StatusNotFound, fmt.Errorf("product not found"))
		return
	}
	app.ResponseJSON(w, http.StatusOK, prod)
}

func (wh *WrapHttpHandlers) GetWarehouses(w http.ResponseWriter, r *http.Request) {
	whsService := wh.Storage.GetWhsService()
	whss, err := whsService.GetWarehouses()
	if err != nil {
		l.Warning.Printf("data fetch error, %v", err)
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
		return
	}

	app.ResponseJSON(w, http.StatusOK, whss)
}

func (wh *WrapHttpHandlers) GetWarehouse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if _, ok := vars["whs_id"]; !ok {
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("invalid path params"))
		return
	}

	valId, err := strconv.Atoi(vars["whs_id"])
	if err != nil {
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("invalid query params"))
		return
	}

	whsService := wh.Storage.GetWhsService()
	whss, err := whsService.FindWhsById(int64(valId))
	if err != nil {
		l.Warning.Printf("data fetch error, %v", err)
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
		return
	}

	app.ResponseJSON(w, http.StatusOK, whss)
}

func (wh *WrapHttpHandlers) GetZonesOfWarehouse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if _, ok := vars["whs_id"]; !ok {
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("invalid path params"))
		return
	}

	valId, err := strconv.Atoi(vars["whs_id"])
	if err != nil {
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("invalid query params"))
		return
	}

	whsService := wh.Storage.GetWhsService()
	whs, err := whsService.FindWhsById(int64(valId))
	if err != nil {
		l.Warning.Printf("data fetch error, %v", err)
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
		return
	}
	whss, err := whsService.GetZones(whs)
	if err != nil {
		l.Warning.Printf("data fetch error, %v", err)
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
		return
	}
	app.ResponseJSON(w, http.StatusOK, whss)
}

func (wh *WrapHttpHandlers) CreateProduct(w http.ResponseWriter, r *http.Request) {
	// читаем данные
	body, err := ioutil.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("can't read body"))
		return
	}

	data := new(models.Product)
	err = json.Unmarshal(body, data)
	if err != nil {
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("can't unmarshal body, %s", err))
		return
	}

	prodService := wh.Storage.GetProductService()
	whss, err := prodService.CreateProduct(data)
	if err != nil {
		if err, ok := err.(*core.WrapError); !ok {
			fmt.Println(err)
		} else {
			fmt.Println(err)
		}

		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
		return
	}

	app.ResponseJSON(w, http.StatusOK, whss)
}

func (wh *WrapHttpHandlers) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// читаем данные
	body, err := ioutil.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("can't read body"))
		return
	}

	data := new(models.Product)
	err = json.Unmarshal(body, data)
	if err != nil {
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("can't unmarshal body, %s", err))
		return
	}

	prodService := wh.Storage.GetProductService()
	whss, err := prodService.UpdateProduct(data)
	if err != nil {
		if e, ok := err.(*core.WrapError); !ok {
			fmt.Println(e)
		} else {
			fmt.Println(e)
		}

		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
		return
	}

	app.ResponseJSON(w, http.StatusOK, whss)
}

func (wh *WrapHttpHandlers) CreateManufacturer(w http.ResponseWriter, r *http.Request) {
	// читаем данные
	body, err := ioutil.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("can't read body"))
		return
	}

	data := new(models.Manufacturer)
	err = json.Unmarshal(body, data)
	if err != nil {
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("can't unmarshal body, %s", err))
		return
	}

	prodService := wh.Storage.GetProductService()
	whss, err := prodService.CreateManufacturer(data)
	if err != nil {
		if err, ok := err.(*core.WrapError); !ok {
			fmt.Println(err)
		} else {
			fmt.Println(err)
		}

		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
		return
	}

	app.ResponseJSON(w, http.StatusOK, whss)
}

func (wh *WrapHttpHandlers) UpdateManufacturer(w http.ResponseWriter, r *http.Request) {
	// читаем данные
	body, err := ioutil.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("can't read body"))
		return
	}

	data := new(models.Manufacturer)
	err = json.Unmarshal(body, data)
	if err != nil {
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("can't unmarshal body, %s", err))
		return
	}

	prodService := wh.Storage.GetProductService()
	whss, err := prodService.UpdateManufacturer(data)
	if err != nil {
		if err, ok := err.(*core.WrapError); !ok {
			fmt.Println(err)
		} else {
			fmt.Println(err)
		}

		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
		return
	}

	app.ResponseJSON(w, http.StatusOK, whss)
}
