package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	app "github.com/mlplabs/app-utils"
	"github.com/mlplabs/app-utils/pkg/http/errors"
	"github.com/mlplabs/app-utils/pkg/http/response"
	"github.com/mlplabs/microwms/internal/controllers/query"
	"github.com/mlplabs/microwms/internal/usecase/service"
	"github.com/mlplabs/mwms-core/whs/model"
	"io"
	"net/http"
)

type WarehouseController struct {
	service *service.WhsService
}

func NewWarehouseController(service *service.WhsService) *WarehouseController {
	return &WarehouseController{
		service: service,
	}
}

func (c *WarehouseController) RegisterHandlers(routeItems app.Routes) app.Routes {
	wrap := response.NewWrapper()
	routeItems = append(routeItems, app.Route{
		Name:          "GetWarehouse",
		Method:        "GET",
		Pattern:       "/warehouses/{itemId}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.DataPlain(c.getById),
	})
	routeItems = append(routeItems, app.Route{
		Name:          "GetWarehouse",
		Method:        "GET",
		Pattern:       "/warehouses",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.DataPages(c.getWarehouses),
	})
	routeItems = append(routeItems, app.Route{
		Name:          "CreateWarehouse",
		Method:        "POST",
		Pattern:       "/warehouses",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.DataPlain(c.createWarehouse),
	})
	routeItems = append(routeItems, app.Route{
		Name:          "UpdateWarehouse",
		Method:        "PUT",
		Pattern:       "/warehouses/{itemId}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.DataPlain(c.updateWarehouse),
	})
	routeItems = append(routeItems, app.Route{
		Name:          "DeleteWarehouse",
		Method:        "DELETE",
		Pattern:       "/warehouses/{itemId}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.Empty(c.deleteWarehouse),
	})
	routeItems = append(routeItems, app.Route{
		Name:          "GetSuggestionUsers",
		Method:        "GET",
		Pattern:       "/suggestion/warehouses/{text}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.DataList(c.getSuggestion),
	})
	return routeItems
}

// @Summary get warehouse by ID
// @Tags warehouses
// @Success		200	{object} response.PlainData{data=Warehouse}
// @Param itemID path integer true "warehouse ID"
// @Router /warehouses/{itemID} [get]
func (c *WarehouseController) getById(r *http.Request) (interface{}, error) {
	var err error
	itemId, err := query.GetItemId(r)
	if err != nil {
		return nil, err
	}
	item, err := c.service.GetWarehouseById(r.Context(), int64(itemId))
	if err != nil {
		return item, errors.NewServerError(err)
	}

	return item, nil
}

// @Summary get warehouses list
// @Tags warehouses
// @Success		200	{object} response.Pagination{data=[]Warehouse}
// @Param o query integer false "offset"
// @Param l query integer false "limit"
// @Router /warehouses [get]
func (c *WarehouseController) getWarehouses(r *http.Request) (interface{}, *response.DataRange, error) {
	ctx := r.Context()
	offset, limit := query.GetOffsetLimit(r)
	items, count, err := c.service.GetWarehouses(ctx, offset, limit)
	if err != nil {
		return items, &response.DataRange{}, err
	}

	return items, &response.DataRange{
		Count:  int(count),
		Limit:  limit,
		Offset: offset,
	}, nil
}

// @Summary create warehouse
// @Tags warehouses
// @Param request body Warehouse true "warehouse data"
// @Success		200	{object} response.PlainData{data=integer}
// @Router /warehouses [post]
func (c *WarehouseController) createWarehouse(r *http.Request) (interface{}, error) {
	ctx := r.Context()
	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		return nil, errors.NewInvalidInputData(fmt.Errorf("can't read body"))
	}

	data := new(model.Warehouse)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, errors.NewInvalidInputData(fmt.Errorf("can't unmarshal body, %s", err))
	}

	itemId, err := c.service.CreateWarehouse(ctx, data)
	if err != nil {
		return nil, errors.NewBadRequest(fmt.Errorf("data fetch error, %s", err))
	}
	return itemId, nil
}

