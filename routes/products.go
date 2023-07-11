package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	app "github.com/mlplabs/app-utils"
	"github.com/mlplabs/microwms-core/core"
	"github.com/mlplabs/microwms-core/whs"
	"io/ioutil"
	"net/http"
	"strconv"
)

type GetProductsResponse struct {
	Header struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
		Count  int `json:"count"`
	} `json:"header"`
	Data []whs.Product `json:"data"`
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
		Name:          "GetProduct",
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
		Name:          "DeleteProduct",
		Method:        "DELETE",
		Pattern:       "/products/{id}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.DeleteProduct,
	})

	routeItems = append(routeItems, app.Route{
		Name:          "GetZonesOfWarehouse",
		Method:        "GET",
		Pattern:       "/warehouses/{whs_id}/zones/",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.GetZonesOfWarehouse,
	})
	routeItems = append(routeItems, app.Route{
		Name:          "GetSuggestionProducts",
		Method:        "GET",
		Pattern:       "/suggestion/products/{text}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.GetSuggestionProducts,
	})

	return routeItems
}

func (wh *WrapHttpHandlers) GetProducts(w http.ResponseWriter, r *http.Request) {
	varOffset := r.URL.Query().Get("o")
	varLimit := r.URL.Query().Get("l")

	offset, err := strconv.Atoi(varOffset)
	if err != nil {
		offset = 0
	}
	limit, err := strconv.Atoi(varLimit)
	if err != nil {
		limit = 0
	}

	products, count, err := wh.Storage.GetProductsItems(offset, limit, 0)
	if err != nil {
		app.Log.Warning.Printf("data fetch error, %v", err)
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
	}

	response := GetProductsResponse{}
	response.Data = make([]whs.Product, 0)

	response.Header.Limit = limit
	response.Header.Offset = offset
	response.Header.Count = count
	response.Data = products

	app.ResponseJSON(w, http.StatusOK, response)
}

func (wh *WrapHttpHandlers) GetProductsByBarcode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	val := vars["barcode"]

	prod, err := wh.Storage.FindProductsByBarcode(val)
	if err != nil {
		app.ResponseERROR(w, http.StatusNotFound, fmt.Errorf("product not found"))
		return
	}
	app.ResponseJSON(w, http.StatusOK, prod)
}

func (wh *WrapHttpHandlers) GetProductById(w http.ResponseWriter, r *http.Request) {
	var err error
	var prod *whs.Product
	vars := mux.Vars(r)

	if v, ok := vars["id"]; !ok || v == "0" {
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("invalid path params"))
		return
	}

	valId, err := strconv.Atoi(vars["id"])
	if err != nil {
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("invalid query params"))
		return
	}
	prod, err = wh.Storage.FindProductById(int64(valId))
	if err != nil {
		app.ResponseERROR(w, http.StatusNotFound, fmt.Errorf("product not found"))
		return
	}
	app.ResponseJSON(w, http.StatusOK, prod)
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

	whs, err := wh.Storage.FindWhsById(int64(valId))
	if err != nil {
		app.Log.Warning.Printf("data fetch error, %v", err)
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
		return
	}

	whss, err := wh.Storage.GetWhsZones(whs)
	if err != nil {
		app.Log.Warning.Printf("data fetch error, %v", err)
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

	data := new(whs.Product)
	err = json.Unmarshal(body, data)
	if err != nil {
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("can't unmarshal body, %s", err))
		return
	}

	whss, err := wh.Storage.CreateProduct(data)
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

	data := new(whs.Product)
	err = json.Unmarshal(body, data)
	if err != nil {
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("can't unmarshal body, %s", err))
		return
	}

	_, err = wh.Storage.FindProductById(data.Id)
	if err != nil {
		app.ResponseERROR(w, http.StatusNotFound, fmt.Errorf("product not found"))
		return
	}

	resultData, err := wh.Storage.UpdateProduct(data)
	if err != nil {
		if e, ok := err.(*core.WrapError); !ok {
			fmt.Println(e)
		} else {
			fmt.Println(e)
		}

		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
		return
	}
	app.ResponseJSON(w, http.StatusOK, resultData)
}

func (wh *WrapHttpHandlers) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if v, ok := vars["id"]; !ok || v == "0" {
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("invalid path params"))
		return
	}

	valId, err := strconv.Atoi(vars["id"])
	if err != nil {
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("invalid query params"))
		return
	}

	p, err := wh.Storage.FindProductById(int64(valId))
	if err != nil {
		app.ResponseERROR(w, http.StatusNotFound, fmt.Errorf("product not found"))
		return
	}
	resultData, err := wh.Storage.DeleteProduct(p)
	if err != nil {
		app.Log.Warning.Printf("item deleting error, %v", err)
		app.ResponseERROR(w, http.StatusInternalServerError, fmt.Errorf("item deleting error"))
		return
	}
	app.ResponseJSON(w, http.StatusOK, resultData)
}

func (wh *WrapHttpHandlers) GetSuggestionProducts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if _, ok := vars["text"]; !ok {
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("invalid path params"))
		return
	}

	data, err := wh.Storage.GetProductsSuggestion(vars["text"], 10)
	if err != nil {
		app.Log.Warning.Printf("data fetch error, %v", err)
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
		return
	}
	app.ResponseJSON(w, http.StatusOK, data)
}
