package routes

import (
	app "github.com/mlplabs/app-utils"
	"net/http"
)

type PingResponse struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
	Version     string `json:"version"`
}

type HealthResponse struct {
	Database bool `json:"database"`
}

func RegisterControlHandlers(routeItems app.Routes) app.Routes {
	routeItems = append(routeItems, app.Route{
		Name:          "Ping",
		Method:        "GET",
		Pattern:       "/ping",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   Ping,
	})
	routeItems = append(routeItems, app.Route{
		Name:          "health",
		Method:        "GET",
		Pattern:       "/health",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   GetHealth,
	})
	return routeItems
}

func Ping(w http.ResponseWriter, r *http.Request) {
	app.ResponseJSON(w, http.StatusOK, PingResponse{
		Code:        http.StatusOK,
		Description: "",
		Version:     "v1.0",
	})
}

func GetHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	return
	/*
		app.ResponseJSON(w, http.StatusOK, HealthResponse{
			Database: falses,
		})
	*/
}

func Custom404(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Vary", "Accept-Encoding")

	w.WriteHeader(http.StatusNotFound)
}
