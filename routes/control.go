package routes

import (
	"fmt"
	l "github.com/mikelpsv/mod_logging"
	app "github.com/mikelpsv/mod_micro_app"
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
	pingErr := app.Db.Ping()
	if pingErr != nil {
		l.Error.Printf("get health returned db error, %v", pingErr)
	}
	app.ResponseJSON(w, http.StatusOK, HealthResponse{
		Database: pingErr == nil,
	})
}

func Custom404(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RequestURI)
	w.WriteHeader(http.StatusNotFound)
}
