package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	app "github.com/mlplabs/app-utils"
	"github.com/mlplabs/app-utils/pkg/http/errors"
	"github.com/mlplabs/app-utils/pkg/http/response"
	"github.com/mlplabs/microwms/internal/usecase/service"
	"net/http"
)

type CellsController struct {
	service *service.WhsService
}

func NewCellsController(service *service.WhsService) *CellsController {
	return &CellsController{
		service: service,
	}
}

func (c *CellsController) RegisterHandlers(routeItems app.Routes) app.Routes {
	wrap := response.NewWrapper()

	routeItems = append(routeItems, app.Route{
		Name:          "GetSuggestionCells",
		Method:        "GET",
		Pattern:       "/suggestion/cells/{text}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.DataList(c.getSuggestion),
	})
	return routeItems
}

// @Summary get manufacturers suggestion
// @Tags manufacturers
// @Success		200	{object} response.List{data=[]Cell}
// @Param text path string true "text for suggestion"
// @Router /suggestion/cells/{text} [get]
func (c *CellsController) getSuggestion(r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	if _, ok := vars["text"]; !ok {
		return nil, errors.NewInvalidInputData(fmt.Errorf("invalid path params"))
	}

	data, err := c.service.GetCellSuggestion(r.Context(), vars["text"], 10)
	if err != nil {
		return nil, errors.NewServerError(err)
	}
	return data, nil
}
