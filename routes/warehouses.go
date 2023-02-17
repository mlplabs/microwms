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

type GetWhsResponse struct {
	Header struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
		Count  int `json:"count"`
	} `json:"header"`
	Data []whs.Whs `json:"data"`
}

func RegisterWhsHandlers(routeItems app.Routes, wHandlers *WrapHttpHandlers) app.Routes {

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
		Name:          "CreateWhs",
		Method:        "POST",
		Pattern:       "/warehouses",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.CreateWhs,
	})
	routeItems = append(routeItems, app.Route{
		Name:          "UpdateWhs",
		Method:        "PUT",
		Pattern:       "/warehouses/{id}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.UpdateWhs,
	})
	routeItems = append(routeItems, app.Route{
		Name:          "DeleteWhs",
		Method:        "DELETE",
		Pattern:       "/warehouses/{id}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.DeleteWhs,
	})

	return routeItems
}

func (wh *WrapHttpHandlers) GetWarehouses(w http.ResponseWriter, r *http.Request) {
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

	mnfs, count, err := wh.Storage.GetWhsItems(offset, limit, 0)
	if err != nil {
		app.Log.Warning.Printf("data fetch error, %v", err)
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
	}

	response := GetWhsResponse{}
	response.Data = make([]whs.Whs, 0)

	response.Header.Limit = limit
	response.Header.Offset = offset
	response.Header.Count = count
	response.Data = mnfs

	app.ResponseJSON(w, http.StatusOK, response)
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

	whss, err := wh.Storage.GetWhsById(int64(valId))
	if err != nil {
		app.Log.Warning.Printf("data fetch error, %v", err)
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
		return
	}

	app.ResponseJSON(w, http.StatusOK, whss)
}

func (wh *WrapHttpHandlers) CreateWhs(w http.ResponseWriter, r *http.Request) {
	// читаем данные
	body, err := ioutil.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("can't read body"))
		return
	}

	data := new(whs.Whs)
	err = json.Unmarshal(body, data)
	if err != nil {
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("can't unmarshal body, %s", err))
		return
	}

	whss, err := wh.Storage.CreateWhs(data)
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

func (wh *WrapHttpHandlers) UpdateWhs(w http.ResponseWriter, r *http.Request) {
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

	m, err := wh.Storage.FindWhsById(int64(valId))
	if err != nil {
		app.ResponseERROR(w, http.StatusNotFound, fmt.Errorf("product not found"))
		return
	}

	// читаем данные
	body, err := ioutil.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("can't read body"))
		return
	}

	data := new(whs.Whs)
	err = json.Unmarshal(body, data)
	if err != nil {
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("can't unmarshal body, %s", err))
		return
	}

	data.Id = m.Id

	resultData, err := wh.Storage.UpdateWhs(data)
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

func (wh *WrapHttpHandlers) DeleteWhs(w http.ResponseWriter, r *http.Request) {
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

	m, err := wh.Storage.FindWhsById(int64(valId))
	if err != nil {
		app.ResponseERROR(w, http.StatusNotFound, fmt.Errorf("product not found"))
		return
	}
	resultData, err := wh.Storage.DeleteWhs(m)
	if err != nil {
		app.ResponseERROR(w, http.StatusInternalServerError, fmt.Errorf("item deleting error"))
		return
	}
	app.ResponseJSON(w, http.StatusOK, resultData)
}
