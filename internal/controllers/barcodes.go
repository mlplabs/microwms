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

type BarcodesController struct {
	service *service.WhsService
}

func NewBarcodesController(service *service.WhsService) *BarcodesController {
	return &BarcodesController{
		service: service,
	}
}

func (c *BarcodesController) RegisterHandlers(routeItems app.Routes) app.Routes {
	wrap := response.NewWrapper()
	routeItems = append(routeItems, app.Route{
		Name:          "GetBarcodeById",
		Method:        "GET",
		Pattern:       "/barcodes/{itemId}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.DataPlain(c.getById),
	})
	routeItems = append(routeItems, app.Route{
		Name:          "GetBarcodesAll",
		Method:        "GET",
		Pattern:       "/barcodes",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.DataPages(c.getBarcodes),
	})
	routeItems = append(routeItems, app.Route{
		Name:          "GetBarcodesByOwner",
		Method:        "GET",
		Pattern:       "/barcodes/{ownerRef}/{ownerId}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.DataPages(c.getBarcodesByOwner),
	})
	routeItems = append(routeItems, app.Route{
		Name:          "CreateBarcode",
		Method:        "POST",
		Pattern:       "/barcodes",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.DataPlain(c.createBarcode),
	})
	routeItems = append(routeItems, app.Route{
		Name:          "UpdateBarcode",
		Method:        "PUT",
		Pattern:       "/barcodes/{itemId}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.DataPlain(c.updateBarcode),
	})
	routeItems = append(routeItems, app.Route{
		Name:          "DeleteBarcode",
		Method:        "DELETE",
		Pattern:       "/barcodes/{itemId}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.Empty(c.deleteBarcode),
	})
	routeItems = append(routeItems, app.Route{
		Name:          "GetSuggestionBarcodes",
		Method:        "GET",
		Pattern:       "/suggestion/barcodes/{text}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.DataList(c.getSuggestion),
	})
	routeItems = append(routeItems, app.Route{
		Name:          "EnumBarcodeTypes",
		Method:        "GET",
		Pattern:       "/enum/barcodes/types",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.DataList(c.getBarcodeTypes),
	})
	return routeItems
}

// @Summary get barcodes list
// @Tags barcodes
// @Success		200	{object} response.Pagination{data=[]Barcode}
// @Param o query integer false "offset"
// @Param l query integer false "limit"
// @Router /barcodes [get]
func (c *BarcodesController) getBarcodes(r *http.Request) (interface{}, *response.DataRange, error) {
	ctx := r.Context()

	offset, limit := query.GetOffsetLimit(r)

	items, count, err := c.service.GetBarcodes(ctx, offset, limit)
	if err != nil {
		return items, &response.DataRange{}, err
	}

	return items, &response.DataRange{
		Count:  int(count),
		Limit:  limit,
		Offset: offset,
	}, nil
}

// @Summary get barcodes list by owner
// @Tags barcodes
// @Success		200	{object} response.Pagination{data=[]Barcode}
// @Param o query integer false "offset"
// @Param l query integer false "limit"
// @Param ownerRef path string true "owner table"
// @Param ownerID path integer true "owner ID"
// @Router /barcodes/{ownerRef}/{ownerID} [get]
func (c *BarcodesController) getBarcodesByOwner(r *http.Request) (interface{}, *response.DataRange, error) {
	ctx := r.Context()

	vars := mux.Vars(r)
	if vars["ownerRef"] == "" || vars["ownerId"] == "" {
		return nil, &response.DataRange{}, errors.NewInvalidInputData(fmt.Errorf("invalid query params"))
	}

	ownerId, ownerRef, err := query.GetOwner(r)
	if err != nil {
		return nil, &response.DataRange{}, err
	}

	offset, limit := query.GetOffsetLimit(r)

	items, count, err := c.service.GetBarcodesByOwner(ctx, offset, limit, int64(ownerId), ownerRef)
	if err != nil {
		return items, &response.DataRange{}, err
	}

	return items, &response.DataRange{
		Count:  int(count),
		Limit:  limit,
		Offset: offset,
	}, nil
}

