package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	app "github.com/mlplabs/app-utils"
	"github.com/mlplabs/app-utils/pkg/http/errors"
	"github.com/mlplabs/app-utils/pkg/http/response"
	"github.com/mlplabs/microwms/internal/controllers/query"
	"github.com/mlplabs/microwms/internal/usecase/service"
	"github.com/mlplabs/mwms-core/whs/model"
	"io"
	"net/http"
)

type ProductsController struct {
	service *service.WhsService
}

func NewProductsController(service *service.WhsService) *ProductsController {
	return &ProductsController{
		service: service,
	}
}
func (c *ProductsController) RegisterHandlers(routeItems app.Routes) app.Routes {
	wrap := response.NewWrapper()
	routeItems = append(routeItems, app.Route{
		Name:          "GetProduct",
		Method:        "GET",
		Pattern:       "/products/{itemId}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.DataPlain(c.getById),
	})
	routeItems = append(routeItems, app.Route{
		Name:          "GetProducts",
		Method:        "GET",
		Pattern:       "/products",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.DataPages(c.getProducts),
	})
	routeItems = append(routeItems, app.Route{
		Name:          "FindByName",
		Method:        "GET",
		Pattern:       "/products/find/name/{name}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.DataList(c.findProductsByName),
	})
	routeItems = append(routeItems, app.Route{
		Name:          "FindByBarcode",
		Method:        "GET",
		Pattern:       "/products/find/barcode/{name}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.DataList(c.findProductsByBarcode),
	})
	routeItems = append(routeItems, app.Route{
		Name:          "CreateProduct",
		Method:        "POST",
		Pattern:       "/products",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.DataPlain(c.createProduct),
	})
	routeItems = append(routeItems, app.Route{
		Name:          "UpdateProduct",
		Method:        "PUT",
		Pattern:       "/products/{itemId}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.DataPlain(c.updateProduct),
	})
	routeItems = append(routeItems, app.Route{
		Name:          "DeleteProduct",
		Method:        "DELETE",
		Pattern:       "/products/{itemId}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.Empty(c.deleteProduct),
	})
	routeItems = append(routeItems, app.Route{
		Name:          "GetSuggestionProducts",
		Method:        "GET",
		Pattern:       "/suggestion/products/{text}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.DataList(c.getSuggestion),
	})
	return routeItems
}

// @Summary get product by ID
// @Tags products
// @Success		200	{object} response.PlainData{data=Product}
// @Param itemID path integer true "user ID"
// @Router /products/{itemID} [get]
func (c *ProductsController) getById(r *http.Request) (interface{}, error) {
	var err error
	itemId, err := query.GetItemId(r)
	if err != nil {
		return nil, err
	}
	item, err := c.service.GetProductById(r.Context(), int64(itemId))
	if err != nil {
		return item, errors.NewServerError(err)
	}

	return item, nil
}

// @Summary get products list
// @Tags products
// @Success		200	{object} response.Pagination{data=[]Product}
// @Param o query integer false "offset"
// @Param l query integer false "limit"
// @Router /products [get]
func (c *ProductsController) getProducts(r *http.Request) (interface{}, *response.DataRange, error) {
	ctx := r.Context()
	offset, limit := query.GetOffsetLimit(r)
	search := r.URL.Query().Get("s")
	items, count, err := c.service.GetProducts(ctx, offset, limit, search)
	if err != nil {
		return items, &response.DataRange{}, err
	}

	return items, &response.DataRange{
		Count:  int(count),
		Limit:  limit,
		Offset: offset,
	}, nil
}

// @Summary create product
// @Tags products
// @Param request body Product true "product data"
// @Success		200	{object} response.PlainData{data=integer}
// @Router /products [post]
func (c *ProductsController) createProduct(r *http.Request) (interface{}, error) {
	ctx := r.Context()
	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		return nil, errors.NewInvalidInputData(fmt.Errorf("can't read body"))
	}

	data := new(model.Product)
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, errors.NewInvalidInputData(fmt.Errorf("can't unmarshal body, %s", err))
	}

	userId, err := c.service.CreateProduct(ctx, data)
	if err != nil {
		return nil, errors.NewBadRequest(fmt.Errorf("data fetch error, %s", err))
	}
	return userId, nil
}