// @Summary update warehouse
// @Tags warehouses
// @Param itemID path integer true "warehouse ID"
// @Param request body Warehouse true "warehouse data"
// @Success		200	{object} response.PlainData{data=integer}
// @Router /warehouses/{itemID} [put]
func (c *WarehouseController) updateWarehouse(r *http.Request) (interface{}, error) {
	ctx := r.Context()
	itemReqId, err := query.GetItemId(r)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		return nil, errors.NewInvalidInputData(fmt.Errorf("can't read body"))
	}

	data := new(model.Warehouse)
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, errors.NewInvalidInputData(fmt.Errorf("can't unmarshal body, %s", err))
	}
	data.Id = int64(itemReqId)
	itemId, err := c.service.UpdateWarehouse(ctx, data)
	if err != nil {
		return nil, errors.NewBadRequest(fmt.Errorf("data fetch error, %s", err))
	}
	return itemId, nil
}

// @Summary delete warehouse by ID
// @Tags warehouses
// @Success		200	{object} map[string]interface{}
// @Param itemID path integer true "warehouse ID"
// @Router /warehouses/{itemID} [delete]
func (c *WarehouseController) deleteWarehouse(r *http.Request) error {
	itemId, err := query.GetItemId(r)
	if err != nil {
		return err
	}

	err = c.service.DeleteWarehouse(r.Context(), int64(itemId))
	if err != nil {
		return errors.NewServerError(fmt.Errorf("data fetch error, %s", err))
	}
	return nil
}

// @Summary get warehouses suggestion
// @Tags warehouses
// @Success		200	{object} response.List{data=[]Warehouse}
// @Param text path string true "text for suggestion"
// @Router /suggestion/warehouses/{text} [get]
func (c *WarehouseController) getSuggestion(r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	if _, ok := vars["text"]; !ok {
		return nil, errors.NewInvalidInputData(fmt.Errorf("invalid path params"))
	}

	data, err := c.service.GetWarehouseSuggestion(r.Context(), vars["text"], 10)
	if err != nil {
		return nil, errors.NewServerError(err)
	}
	return data, nil
}

