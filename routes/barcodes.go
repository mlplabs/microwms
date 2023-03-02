package routes

import (
	app "github.com/mlplabs/app-utils"
	"github.com/mlplabs/microwms-core/whs"
	"net/http"
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