// @Summary update product
// @Tags products
// @Param itemID path integer true "product ID"
// @Param request body Product true "product data"
// @Success		200	{object} response.PlainData{data=integer}
// @Router /products/{itemID} [put]
func (c *ProductsController) updateProduct(r *http.Request) (interface{}, error) {
	ctx := r.Context()
	itemReqId, err := query.GetItemId(r)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		return nil, errors.NewInvalidInputData(fmt.Errorf("can't read body"))
	}

	data := new(model.Product)
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, errors.NewInvalidInputData(fmt.Errorf("can't unmarshal body, %s", err))
	}
	data.Id = int64(itemReqId)
	userId, err := c.service.UpdateProduct(ctx, data)
	if err != nil {
		return nil, errors.NewBadRequest(fmt.Errorf("data fetch error, %s", err))
	}
	return userId, nil
}

// @Summary delete product by ID
// @Tags products
// @Success		200	{object} map[string]interface{}
// @Param itemID path integer true "product ID"
// @Router /products/{itemID} [delete]
func (c *ProductsController) deleteProduct(r *http.Request) error {
	itemId, err := query.GetItemId(r)
	if err != nil {
		return err
	}

	err = c.service.DeleteProduct(r.Context(), int64(itemId))
	if err != nil {
		return errors.NewServerError(fmt.Errorf("data fetch error, %s", err))
	}
	return nil
}

// @Summary find products by name
// @Tags products
// @Success		200	{object} response.Pagination{data=[]Product}
// @Param name path string true "product name"
// @Router /products/find/name/{name} [get]
func (c *ProductsController) findProductsByName(r *http.Request) (interface{}, error) {
	ctx := r.Context()
	vars := mux.Vars(r)
	if _, ok := vars["name"]; !ok {
		return nil, errors.NewInvalidInputData(fmt.Errorf("invalid path params"))
	}
	items, err := c.service.FindProductsByName(ctx, vars["name"])
	if err != nil {
		return items, err
	}

	return items, nil
}

// @Summary find products by barcode
// @Tags products
// @Success		200	{object} response.Pagination{data=[]Product}
// @Param name path string true "barcode name"
// @Router /products/find/barcode/{name} [get]
func (c *ProductsController) findProductsByBarcode(r *http.Request) (interface{}, error) {
	ctx := r.Context()
	vars := mux.Vars(r)
	if _, ok := vars["name"]; !ok {
		return nil, errors.NewInvalidInputData(fmt.Errorf("invalid path params"))
	}
	items, err := c.service.FindProductsByBarcode(ctx, vars["name"])
	if err != nil {
		return items, err
	}

	return items, nil
}

// @Summary get products suggestion
// @Tags products
// @Success		200	{object} response.List{data=[]Product}
// @Param text path string true "text for suggestion"
// @Router /suggestion/products/{text} [get]
func (c *ProductsController) getSuggestion(r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	if _, ok := vars["text"]; !ok {
		return nil, errors.NewInvalidInputData(fmt.Errorf("invalid path params"))
	}

	data, err := c.service.GetProductsSuggestion(r.Context(), vars["text"], 10)
	if err != nil {
		return nil, errors.NewServerError(err)
	}
	return data, nil
}

