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

type GetReceiptDocsResponse struct {
	Header struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
		Count  int `json:"count"`
	} `json:"header"`
	Data []whs.DocItem `json:"data"`
}

type ReceiptDoc struct {
	Id     int              `json:"id"`
	Number string           `json:"number"`
	Date   string           `json:"date"`
	Items  []ReceiptDocItem `json:"items"`
}

type ReceiptDocItem struct {
	ProductId             int64  `json:"product_id"`
	ProductName           string `json:"product_name"`
	ProductItemNumber     string `json:"product_item_number"`
	ProductManufacturer   string `json:"product_manufacturer"`
	ProductManufacturerId int64  `json:"product_manufacturer_id"`
	Quantity              int    `json:"quantity"`
	CellId                int64  `json:"cell_id"`
	CellName              string `json:"cell_name"`
}

func (r *ReceiptDoc) ImportData(item *whs.DocItem) {
	r.Id = int(item.Id)
	r.Number = item.Number
	r.Date = item.Date

	for _, v := range item.Items {
		row := ReceiptDocItem{
			ProductId:         v.Product.Id,
			ProductName:       v.Product.Name,
			ProductItemNumber: v.Product.ItemNumber,
			CellId:            v.CellDst.Id,
			CellName:          v.CellDst.Name,
			Quantity:          v.Quantity,
		}
		r.Items = append(r.Items, row)
	}
}

func (r *ReceiptDoc) ExportData() *whs.DocItem {
	docItem := whs.DocItem{
		Id:      int64(r.Id),
		Number:  r.Number,
		Date:    r.Date,
		DocType: 1,
	}
	for _, v := range r.Items {
		docItem.Items = append(docItem.Items, whs.DocRow{
			RowId: "",
			Product: whs.Product{
				RefItem: whs.RefItem{
					Id:   v.ProductId,
					Name: v.ProductName,
				},
				ItemNumber: v.ProductItemNumber,
				Barcodes:   nil,
				Manufacturer: whs.Manufacturer{
					RefItem: whs.RefItem(struct {
						Id       int64
						ParentId int64
						Name     string
					}{Id: v.ProductManufacturerId, ParentId: 0, Name: v.ProductManufacturer}),
				},
				Size: whs.SpecificSize{},
			},
			CellDst: whs.Cell{
				Id:   v.CellId,
				Name: v.CellName,
			},
			Quantity: v.Quantity,
		})
	}
	return &docItem
}

func RegisterReceiptHandlers(routeItems app.Routes, wHandlers *WrapHttpHandlers) app.Routes {

	routeItems = append(routeItems, app.Route{
		Name:          "GetReceiptDocs",
		Method:        "GET",
		Pattern:       "/receipt",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.GetReceiptDocs,
	})
	routeItems = append(routeItems, app.Route{
		Name:          "GetReceiptDoc",
		Method:        "GET",
		Pattern:       "/receipt/{id}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.GetReceiptDoc,
	})
	routeItems = append(routeItems, app.Route{
		Name:          "CreateReceiptDoc",
		Method:        "POST",
		Pattern:       "/receipt",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.CreateReceiptDoc,
	})
	routeItems = append(routeItems, app.Route{
		Name:          "UpdateReceiptDoc",
		Method:        "PUT",
		Pattern:       "/receipt/{id}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.UpdateReceiptDoc,
	})
	routeItems = append(routeItems, app.Route{
		Name:          "DeleteReceiptDoc",
		Method:        "DELETE",
		Pattern:       "/receipt/{id}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.DeleteReceiptDoc,
	})
	return routeItems
}
func (wh *WrapHttpHandlers) GetReceiptDocs(w http.ResponseWriter, r *http.Request) {

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

	docs, count, err := wh.Storage.GetReceiptDocsItems(offset, limit)
	if err != nil {
		app.Log.Warning.Printf("data fetch error, %v", err)
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
	}

	response := GetReceiptDocsResponse{}

	response.Data = make([]whs.DocItem, 0)

	response.Header.Limit = limit
	response.Header.Offset = offset
	response.Header.Count = count
	response.Data = docs

	app.ResponseJSON(w, http.StatusOK, response)
}

func (wh *WrapHttpHandlers) CreateReceiptDoc(w http.ResponseWriter, r *http.Request) {
	// читаем данные
	body, err := ioutil.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("can't read body"))
		return
	}

	data := new(ReceiptDoc)
	err = json.Unmarshal(body, data)
	if err != nil {
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("can't unmarshal body, %s", err))
		return
	}

	id, err := wh.Storage.CreateReceiptDoc(data.ExportData())

	if err != nil {
		if err, ok := err.(*core.WrapError); !ok {
			fmt.Println(err)
		} else {
			fmt.Println(err)
		}

		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
		return
	}

	app.ResponseJSON(w, http.StatusOK, id)
}

func (wh *WrapHttpHandlers) UpdateReceiptDoc(w http.ResponseWriter, r *http.Request) {
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

	_, err = wh.Storage.FindReceiptDocById(int64(valId))
	if err != nil {
		app.ResponseERROR(w, http.StatusNotFound, fmt.Errorf("document not found"))
		return
	}

	// читаем данные
	body, err := ioutil.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("can't read body"))
		return
	}

	data := new(ReceiptDoc)
	err = json.Unmarshal(body, data)
	if err != nil {
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("can't unmarshal body, %s", err))
		return
	}

	resultData, err := wh.Storage.UpdateReceiptDoc(data.ExportData())
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

func (wh *WrapHttpHandlers) GetReceiptDoc(w http.ResponseWriter, r *http.Request) {
	var err error
	var doc *whs.DocItem
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
	doc, err = wh.Storage.GetReceiptDocById(int64(valId))
	if err != nil {
		app.ResponseERROR(w, http.StatusNotFound, fmt.Errorf("document not found"))
		return
	}

	d := new(ReceiptDoc)
	d.ImportData(doc)
	app.ResponseJSON(w, http.StatusOK, d)
}

func (wh *WrapHttpHandlers) DeleteReceiptDoc(w http.ResponseWriter, r *http.Request) {
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

	m, err := wh.Storage.FindReceiptDocById(int64(valId))
	if err != nil {
		app.ResponseERROR(w, http.StatusNotFound, fmt.Errorf("document not found"))
		return
	}
	resultData, err := wh.Storage.DeleteReceiptDoc(m.Id)
	if err != nil {
		app.ResponseERROR(w, http.StatusInternalServerError, fmt.Errorf("item deleting error"))
		return
	}
	app.ResponseJSON(w, http.StatusOK, resultData)
}