//
//import (
//	"encoding/json"
//	"fmt"
//	"github.com/gorilla/mux"
//	app "github.com/mlplabs/app-utils"
//	"github.com/mlplabs/mwms-core/core"
//	"github.com/mlplabs/mwms-core/whs"
//	"io/ioutil"
//	"net/http"
//	"strconv"
//)
//
//type GetWhsResponse struct {
//	Header struct {
//		Limit  int `json:"limit"`
//		Offset int `json:"offset"`
//		Count  int `json:"count"`
//	} `json:"header"`
//	Data []whs.Whs `json:"data"`
//}
//
//func RegisterWhsHandlers(routeItems app.Routes, wHandlers *WrapHttpHandlers) app.Routes {
//
//	routeItems = append(routeItems, app.Route{
//		Name:          "GetWarehouses",
//		Method:        "GET",
//		Pattern:       "/warehouses",
//		SetHeaderJSON: true,
//		ValidateToken: false,
//		HandlerFunc:   wHandlers.GetWarehouses,
//	})
//	routeItems = append(routeItems, app.Route{
//		Name:          "GetWarehouse",
//		Method:        "GET",
//		Pattern:       "/warehouses/{whs_id}",
//		SetHeaderJSON: true,
//		ValidateToken: false,
//		HandlerFunc:   wHandlers.GetWarehouse,
//	})
//	routeItems = append(routeItems, app.Route{
//		Name:          "CreateWhs",
//		Method:        "POST",
//		Pattern:       "/warehouses",
//		SetHeaderJSON: true,
//		ValidateToken: false,
//		HandlerFunc:   wHandlers.CreateWhs,
//	})
//	routeItems = append(routeItems, app.Route{
//		Name:          "UpdateWhs",
//		Method:        "PUT",
//		Pattern:       "/warehouses",
//		SetHeaderJSON: true,
//		ValidateToken: false,
//		HandlerFunc:   wHandlers.UpdateWhs,
//	})
//	routeItems = append(routeItems, app.Route{
//		Name:          "DeleteWhs",
//		Method:        "DELETE",
//		Pattern:       "/warehouses/{id}",
//		SetHeaderJSON: true,
//		ValidateToken: false,
//		HandlerFunc:   wHandlers.DeleteWhs,
//	})
//
//	return routeItems
//}
//
//func (wh *WrapHttpHandlers) GetWarehouses(w http.ResponseWriter, r *http.Request) {
//	varOffset := r.URL.Query().Get("o")
//	varLimit := r.URL.Query().Get("l")
//
//	offset, err := strconv.Atoi(varOffset)
//	if err != nil {
//		offset = 0
//	}
//	limit, err := strconv.Atoi(varLimit)
//	if err != nil {
//		limit = 0
//	}
//
//	mnfs, count, err := wh.Storage.GetWhsItems(offset, limit, 0)
//	if err != nil {
//		app.Log.Warning.Printf("data fetch error, %v", err)
//		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
//	}
//
//	response := GetWhsResponse{}
//	response.Data = make([]whs.Whs, 0)
//
//	response.Header.Limit = limit
//	response.Header.Offset = offset
//	response.Header.Count = count
//	response.Data = mnfs
//
//	app.ResponseJSON(w, http.StatusOK, response)
//}
//
//func (wh *WrapHttpHandlers) GetWarehouse(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//
//	if _, ok := vars["whs_id"]; !ok {
//		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("invalid path params"))
//		return
//	}
//
//	valId, err := strconv.Atoi(vars["whs_id"])
//	if err != nil {
//		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("invalid query params"))
//		return
//	}
//
//	whss, err := wh.Storage.GetWhsById(int64(valId))
//	if err != nil {
//		app.Log.Warning.Printf("data fetch error, %v", err)
//		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
//		return
//	}
//
//	app.ResponseJSON(w, http.StatusOK, whss)
//}
//
//func (wh *WrapHttpHandlers) CreateWhs(w http.ResponseWriter, r *http.Request) {
//	// читаем данные
//	body, err := ioutil.ReadAll(r.Body)
//	if err != nil || len(body) == 0 {
//		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("can't read body"))
//		return
//	}
//
//	data := new(whs.Whs)
//	err = json.Unmarshal(body, data)
//	if err != nil {
//		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("can't unmarshal body, %s", err))
//		return
//	}
//
//	whss, err := wh.Storage.CreateWhs(data)
//	if err != nil {
//		if err, ok := err.(*core.WrapError); !ok {
//			fmt.Println(err)
//		} else {
//			fmt.Println(err)
//		}
//
//		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
//		return
//	}
//
//	app.ResponseJSON(w, http.StatusOK, whss)
//}
//
//func (wh *WrapHttpHandlers) UpdateWhs(w http.ResponseWriter, r *http.Request) {
//	// читаем данные
//	body, err := ioutil.ReadAll(r.Body)
//	if err != nil || len(body) == 0 {
//		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("can't read body"))
//		return
//	}
//
//	data := new(whs.Whs)
//	err = json.Unmarshal(body, data)
//	if err != nil {
//		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("can't unmarshal body, %s", err))
//		return
//	}
//
//	_, err = wh.Storage.FindWhsById(data.Id)
//	if err != nil {
//		app.ResponseERROR(w, http.StatusNotFound, fmt.Errorf("warehouse not found"))
//		return
//	}
//
//	resultData, err := wh.Storage.UpdateWhs(data)
//	if err != nil {
//		if e, ok := err.(*core.WrapError); !ok {
//			fmt.Println(e)
//		} else {
//			fmt.Println(e)
//		}
//
//		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
//		return
//	}
//	app.ResponseJSON(w, http.StatusOK, resultData)
//}
//
//func (wh *WrapHttpHandlers) DeleteWhs(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//
//	if v, ok := vars["id"]; !ok || v == "0" {
//		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("invalid path params"))
//		return
//	}
//
//	valId, err := strconv.Atoi(vars["id"])
//	if err != nil {
//		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("invalid query params"))
//		return
//	}
//
//	m, err := wh.Storage.FindWhsById(int64(valId))
//	if err != nil {
//		app.ResponseERROR(w, http.StatusNotFound, fmt.Errorf("product not found"))
//		return
//	}
//	resultData, err := wh.Storage.DeleteWhs(m)
//	if err != nil {
//		app.ResponseERROR(w, http.StatusInternalServerError, fmt.Errorf("item deleting error"))
//		return
//	}
//	app.ResponseJSON(w, http.StatusOK, resultData)
//}
