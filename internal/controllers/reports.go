package controllers

import (
	app "github.com/mlplabs/app-utils"
	"github.com/mlplabs/app-utils/pkg/http/response"
	"github.com/mlplabs/microwms/internal/controllers/query"
	"github.com/mlplabs/microwms/internal/usecase/service"
	"net/http"
)

type ReportsController struct {
	service *service.WhsService
}

func NewReportsController(service *service.WhsService) *ReportsController {
	return &ReportsController{
		service: service,
	}
}
func (c *ReportsController) RegisterHandlers(routeItems app.Routes) app.Routes {
	wrap := response.NewWrapper()
	routeItems = append(routeItems, app.Route{
		Name:          "GetStock",
		Method:        "GET",
		Pattern:       "/reports/stock",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.DataPages(c.getStockProducts),
	})
	return routeItems
}

func (c *ReportsController) getStockProducts(r *http.Request) (interface{}, *response.DataRange, error) {
	ctx := r.Context()
	offset, limit := query.GetOffsetLimit(r)
	//items, count, err := c.service.GetUsers(ctx, offset, limit)
	//if err != nil {
	//	return items, &response.DataRange{}, err
	//}
	//c.service.GetProducts()
	data, _ := c.service.GetStockData(ctx, 0)
	return data, &response.DataRange{
		Count:  int(0),
		Limit:  limit,
		Offset: offset,
	}, nil
}

//
////
////import (
////	"fmt"
////	"github.com/gorilla/mux"
////	app "github.com/mlplabs/app-utils"
////	"github.com/mlplabs/mwms-core/whs"
////	"net/http"
////	"strconv"
////)
////
////type RemainingParams struct {
////	Products []int64 `json:"products"`
////}
////
////func RegisterReportsHandlers(routeItems app.Routes, wHandlers *WrapHttpHandlers) app.Routes {
////
////	routeItems = append(routeItems, app.Route{
////		Name:          "GetHistoryProducts",
////		Method:        "GET",
////		Pattern:       "/reports/history/item/{id}",
////		SetHeaderJSON: true,
////		ValidateToken: false,
////		HandlerFunc:   wHandlers.GetHistoryProduct,
////	})
////
////	routeItems = append(routeItems, app.Route{
////		Name:          "GetRemainingProduct",
////		Method:        "GET",
////		Pattern:       "/reports/remaining/item/{id}",
////		SetHeaderJSON: true,
////		ValidateToken: false,
////		HandlerFunc:   wHandlers.GetRemainingProduct,
////	})
////
////	return routeItems
////}
////
////func (wh *WrapHttpHandlers) GetHistoryProduct(w http.ResponseWriter, r *http.Request) {
////	var data []whs.RemainingProductRow
////
////	vars := mux.Vars(r)
////	paramItemId := vars["id"]
////	itemId, err := strconv.Atoi(paramItemId)
////	if err != nil {
////		app.Log.Warning.Printf("failed converting item id = %v, %v", paramItemId, err)
////		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("invalid request"))
////		return
////	}
////
////	if itemId != 0 {
////		params := RemainingParams{}
////		params.Products = append(params.Products, int64(itemId))
////
////		data, err = wh.Storage.GetRemainingProductsByIds(params.Products)
////		if err != nil {
////			app.Log.Warning.Printf("data fetch error, %v", err)
////			app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
////			return
////		}
////	} else {
////		data, err = wh.Storage.GetRemainingProducts()
////		if err != nil {
////			app.Log.Warning.Printf("data fetch error, %v", err)
////			app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
////			return
////		}
////	}
////
////	app.ResponseJSON(w, http.StatusOK, data)
////}
////
////func (wh *WrapHttpHandlers) GetRemainingProduct(w http.ResponseWriter, r *http.Request) {
////	var data []whs.RemainingProductRow
////
////	vars := mux.Vars(r)
////	paramItemId := vars["id"]
////	itemId, err := strconv.Atoi(paramItemId)
////	if err != nil {
////		app.Log.Warning.Printf("failed converting item id = %v, %v", paramItemId, err)
////		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("invalid request"))
////		return
////	}
////
////	if itemId != 0 {
////		params := RemainingParams{}
////		params.Products = append(params.Products, int64(itemId))
////
////		data, err = wh.Storage.GetRemainingProductsByIds(params.Products)
////		if err != nil {
////			app.Log.Warning.Printf("data fetch error, %v", err)
////			app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
////			return
////		}
////	} else {
////		data, err = wh.Storage.GetRemainingProducts()
////		if err != nil {
////			app.Log.Warning.Printf("data fetch error, %v", err)
////			app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
////			return
////		}
////	}
////
////	app.ResponseJSON(w, http.StatusOK, data)
////}