// @Summary get barcode by ID
// @Tags barcodes
// @Success		200	{object} response.PlainData{data=Barcode}
// @Param itemID path integer true "barcode ID"
// @Router /barcodes/{itemID} [get]
func (c *BarcodesController) getById(r *http.Request) (interface{}, error) {
	var err error
	itemId, err := query.GetItemId(r)
	if err != nil {
		return nil, err
	}
	item, err := c.service.GetBarcodeById(r.Context(), int64(itemId))
	if err != nil {
		return item, errors.NewServerError(err)
	}

	return item, nil
}

// @Summary create barcode
// @Tags barcodes
// @Param request body Barcode true "barcode data"
// @Success		200	{object} response.PlainData{data=integer}
// @Router /barcodes [post]
func (c *BarcodesController) createBarcode(r *http.Request) (interface{}, error) {
	ctx := r.Context()
	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		return nil, errors.NewInvalidInputData(fmt.Errorf("can't read body"))
	}

	data := new(model.Barcode)
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, errors.NewInvalidInputData(fmt.Errorf("can't unmarshal body, %s", err))
	}

	if data.Name == "" || data.OwnerId == 0 || data.OwnerRef == "" {
		return nil, errors.NewInvalidInputData(fmt.Errorf("incorrect values"))
	}

	bcId, err := c.service.CreateBarcode(ctx, data)
	if err != nil {
		return nil, errors.NewBadRequest(fmt.Errorf("data fetch error, %s", err))
	}
	return bcId, nil
}

// @Summary update barcode
// @Tags barcodes
// @Param itemID path integer true "barcode ID"
// @Param request body Barcode true "barcode data"
// @Success		200	{object} response.PlainData{data=integer}
// @Router /barcodes/{itemID} [put]
func (c *BarcodesController) updateBarcode(r *http.Request) (interface{}, error) {
	ctx := r.Context()
	itemReqId, err := query.GetItemId(r)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		return nil, errors.NewInvalidInputData(fmt.Errorf("can't read body"))
	}

	data := new(model.Barcode)
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, errors.NewInvalidInputData(fmt.Errorf("can't unmarshal body, %s", err))
	}
	data.Id = int64(itemReqId)
	if data.Name == "" || data.OwnerId == 0 || data.OwnerRef == "" {
		return nil, errors.NewInvalidInputData(fmt.Errorf("incorrect values"))
	}

	bcId, err := c.service.UpdateBarcode(ctx, data)
	if err != nil {
		return nil, errors.NewBadRequest(fmt.Errorf("data fetch error, %s", err))
	}
	return bcId, nil
}

// @Summary delete barcode by ID
// @Tags barcodes
// @Success		200	{object} map[string]interface{}
// @Param itemID path integer true "barcode ID"
// @Router /barcodes/{itemID} [delete]
func (c *BarcodesController) deleteBarcode(r *http.Request) error {
	itemId, err := query.GetItemId(r)
	if err != nil {
		return err
	}

	err = c.service.DeleteBarcode(r.Context(), int64(itemId))
	if err != nil {
		return errors.NewServerError(fmt.Errorf("data fetch error, %s", err))
	}
	return nil
}

// @Summary get barcodes suggestion
// @Tags barcodes
// @Success		200	{object} response.List{data=[]Barcode}
// @Param text path string true "text for suggestion"
// @Router /suggestion/barcodes/{text} [get]
func (c *BarcodesController) getSuggestion(r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	if _, ok := vars["text"]; !ok {
		return nil, errors.NewInvalidInputData(fmt.Errorf("invalid path params"))
	}

	data, err := c.service.GetBarcodeSuggestion(r.Context(), vars["text"], 10)
	if err != nil {
		return nil, errors.NewServerError(err)
	}
	return data, nil
}

// @Summary enum barcode types
// @Tags barcodes
// @Success		200	{object} response.List{data=[]Type}
// @Router /enum/barcodes/types [get]
func (c *BarcodesController) getBarcodeTypes(r *http.Request) (interface{}, error) {
	data, err := c.service.GetBarcodeTypes(r.Context())
	if err != nil {
		return nil, errors.NewServerError(err)
	}
	return data, nil
}
