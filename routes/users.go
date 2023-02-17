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

type GetUsersResponse struct {
	Header struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
		Count  int `json:"count"`
	} `json:"header"`
	Data []whs.User `json:"data"`
}

func RegisterUsersHandlers(routeItems app.Routes, wHandlers *WrapHttpHandlers) app.Routes {

	routeItems = append(routeItems, app.Route{
		Name:          "GetUsers",
		Method:        "GET",
		Pattern:       "/users",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.GetUsers,
	})
	routeItems = append(routeItems, app.Route{
		Name:          "GetUser",
		Method:        "GET",
		Pattern:       "/users/{id}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.GetUserById,
	})
	routeItems = append(routeItems, app.Route{
		Name:          "CreateUser",
		Method:        "POST",
		Pattern:       "/users",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.CreateUser,
	})
	routeItems = append(routeItems, app.Route{
		Name:          "UpdateUser",
		Method:        "PUT",
		Pattern:       "/users/{id}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.UpdateUser,
	})
	routeItems = append(routeItems, app.Route{
		Name:          "DeleteUser",
		Method:        "DELETE",
		Pattern:       "/users/{id}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.DeleteUser,
	})
	routeItems = append(routeItems, app.Route{
		Name:          "GetSuggestionUsers",
		Method:        "GET",
		Pattern:       "/suggestion/users/{text}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.GetSuggestionUsers,
	})

	return routeItems
}

func (wh *WrapHttpHandlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	// читаем данные
	body, err := ioutil.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("can't read body"))
		return
	}

	data := new(whs.User)
	err = json.Unmarshal(body, data)
	if err != nil {
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("can't unmarshal body, %s", err))
		return
	}

	whss, err := wh.Storage.CreateUser(data)
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

func (wh *WrapHttpHandlers) GetUserById(w http.ResponseWriter, r *http.Request) {
	var err error
	var usr *whs.User
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
	usr, err = wh.Storage.FindUserById(int64(valId))
	if err != nil {
		app.ResponseERROR(w, http.StatusNotFound, fmt.Errorf("product not found"))
		return
	}
	app.ResponseJSON(w, http.StatusOK, usr)
}

func (wh *WrapHttpHandlers) GetUsers(w http.ResponseWriter, r *http.Request) {

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

	usrs, count, err := wh.Storage.GetUsersItems(offset, limit, 0)
	if err != nil {
		app.Log.Warning.Printf("data fetch error, %v", err)
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
	}

	response := GetUsersResponse{}

	response.Data = make([]whs.User, 0)

	response.Header.Limit = limit
	response.Header.Offset = offset
	response.Header.Count = count
	response.Data = usrs

	app.ResponseJSON(w, http.StatusOK, response)
}

func (wh *WrapHttpHandlers) UpdateUser(w http.ResponseWriter, r *http.Request) {
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

	m, err := wh.Storage.FindUserById(int64(valId))
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

	data := new(whs.User)
	err = json.Unmarshal(body, data)
	if err != nil {
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("can't unmarshal body, %s", err))
		return
	}

	data.Id = m.Id

	resultData, err := wh.Storage.UpdateUser(data)
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

func (wh *WrapHttpHandlers) DeleteUser(w http.ResponseWriter, r *http.Request) {
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

	m, err := wh.Storage.FindUserById(int64(valId))
	if err != nil {
		app.ResponseERROR(w, http.StatusNotFound, fmt.Errorf("product not found"))
		return
	}
	resultData, err := wh.Storage.DeleteUser(m)
	if err != nil {
		app.ResponseERROR(w, http.StatusInternalServerError, fmt.Errorf("item deleting error"))
		return
	}
	app.ResponseJSON(w, http.StatusOK, resultData)
}

func (wh *WrapHttpHandlers) GetSuggestionUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if _, ok := vars["text"]; !ok {
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("invalid path params"))
		return
	}

	data, err := wh.Storage.GetUsersSuggestion(vars["text"], 10)
	if err != nil {
		app.Log.Warning.Printf("data fetch error, %v", err)
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
		return
	}

	app.ResponseJSON(w, http.StatusOK, data)
}
