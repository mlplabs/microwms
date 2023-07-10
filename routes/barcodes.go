package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	app "github.com/mlplabs/app-utils"
	"github.com/mlplabs/microwms-core/whs"
	"net/http"
	"strconv"
)

func RegisterBarcodesHandlers(routeItems app.Routes, wHandlers *WrapHttpHandlers) app.Routes {

	routeItems = append(routeItems, app.Route{
		Name:          "GetBarcodeEnumType",
		Method:        "GET",
		Pattern:       "/enum/barcode_type",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.GetBarcodeEnumType,
	})

	routeItems = append(routeItems, app.Route{
		Name:          "DeleteBarcode",
		Method:        "DELETE",
		Pattern:       "/barcodes/{id}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.DeleteBarcode,
	})

	return routeItems
}

func (wh *WrapHttpHandlers) GetBarcodeEnumType(w http.ResponseWriter, r *http.Request) {
	type bcItemType struct {
		Key int    `json:"key"`
		Val string `json:"val"`
	}

	bc := make([]bcItemType, 0)
	bc = append(bc, bcItemType{Key: whs.BarcodeTypeUnknown, Val: "-"})
	bc = append(bc, bcItemType{Key: whs.BarcodeTypeEAN13, Val: "EAN13"})
	bc = append(bc, bcItemType{Key: whs.BarcodeTypeEAN8, Val: "EAN8"})
	bc = append(bc, bcItemType{Key: whs.BarcodeTypeEAN14, Val: "EAN14"})
	bc = append(bc, bcItemType{Key: whs.BarcodeTypeCode128, Val: "CODE128"})

	app.ResponseJSON(w, http.StatusOK, bc)
}

func (wh *WrapHttpHandlers) DeleteBarcode(w http.ResponseWriter, r *http.Request) {
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

	delId, err := wh.Storage.DeleteBarcodeById(int64(valId))
	if err != nil {
		app.ResponseERROR(w, http.StatusInternalServerError, fmt.Errorf("item deleting error"))
		return
	}
	app.ResponseJSON(w, http.StatusOK, delId)
}
