package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	l "github.com/mikelpsv/mod_logging"
	app "github.com/mikelpsv/mod_micro_app"
	"github.com/mlplabs/microwms-core"
	"github.com/mlplabs/microwms-core/models"
	"github.com/mlplabs/microwms/routes"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var Storage *models.Storage

func main() {
	l.Init("", "")

	Storage = microwms_core.GetStorage()

	wHandlers := new(routes.WrapHttpHandlers)
	wHandlers.Storage = Storage
	err := wHandlers.Storage.Init("localhost", "wmsdb", "devuser", "devuser")
	if err != nil {
		l.Error.Fatalf("storage initialization failed, %v", err)
	}

	routeItems := app.Routes{}
	routeItems = RegisterHandlers(routeItems, wHandlers)
	router := NewRouter(routeItems)
	router.NotFoundHandler = http.HandlerFunc(routes.Custom404)
	StartHttpServer(router)

	/*
		ps := s.GetProductService()
		ps.GetProductBarcodes(1)

		p := new(models.Product)
		p.Name = "Тестовый продукт 2"
		p.Size.SetSize(1, 2,3, 0.8)
		//p.Store(s)

		p, err := ps.FindProductsByBarcode("1234567890")
		if err != nil{
			fmt.Println(err)
		}
		fmt.Println(p)
		products, err := ps.GetProducts()
		if err != nil{
			fmt.Println(err)
		}

		fmt.Println(products)
	*/

	/*
		mfs, err := ps.GetManufacturers()
		if err != nil{
			fmt.Println(err)
		}
		fmt.Println(mfs)
	*/
}

func StartHttpServer(router *mux.Router) {

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", "127.0.0.1", 7777),
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			l.Error.Fatal(err)
		}
	}()

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	<-sigterm

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		l.Error.Fatalf("server shutdown failed:%+v", err)
	}
	if err := Storage.Db.Close(); err != nil {
		l.Error.Fatalf("database close failed:%+v", err)
	}

	l.Info.Printf("Stopped Api server on %d port", 777)
}
func RegisterHandlers(routeItems app.Routes, wHandlers *routes.WrapHttpHandlers) app.Routes {
	routeItems = routes.RegisterControlHandlers(routeItems)
	routeItems = routes.RegisterProductsHandlers(routeItems, wHandlers)
	return routeItems
}

func NewRouter(routeItems app.Routes) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routeItems {
		handlerFunc := route.HandlerFunc
		if route.ValidateToken {
			handlerFunc = app.SetMiddlewareAuth(handlerFunc)
		}

		if route.SetHeaderJSON {
			handlerFunc = app.SetMiddlewareJSON(handlerFunc)
		}

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			HandlerFunc(handlerFunc)
	}

	return router
}
