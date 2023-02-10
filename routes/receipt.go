package routes

import (
	"encoding/json"
	"fmt"
	app "github.com/mlplabs/app-utils"
	"github.com/mlplabs/microwms-core/core"
	"github.com/mlplabs/microwms-core/models"
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
	Data []models.DocItem `json:"data"`
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
	//routeItems = append(routeItems, app.Route{
	//	Name:          "GetUser",
	//	Method:        "GET",
	//	Pattern:       "/users/{id}",
	//	SetHeaderJSON: true,
	//	ValidateToken: false,
	//	HandlerFunc:   wHandlers.GetUserById,
	//})
	routeItems = append(routeItems, app.Route{
		Name:          "CreateReceiptDoc",
		Method:        "POST",
		Pattern:       "/receipt",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.CreateReceiptDoc,
	})
	//routeItems = append(routeItems, app.Route{
	//	Name:          "UpdateUser",
	//	Method:        "PUT",
	//	Pattern:       "/users/{id}",
	//	SetHeaderJSON: true,
	//	ValidateToken: false,
	//	HandlerFunc:   wHandlers.UpdateUser,
	//})
	//routeItems = append(routeItems, app.Route{
	//	Name:          "DeleteUser",
	//	Method:        "DELETE",
	//	Pattern:       "/users/{id}",
	//	SetHeaderJSON: true,
	//	ValidateToken: false,
	//	HandlerFunc:   wHandlers.DeleteUser,
	//})
	//routeItems = append(routeItems, app.Route{
	//	Name:          "GetSuggestionUsers",
	//	Method:        "GET",
	//	Pattern:       "/suggestion/users/{text}",
	//	SetHeaderJSON: true,
	//	ValidateToken: false,
	//	HandlerFunc:   wHandlers.GetSuggestionUsers,
	//})

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

	doc := wh.Storage.GetDocReceipt()
	docs, count, err := doc.GetItems(offset, limit)
	if err != nil {
		app.Log.Warning.Printf("data fetch error, %v", err)
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
	}

	response := GetReceiptDocsResponse{}

	response.Data = make([]models.DocItem, 0)

	response.Header.Limit = limit
	response.Header.Offset = offset
	response.Header.Count = count
	response.Data = docs

	app.ResponseJSON(w, http.StatusOK, response)
}

func (wh *WrapHttpHandlers) CreateReceiptDoc(w http.ResponseWriter, r *http.Request) {
	type InDoc struct {
		Id     int    `json:"id"`
		Number string `json:"number"`
		Date   string `json:"date"`
		Items  []struct {
			ProductId   int    `json:"product_id"`
			ProductName string `json:"product_name"`
			Quantity    int
		}
	}

	// читаем данные
	body, err := ioutil.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("can't read body"))
		return
	}

	data := new(InDoc)
	err = json.Unmarshal(body, data)
	if err != nil {
		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("can't unmarshal body, %s", err))
		return
	}
	docItem := models.DocItem{
		Id:      int64(data.Id),
		Number:  data.Number,
		Date:    data.Date,
		DocType: 1,
	}

	for _, v := range data.Items {
		docItem.Items = append(docItem.Items, models.DocRow{
			RowNum: "",
			Product: models.Product{
				Id:           0,
				Name:         v.ProductName,
				ItemNumber:   "",
				Barcodes:     nil,
				Manufacturer: models.Manufacturer{},
				Size:         models.SpecificSize{},
			},
			Quantity: v.Quantity,
		})
	}

	docReceipt := wh.Storage.GetDocReceipt()
	id, err := docReceipt.Create(&docItem)

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
