package routes

import (
	"fmt"
	app "github.com/mlplabs/app-utils"
	"net/http"
)

func RegisterReportsHandlers(routeItems app.Routes, wHandlers *WrapHttpHandlers) app.Routes {

	routeItems = append(routeItems, app.Route{
		Name:          "GetRemainingProducts",
		Method:        "GET",
		Pattern:       "/reports/remaining",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.GetRemainingProducts,
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

	data, err := wh.Storage.GetRemainingProducts()
	if err != nil {
		app.Log.Warning.Printf("data fetch error, %v", err)
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
	}

	app.ResponseJSON(w, http.StatusOK, data)
}
