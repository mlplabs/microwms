package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	app "github.com/mlplabs/app-utils"
	"github.com/mlplabs/microwms-core"
	"github.com/mlplabs/microwms-core/whs"
	"github.com/mlplabs/microwms/conf"
	"github.com/mlplabs/microwms/routes"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var Storage *whs.Storage

func main() {
	app.Log.Init("", "")

	conf.ReadEnv()

	Storage = microwms_core.GetStorage()

	wHandlers := new(routes.WrapHttpHandlers)
	wHandlers.Storage = Storage

	err := wHandlers.Storage.Init(conf.Cfg.DbHost, conf.Cfg.DbName, conf.Cfg.DbUser, conf.Cfg.DbPassword)

	if err != nil {
		app.Log.Error.Fatalf("storage initialization failed, %v", err)
	}

	routeItems := app.Routes{}
	routeItems = RegisterHandlers(routeItems, wHandlers)
	router := NewRouter(routeItems)
	router.NotFoundHandler = http.HandlerFunc(routes.Custom404)
	router.Use(mux.CORSMethodMiddleware(router))
	StartHttpServer(router)
}

func StartHttpServer(router *mux.Router) {

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", conf.Cfg.AppAddr, conf.Cfg.AppPort),
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			app.Log.Error.Fatal(err)
		}
	}()

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	<-sigterm

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		app.Log.Error.Fatalf("server shutdown failed:%+v", err)
	}
	if err := Storage.Db.Close(); err != nil {
		app.Log.Error.Fatalf("database close failed:%+v", err)
	}

	app.Log.Info.Printf("Stopped Api server on %d port", 777)
}
func RegisterHandlers(routeItems app.Routes, wHandlers *routes.WrapHttpHandlers) app.Routes {
	routeItems = routes.RegisterControlHandlers(routeItems)
	routeItems = routes.RegisterProductsHandlers(routeItems, wHandlers)
	routeItems = routes.RegisterWhsHandlers(routeItems, wHandlers)
	routeItems = routes.RegisterManufacturersHandlers(routeItems, wHandlers)
	routeItems = routes.RegisterHardwareHandlers(routeItems, wHandlers)
	routeItems = routes.RegisterUsersHandlers(routeItems, wHandlers)
	routeItems = routes.RegisterReceiptHandlers(routeItems, wHandlers)
	routeItems = routes.RegisterBarcodesHandlers(routeItems, wHandlers)
	routeItems = routes.RegisterReportsHandlers(routeItems, wHandlers)

	return routeItems
}

func NewRouter(routeItems app.Routes) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routeItems {
		handlerFunc := route.HandlerFunc
		//if route.ValidateToken {
		//	handlerFunc = app.SetMiddlewareAuth(handlerFunc)
		//}

		if route.SetHeaderJSON {
			handlerFunc = app.SetMiddlewareJSON(handlerFunc)
		}

		router.
			Methods(route.Method, "OPTIONS").
			Path(route.Pattern).
			Name(route.Name).
			HandlerFunc(handlerFunc)
	}

	return router
}
