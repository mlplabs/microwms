package controllers

import (
	"encoding/json"
	"fmt"
	app "github.com/mlplabs/app-utils"
	"github.com/mlplabs/app-utils/pkg/http/errors"
	"github.com/mlplabs/app-utils/pkg/http/response"
	"github.com/mlplabs/microwms/internal/usecase/service"
	"io"
	"net/http"
)

type paramStorage struct {
	ProductId int64 `json:"product_id"`
	CellId    int64 `json:"cell_id"`
	Quantity  int   `json:"quantity"`
}

type paramStorageMove struct {
	ProductId int64 `json:"product_id"`
	SrcCellId int64 `json:"src_cell_id"`
	DstCellId int64 `json:"dst_cell_id"`
	Quantity  int   `json:"quantity"`
}

type StorageController struct {
	service *service.WhsService
}

func NewStorageController(service *service.WhsService) *StorageController {
	return &StorageController{
		service: service,
	}
}

func (s *StorageController) RegisterHandlers(routeItems app.Routes) app.Routes {
	wrap := response.NewWrapper()
	routeItems = append(routeItems, app.Route{
		Name:          "GetProduct",
		Method:        "POST",
		Pattern:       "/storage/get",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.DataPlain(s.getProduct),
	})
	routeItems = append(routeItems, app.Route{
		Name:          "PutProduct",
		Method:        "POST",
		Pattern:       "/storage/put",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.DataPlain(s.putProduct),
	})
	routeItems = append(routeItems, app.Route{
		Name:          "MoveProduct",
		Method:        "POST",
		Pattern:       "/storage/move",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.DataPlain(s.moveProduct),
	})
	routeItems = append(routeItems, app.Route{
		Name:          "QuantityProduct",
		Method:        "GET",
		Pattern:       "/storage/quantity",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.DataPlain(s.quantityProduct),
	})
	return routeItems
}

func (s *StorageController) getProduct(r *http.Request) (interface{}, error) {
	ctx := r.Context()
	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		return nil, errors.NewInvalidInputData(fmt.Errorf("can't read body"))
	}

	data := new(paramStorage)
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, errors.NewInvalidInputData(fmt.Errorf("can't unmarshal body, %s", err))
	}

	res, err := s.service.GetItemFromCell(ctx, data.ProductId, data.CellId, data.Quantity)
	if err != nil {
		return nil, errors.NewBadRequest(fmt.Errorf("data fetch error, %s", err))
	}
	return res, nil
}

func (s *StorageController) putProduct(r *http.Request) (interface{}, error) {
	ctx := r.Context()
	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		return nil, errors.NewInvalidInputData(fmt.Errorf("can't read body"))
	}

	data := new(paramStorage)
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, errors.NewInvalidInputData(fmt.Errorf("can't unmarshal body, %s", err))
	}

	res, err := s.service.PutItemToCell(ctx, data.ProductId, data.CellId, data.Quantity)
	if err != nil {
		return nil, errors.NewBadRequest(fmt.Errorf("data fetch error, %s", err))
	}
	return res, nil
}

func (s *StorageController) moveProduct(r *http.Request) (interface{}, error) {
	ctx := r.Context()
	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		return nil, errors.NewInvalidInputData(fmt.Errorf("can't read body"))
	}

	data := new(paramStorageMove)
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, errors.NewInvalidInputData(fmt.Errorf("can't unmarshal body, %s", err))
	}

	res, err := s.service.MoveItemToCell(ctx, data.ProductId, data.SrcCellId, data.DstCellId, data.Quantity)
	if err != nil {
		return nil, errors.NewBadRequest(fmt.Errorf("data fetch error, %s", err))
	}
	return res, nil
}

func (s *StorageController) quantityProduct(r *http.Request) (interface{}, error) {
	ctx := r.Context()
	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		return nil, errors.NewInvalidInputData(fmt.Errorf("can't read body"))
	}

	data := new(paramStorage)
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, errors.NewInvalidInputData(fmt.Errorf("can't unmarshal body, %s", err))
	}

	res, err := s.service.GetItemFromCell(ctx, data.ProductId, data.CellId, data.Quantity)
	if err != nil {
		return nil, errors.NewBadRequest(fmt.Errorf("data fetch error, %s", err))
	}
	return res, nil
}