//
//import (
//	"encoding/json"
//	"fmt"
//	"github.com/gorilla/mux"
//	app "github.com/mlplabs/app-utils"
//	"github.com/mlplabs/mwms-core/core"
//	"github.com/mlplabs/mwms-core/whs"
//	"io/ioutil"
//	"net/http"
//	"strconv"
//)
//
//type GetProductsResponse struct {
//	Header struct {
//		Limit  int `json:"limit"`
//		Offset int `json:"offset"`
//		Count  int `json:"count"`
//	} `json:"header"`
//	Data []whs.Product `json:"data"`
//}
//
//func RegisterProductsHandlers(routeItems app.Routes, wHandlers *WrapHttpHandlers) app.Routes {
//	routeItems = append(routeItems, app.Route{
//		Name:          "GetProducts",
//		Method:        "GET",
//		Pattern:       "/products",
//		SetHeaderJSON: true,
//		ValidateToken: false,
//		HandlerFunc:   wHandlers.GetProducts,
//	})
//	routeItems = append(routeItems, app.Route{
//		Name:          "GetProduct",
//		Method:        "GET",
//		Pattern:       "/products/{id}",
//		SetHeaderJSON: true,
//		ValidateToken: false,
//		HandlerFunc:   wHandlers.GetProductById,
//	})
//	routeItems = append(routeItems, app.Route{
//		Name:          "GetProductByBarcode",
//		Method:        "GET",
//		Pattern:       "/products/barcode/{barcode}",
//		SetHeaderJSON: true,
//		ValidateToken: false,
//		HandlerFunc:   wHandlers.GetProductsByBarcode,
//	})
//	routeItems = append(routeItems, app.Route{
//		Name:          "CreateProduct",
//		Method:        "POST",
//		Pattern:       "/products",
//		SetHeaderJSON: true,
//		ValidateToken: false,
//		HandlerFunc:   wHandlers.CreateProduct,
//	})
//	routeItems = append(routeItems, app.Route{
//		Name:          "UpdateProduct",
//		Method:        "PUT",
//		Pattern:       "/products",
//		SetHeaderJSON: true,
//		ValidateToken: false,
//		HandlerFunc:   wHandlers.UpdateProduct,
//	})
//	routeItems = append(routeItems, app.Route{
//		Name:          "DeleteProduct",
//		Method:        "DELETE",
//		Pattern:       "/products/{id}",
//		SetHeaderJSON: true,
//		ValidateToken: false,
//		HandlerFunc:   wHandlers.DeleteProduct,
//	})
//
//	routeItems = append(routeItems, app.Route{
//		Name:          "GetZonesOfWarehouse",
//		Method:        "GET",
//		Pattern:       "/warehouses/{whs_id}/zones/",
//		SetHeaderJSON: true,
//		ValidateToken: false,
//		HandlerFunc:   wHandlers.GetZonesOfWarehouse,
//	})
//	routeItems = append(routeItems, app.Route{
//		Name:          "GetSuggestionProducts",
//		Method:        "GET",
//		Pattern:       "/suggestion/products/{text}",
//		SetHeaderJSON: true,
//		ValidateToken: false,
//		HandlerFunc:   wHandlers.GetSuggestionProducts,
//	})
//
//	return routeItems
//}
//
//func (wh *WrapHttpHandlers) GetProducts(w http.ResponseWriter, r *http.Request) {
//	varOffset := r.URL.Query().Get("o")
//	varLimit := r.URL.Query().Get("l")
//
//	offset, err := strconv.Atoi(varOffset)
//	if err != nil {
//		offset = 0
//	}
//	limit, err := strconv.Atoi(varLimit)
//	if err != nil {
//		limit = 0
//	}
//
//	products, count, err := wh.Storage.GetProductsItems(offset, limit, 0)
//	if err != nil {
//		app.Log.Warning.Printf("data fetch error, %v", err)
//		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
//	}
//
//	response := GetProductsResponse{}
//	response.Data = make([]whs.Product, 0)
//
//	response.Header.Limit = limit
//	response.Header.Offset = offset
//	response.Header.Count = count
//	response.Data = products
//
//	app.ResponseJSON(w, http.StatusOK, response)
//}
//
//func (wh *WrapHttpHandlers) GetProductsByBarcode(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	val := vars["barcode"]
//
//	prod, err := wh.Storage.FindProductsByBarcode(val)
//	if err != nil {
//		app.ResponseERROR(w, http.StatusNotFound, fmt.Errorf("product not found"))
//		return
//	}
//	app.ResponseJSON(w, http.StatusOK, prod)
//}
//
//func (wh *WrapHttpHandlers) GetProductById(w http.ResponseWriter, r *http.Request) {
//	var err error
//	var prod *whs.Product
//	vars := mux.Vars(r)
//
//	if v, ok := vars["id"]; !ok || v == "0" {
//		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("invalid path params"))
//		return
//	}
//
//	valId, err := strconv.Atoi(vars["id"])
//	if err != nil {
//		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("invalid query params"))
//		return
//	}
//	prod, err = wh.Storage.FindProductById(int64(valId))
//	if err != nil {
//		app.ResponseERROR(w, http.StatusNotFound, fmt.Errorf("product not found"))
//		return
//	}
//	app.ResponseJSON(w, http.StatusOK, prod)
//}
//
//func (wh *WrapHttpHandlers) GetZonesOfWarehouse(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//
//	if _, ok := vars["whs_id"]; !ok {
//		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("invalid path params"))
//		return
//	}
//
//	valId, err := strconv.Atoi(vars["whs_id"])
//	if err != nil {
//		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("invalid query params"))
//		return
//	}
//
//	whs, err := wh.Storage.FindWhsById(int64(valId))
//	if err != nil {
//		app.Log.Warning.Printf("data fetch error, %v", err)
//		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
//		return
//	}
//
//	whss, err := wh.Storage.GetWhsZones(whs)
//	if err != nil {
//		app.Log.Warning.Printf("data fetch error, %v", err)
//		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
//		return
//	}
//	app.ResponseJSON(w, http.StatusOK, whss)
//}
//
//func (wh *WrapHttpHandlers) CreateProduct(w http.ResponseWriter, r *http.Request) {
//	// читаем данные
//	body, err := ioutil.ReadAll(r.Body)
//	if err != nil || len(body) == 0 {
//		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("can't read body"))
//		return
//	}
//
//	data := new(whs.Product)
//	err = json.Unmarshal(body, data)
//	if err != nil {
//		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("can't unmarshal body, %s", err))
//		return
//	}
//
//	whss, err := wh.Storage.CreateProduct(data)
//	if err != nil {
//		if err, ok := err.(*core.WrapError); !ok {
//			fmt.Println(err)
//		} else {
//			fmt.Println(err)
//		}
//
//		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
//		return
//	}
//
//	app.ResponseJSON(w, http.StatusOK, whss)
//}
//
//func (wh *WrapHttpHandlers) UpdateProduct(w http.ResponseWriter, r *http.Request) {
//	// читаем данные
//	body, err := ioutil.ReadAll(r.Body)
//	if err != nil || len(body) == 0 {
//		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("can't read body"))
//		return
//	}
//
//	data := new(whs.Product)
//	err = json.Unmarshal(body, data)
//	if err != nil {
//		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("can't unmarshal body, %s", err))
//		return
//	}
//
//	_, err = wh.Storage.FindProductById(data.Id)
//	if err != nil {
//		app.ResponseERROR(w, http.StatusNotFound, fmt.Errorf("product not found"))
//		return
//	}
//
//	resultData, err := wh.Storage.UpdateProduct(data)
//	if err != nil {
//		if e, ok := err.(*core.WrapError); !ok {
//			fmt.Println(e)
//		} else {
//			fmt.Println(e)
//		}
//
//		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
//		return
//	}
//	app.ResponseJSON(w, http.StatusOK, resultData)
//}
//
//func (wh *WrapHttpHandlers) DeleteProduct(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//
//	if v, ok := vars["id"]; !ok || v == "0" {
//		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("invalid path params"))
//		return
//	}
//
//	valId, err := strconv.Atoi(vars["id"])
//	if err != nil {
//		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("invalid query params"))
//		return
//	}
//
//	p, err := wh.Storage.FindProductById(int64(valId))
//	if err != nil {
//		app.ResponseERROR(w, http.StatusNotFound, fmt.Errorf("product not found"))
//		return
//	}
//	resultData, err := wh.Storage.DeleteProduct(p)
//	if err != nil {
//		if err.(*core.WrapError).Code == core.ForeignKeyError {
//			app.ResponseERROR(w, http.StatusInternalServerError, fmt.Errorf("item deleting error, foreign key"))
//			return
//		}
//		app.ResponseERROR(w, http.StatusInternalServerError, fmt.Errorf("item deleting error"))
//		return
//	}
//	app.ResponseJSON(w, http.StatusOK, resultData)
//}
//
//func (wh *WrapHttpHandlers) GetSuggestionProducts(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	if _, ok := vars["text"]; !ok {
//		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("invalid path params"))
//		return
//	}
//
//	data, err := wh.Storage.GetProductsSuggestion(vars["text"], 10)
//	if err != nil {
//		app.Log.Warning.Printf("data fetch error, %v", err)
//		app.ResponseERROR(w, http.StatusBadRequest, fmt.Errorf("data fetch error"))
//		return
//	}
//	app.ResponseJSON(w, http.StatusOK, data)
//}
