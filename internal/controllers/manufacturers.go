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

type ManufacturersController struct {
	service *service.WhsService
}

func NewManufacturersController(service *service.WhsService) *ManufacturersController {
	return &ManufacturersController{
		service: service,
	}
}

func (c *ManufacturersController) RegisterHandlers(routeItems app.Routes) app.Routes {
	wrap := response.NewWrapper()
	routeItems = append(routeItems, app.Route{
		Name:          "GetManufacturer",
		Method:        "GET",
		Pattern:       "/manufacturers/{itemId}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.DataPlain(c.getManufacturerById),
	})
	routeItems = append(routeItems, app.Route{
		Name:          "GetManufacturers",
		Method:        "GET",
		Pattern:       "/manufacturers",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.DataPages(c.getManufacturers),
	})
	routeItems = append(routeItems, app.Route{
		Name:          "CreateManufacturer",
		Method:        "POST",
		Pattern:       "/manufacturers",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.DataPlain(c.createManufacturer),
	})
	routeItems = append(routeItems, app.Route{
		Name:          "UpdateManufacturer",
		Method:        "PUT",
		Pattern:       "/manufacturers/{itemId}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.DataPlain(c.updateManufacturer),
	})
	routeItems = append(routeItems, app.Route{
		Name:          "DeleteManufacturer",
		Method:        "DELETE",
		Pattern:       "/manufacturers/{itemId}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.Empty(c.deleteManufacturer),
	})
	routeItems = append(routeItems, app.Route{
		Name:          "GetSuggestionProducts",
		Method:        "GET",
		Pattern:       "/suggestion/manufacturers/{text}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.DataList(c.getSuggestion),
	})
	return routeItems
}

// @Summary get manufacturer by ID
// @Tags manufacturers
// @Success		200	{object} response.PlainData{data=Manufacturer}
// @Param itemID path integer true "manufacturer ID"
// @Router /manufacturers/{itemID} [get]
func (c *ManufacturersController) getManufacturerById(r *http.Request) (interface{}, error) {
	itemId, err := query.GetItemId(r)
	if err != nil {
		return nil, err
	}
	items, err := c.service.GetManufacturerById(r.Context(), int64(itemId))
	if err != nil {
		return items, errors.NewServerError(err)
	}
	return items, nil
}

// @Summary get manufacturers list
// @Tags manufacturers
// @Success		200	{object} response.Pagination{data=[]Manufacturer}
// @Param o query integer false "offset"
// @Param l query integer false "limit"
// @Router /manufacturers [get]
func (c *ManufacturersController) getManufacturers(r *http.Request) (interface{}, *response.DataRange, error) {
	offset, limit := query.GetOffsetLimit(r)
	search := r.URL.Query().Get("s")
	items, count, err := c.service.GetManufacturers(r.Context(), offset, limit, search)
	if err != nil {
		return nil, &response.DataRange{}, errors.NewServerError(err)
	}
	return items, &response.DataRange{
		Count:  int(count),
		Limit:  limit,
		Offset: offset,
	}, nil
}

// @Summary create manufacturer
// @Tags manufacturers
// @Param request body Manufacturer true "manufacturer data"
// @Success		200	{object} response.PlainData{data=integer}
// @Router /manufacturers [post]
func (c *ManufacturersController) createManufacturer(r *http.Request) (interface{}, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		return 0, errors.NewInvalidInputData(fmt.Errorf("can't read body"))
	}

	data := model.Manufacturer{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return 0, errors.NewInvalidInputData(fmt.Errorf("can't unmarshal body, %s", err))
	}

	itemId, err := c.service.CreateManufacturer(r.Context(), &data)
	if err != nil {
		return nil, errors.NewServerError(err)
	}
	return itemId, nil
}

// @Summary update manufacturer
// @Tags manufacturers
// @Param itemID path integer true "manufacturer ID"
// @Param request body Manufacturer true "manufacturer data"
// @Success		200	{object} response.PlainData{data=integer}
// @Router /manufacturers/{itemID} [put]
func (c *ManufacturersController) updateManufacturer(r *http.Request) (interface{}, error) {
	itemReqId, err := query.GetItemId(r)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		return 0, errors.NewInvalidInputData(fmt.Errorf("can't read body"))
	}

	data := model.Manufacturer{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return 0, errors.NewInvalidInputData(fmt.Errorf("can't unmarshal body, %s", err))
	}

	data.Id = int64(itemReqId)
	itemId, err := c.service.UpdateManufacturer(r.Context(), &data)
	if err != nil {
		return nil, errors.NewServerError(err)
	}
	return itemId, nil
}

// @Summary delete manufacturer by ID
// @Tags manufacturers
// @Success		200	{object} map[string]interface{}
// @Param itemID path integer true "manufacturer ID"
// @Router /manufacturers/{itemID} [delete]
func (c *ManufacturersController) deleteManufacturer(r *http.Request) error {
	itemId, err := query.GetItemId(r)
	if err != nil {
		return err
	}

	err = c.service.DeleteManufacturer(r.Context(), int64(itemId))
	if err != nil {
		return errors.NewServerError(fmt.Errorf("data fetch error, %s", err))
	}
	return nil
}

// @Summary get manufacturers suggestion
// @Tags manufacturers
// @Success		200	{object} response.List{data=[]Manufacturer}
// @Param text path string true "text for suggestion"
// @Router /suggestion/manufacturers/{text} [get]
func (c *ManufacturersController) getSuggestion(r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	if _, ok := vars["text"]; !ok {
		return nil, errors.NewInvalidInputData(fmt.Errorf("invalid path params"))
	}

	data, err := c.service.GetManufacturerSuggestion(r.Context(), vars["text"], 10)
	if err != nil {
		return nil, errors.NewServerError(err)
	}
	return data, nil
}
