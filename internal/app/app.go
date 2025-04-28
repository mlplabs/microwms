package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	utils "github.com/mlplabs/app-utils"
	"github.com/mlplabs/microwms/internal/config"
	"github.com/mlplabs/microwms/internal/controllers"
	"github.com/mlplabs/microwms/internal/usecase/service"
	"github.com/mlplabs/mwms-core/whs"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Init(cfg *config.AppCfg) {
	ctx, cancel := context.WithCancel(context.Background())
	utils.Log.Init("", "")
	utils.InitDb(cfg.DbHost, cfg.DbPort, cfg.DbName, cfg.DbUser, cfg.DbPassword)
	defer utils.Db.Close()

	thisService := service.NewWhsService(whs.NewWms(utils.Db))
	ctrlUsers := controllers.NewUsersController(thisService)
	ctrlBarcodes := controllers.NewBarcodesController(thisService)
	ctrlManufacturers := controllers.NewManufacturersController(thisService)
	ctrlProducts := controllers.NewProductsController(thisService)
	ctrlWarehouses := controllers.NewWarehouseController(thisService)
	ctrlCells := controllers.NewCellsController(thisService)
	ctrlReports := controllers.NewReportsController(thisService)
	ctrlStorage := controllers.NewStorageController(thisService)

	//ctrlBase := service.NewController(thisService)
	//ctrlUtils := service_utils.NewController(thisService)
	routeItems := utils.Routes{}
	routeItems = ctrlUsers.RegisterHandlers(routeItems)
	routeItems = ctrlBarcodes.RegisterHandlers(routeItems)
	routeItems = ctrlManufacturers.RegisterHandlers(routeItems)
	routeItems = ctrlProducts.RegisterHandlers(routeItems)
	routeItems = ctrlWarehouses.RegisterHandlers(routeItems)
	routeItems = ctrlCells.RegisterHandlers(routeItems)
	routeItems = ctrlReports.RegisterHandlers(routeItems)
	routeItems = ctrlStorage.RegisterHandlers(routeItems)

	router := NewRouter(routeItems)
	router.Use(mux.CORSMethodMiddleware(router))

	server := &http.Server{
		Addr:     fmt.Sprintf("%s:%s", cfg.AppAddr, cfg.AppPort),
		Handler:  router,
		ErrorLog: utils.Log.Error,
	}
	fmt.Printf("Service started. Listen %s:%s", cfg.AppAddr, cfg.AppPort)
	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			utils.Log.Error.Fatal(err)
		}
	}()
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	<-sigterm

	cancel()

	if err := server.Shutdown(ctx); err != nil {
		utils.Log.Error.Fatalf("server shutdown failed:%+v", err)
	}

	fmt.Println("Service stopped")
	time.Sleep(5 * time.Second)
}

func NewRouter(routeItems utils.Routes) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routeItems {
		handlerFunc := route.HandlerFunc
		if route.SetHeaderJSON {
			handlerFunc = utils.SetMiddlewareJSON(handlerFunc)
		}
		router.
			Methods(route.Method, "OPTIONS").
			Path(route.Pattern).
			Name(route.Name).
			HandlerFunc(handlerFunc)
	}
	return router
}
