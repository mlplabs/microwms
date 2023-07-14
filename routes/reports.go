package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	app "github.com/mlplabs/app-utils"
	"github.com/mlplabs/microwms-core/whs"
	"io"
	"net/http"
	"strconv"
)

type RemainingParams struct {
	Products []int64 `json:"products"`
}

func RegisterReportsHandlers(routeItems app.Routes, wHandlers *WrapHttpHandlers) app.Routes {

	routeItems = append(routeItems, app.Route{
		Name:          "GetRemainingProducts",
		Method:        "GET",
		Pattern:       "/reports/remaining",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.GetRemainingProducts,
	})

	routeItems = append(routeItems, app.Route{
		Name:          "GetRemainingProduct",
		Method:        "GET",
		Pattern:       "/reports/remaining/item/{id}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.GetRemainingProduct,
	})

	return routeItems
}

func (wh *WrapHttpHandlers) GetRemainingProducts(w http.ResponseWriter, r *http.Request) {

	//varOffset := r.URL.Query().Get("o")
	//varLimit := r.URL.Query().Get("l")
	//
	//offset, err := strconv.Atoi(varOffset)
	//if err != nil {
	//	offset = 0
	//}
	//limit, err := strconv.Atoi(varLimit)
	//if err != nil {
	//	limit = 0
	//}
	var data []whs.RemainingProductRow

	params := RemainingParams{}

	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		//errText := fmt.Sprintf("can't read body, %v", err)
		//models.HttpError(w, errText, http.StatusBadRequest)
		return
	}
	json.Unmarshal(body, &params)
	if len(params.Products) > 0 {
		data, err = wh.Storage.GetRemainingProductsByIds(params.Products)
		if err != nil {
			app.Log.Warning.Printf("data fetch error, %v", err)
			app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
		}
	} else {
		data, err = wh.Storage.GetRemainingProducts()
		if err != nil {
			app.Log.Warning.Printf("data fetch error, %v", err)
			app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
		}
	}

	app.ResponseJSON(w, http.StatusOK, data)
}

func (wh *WrapHttpHandlers) GetRemainingProduct(w http.ResponseWriter, r *http.Request) {
	var data []whs.RemainingProductRow

	vars := mux.Vars(r)
	paramItemId := vars["id"]
	itemId, err := strconv.Atoi(paramItemId)
	if err != nil {
		app.Log.Warning.Printf("failed converting item id = %v, %v", paramItemId, err)
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("invalid request"))
		return
	}

	if itemId != 0 {
		params := RemainingParams{}
		params.Products = append(params.Products, int64(itemId))

		data, err = wh.Storage.GetRemainingProductsByIds(params.Products)
		if err != nil {
			app.Log.Warning.Printf("data fetch error, %v", err)
			app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
			return
		}
	} else {
		data, err = wh.Storage.GetRemainingProducts()
		if err != nil {
			app.Log.Warning.Printf("data fetch error, %v", err)
			app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
			return
		}
	}

	app.ResponseJSON(w, http.StatusOK, data)
}
