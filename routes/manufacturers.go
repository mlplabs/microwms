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

type GetManufacturersResponse struct {
	Header struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
		Count  int `json:"count"`
	} `json:"header"`
	Data []whs.Manufacturer `json:"data"`
}

func RegisterManufacturersHandlers(routeItems app.Routes, wHandlers *WrapHttpHandlers) app.Routes {

	routeItems = append(routeItems, app.Route{
		Name:          "GetManufacturers",
		Method:        "GET",
		Pattern:       "/manufacturers",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.GetManufacturers,
	})
	routeItems = append(routeItems, app.Route{
		Name:          "GetManufacturer",
		Method:        "GET",
		Pattern:       "/manufacturers/{id}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.GetManufacturerById,
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
		Name:          "DeleteManufacturer",
		Method:        "DELETE",
		Pattern:       "/manufacturers/{id}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.DeleteManufacturer,
	})
	routeItems = append(routeItems, app.Route{
		Name:          "GetSuggestionProducts",
		Method:        "GET",
		Pattern:       "/suggestion/manufacturers/{text}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.GetSuggestionManufacturers,
	})

	return routeItems
}

func (wh *WrapHttpHandlers) CreateManufacturer(w http.ResponseWriter, r *http.Request) {
	// читаем данные
	body, err := ioutil.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("can't read body"))
		return
	}

	data := new(whs.Manufacturer)
	err = json.Unmarshal(body, data)
	if err != nil {
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("can't unmarshal body, %s", err))
		return
	}

	mnf, err := wh.Storage.CreateManufacturer(data)
	if err != nil {
		if err, ok := err.(*core.WrapError); !ok {
			fmt.Println(err)
		} else {
			fmt.Println(err)
		}

		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
		return
	}

	app.ResponseJSON(w, http.StatusOK, mnf)
}

func (wh *WrapHttpHandlers) GetManufacturerById(w http.ResponseWriter, r *http.Request) {
	var err error
	var mnf *whs.Manufacturer
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
	mnf, err = wh.Storage.FindManufacturerById(int64(valId))
	if err != nil {
		app.ResponseERROR(w, http.StatusNotFound, fmt.Errorf("product not found"))
		return
	}
	app.ResponseJSON(w, http.StatusOK, mnf)
}

func (wh *WrapHttpHandlers) GetManufacturers(w http.ResponseWriter, r *http.Request) {
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

	mnfs, count, err := wh.Storage.GetManufacturersItems(offset, limit, 0)
	if err != nil {
		app.Log.Warning.Printf("data fetch error, %v", err)
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
	}

	response := GetManufacturersResponse{}
	response.Data = make([]whs.Manufacturer, 0)

	response.Header.Limit = limit
	response.Header.Offset = offset
	response.Header.Count = count
	response.Data = mnfs

	app.ResponseJSON(w, http.StatusOK, response)
}

func (wh *WrapHttpHandlers) UpdateManufacturer(w http.ResponseWriter, r *http.Request) {
	// читаем данные
	body, err := ioutil.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("can't read body"))
		return
	}

	data := new(whs.Manufacturer)
	err = json.Unmarshal(body, data)
	if err != nil {
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("can't unmarshal body, %s", err))
		return
	}

	_, err = wh.Storage.FindManufacturerById(data.Id)
	if err != nil {
		app.ResponseERROR(w, http.StatusNotFound, fmt.Errorf("product not found"))
		return
	}

	resultData, err := wh.Storage.UpdateManufacturer(data)
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

func (wh *WrapHttpHandlers) DeleteManufacturer(w http.ResponseWriter, r *http.Request) {
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

	m, err := wh.Storage.FindManufacturerById(int64(valId))
	if err != nil {
		app.ResponseERROR(w, http.StatusNotFound, fmt.Errorf("product not found"))
		return
	}
	resultData, err := wh.Storage.DeleteManufacturer(m)
	if err != nil {
		app.ResponseERROR(w, http.StatusInternalServerError, fmt.Errorf("item deleting error"))
		return
	}
	app.ResponseJSON(w, http.StatusOK, resultData)
}

func (wh *WrapHttpHandlers) GetSuggestionManufacturers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if _, ok := vars["text"]; !ok {
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("invalid path params"))
		return
	}

	data, err := wh.Storage.GetManufacturersSuggestion(vars["text"], 10)
	if err != nil {
		app.Log.Warning.Printf("data fetch error, %v", err)
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
		return
	}

	app.ResponseJSON(w, http.StatusOK, data)
}
