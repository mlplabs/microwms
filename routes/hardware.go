package routes

import (
	app "github.com/mlplabs/app-utils"
	"net/http"
)

type Printer struct {
	Name     string `json:"name"`
	Instance string `json:"instance"`
	Status   string `json:"status"`
}

func RegisterHardwareHandlers(routeItems app.Routes, wHandlers *WrapHttpHandlers) app.Routes {
	routeItems = append(routeItems, app.Route{
		Name:          "GetPrinters",
		Method:        "GET",
		Pattern:       "/printers",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.GetPrinters,
	})
	return routeItems
}

func (wh *WrapHttpHandlers) GetPrinters(w http.ResponseWriter, r *http.Request) {
	//resPrinters := make([]Printer, 0)
	//printers := cups.NewConnection()
	//n, err := printers.EnumDestinations()
	//if err != nil {
	//	fmt.Printf("%v", err.Error())
	//	return
	//}
	//fmt.Printf("found %d\n", n)
	//
	//for _, dest := range printers.Dests {
	//	p := Printer{}
	//	p.Name = dest.Name
	//	p.Instance = dest.Instance
	//	p.Status = dest.Status()
	//	resPrinters = append(resPrinters, p)
	//}
	//app.ResponseJSON(w, http.StatusOK, resPrinters)
}
