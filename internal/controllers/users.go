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

type UsersController struct {
	service *service.WhsService
}

func NewUsersController(service *service.WhsService) *UsersController {
	return &UsersController{
		service: service,
	}
}

func (c *UsersController) RegisterHandlers(routeItems app.Routes) app.Routes {
	wrap := response.NewWrapper()
	routeItems = append(routeItems, app.Route{
		Name:          "GetUser",
		Method:        "GET",
		Pattern:       "/users/{itemId}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.DataPlain(c.getById),
	})
	routeItems = append(routeItems, app.Route{
		Name:          "GetUsers",
		Method:        "GET",
		Pattern:       "/users",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.DataPages(c.getUsers),
	})
	routeItems = append(routeItems, app.Route{
		Name:          "CreateUser",
		Method:        "POST",
		Pattern:       "/users",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.DataPlain(c.createUser),
	})
	routeItems = append(routeItems, app.Route{
		Name:          "UpdateUser",
		Method:        "PUT",
		Pattern:       "/users/{itemId}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.DataPlain(c.updateUser),
	})
	routeItems = append(routeItems, app.Route{
		Name:          "DeleteUser",
		Method:        "DELETE",
		Pattern:       "/users/{itemId}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.Empty(c.deleteUser),
	})
	routeItems = append(routeItems, app.Route{
		Name:          "GetSuggestionUsers",
		Method:        "GET",
		Pattern:       "/suggestion/users/{text}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.DataList(c.getSuggestion),
	})
	return routeItems
}

// @Summary get user by ID
// @Tags users
// @Success		200	{object} response.PlainData{data=User}
// @Param itemID path integer true "user ID"
// @Router /users/{itemID} [get]
func (c *UsersController) getById(r *http.Request) (interface{}, error) {
	var err error
	itemId, err := query.GetItemId(r)
	if err != nil {
		return nil, err
	}
	item, err := c.service.GetUserById(r.Context(), int64(itemId))
	if err != nil {
		return item, errors.NewServerError(err)
	}

	return item, nil
}

// @Summary get users list
// @Tags users
// @Success		200	{object} response.Pagination{data=[]User}
// @Param o query integer false "offset"
// @Param l query integer false "limit"
// @Router /users [get]
func (c *UsersController) getUsers(r *http.Request) (interface{}, *response.DataRange, error) {
	ctx := r.Context()
	offset, limit := query.GetOffsetLimit(r)
	items, count, err := c.service.GetUsers(ctx, offset, limit)
	if err != nil {
		return items, &response.DataRange{}, err
	}

	return items, &response.DataRange{
		Count:  int(count),
		Limit:  limit,
		Offset: offset,
	}, nil
}

// @Summary create user
// @Tags users
// @Param request body User true "user data"
// @Success		200	{object} response.PlainData{data=integer}
// @Router /users [post]
func (c *UsersController) createUser(r *http.Request) (interface{}, error) {
	ctx := r.Context()
	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		return nil, errors.NewInvalidInputData(fmt.Errorf("can't read body"))
	}

	data := new(model.User)
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, errors.NewInvalidInputData(fmt.Errorf("can't unmarshal body, %s", err))
	}

	userId, err := c.service.CreateUser(ctx, data)
	if err != nil {
		return nil, errors.NewBadRequest(fmt.Errorf("data fetch error, %s", err))
	}
	return userId, nil
}

// @Summary update user
// @Tags users
// @Param itemID path integer true "user ID"
// @Param request body User true "user data"
// @Success		200	{object} response.PlainData{data=integer}
// @Router /users/{itemID} [put]
func (c *UsersController) updateUser(r *http.Request) (interface{}, error) {
	ctx := r.Context()
	itemReqId, err := query.GetItemId(r)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		return nil, errors.NewInvalidInputData(fmt.Errorf("can't read body"))
	}

	data := new(model.User)
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, errors.NewInvalidInputData(fmt.Errorf("can't unmarshal body, %s", err))
	}
	data.Id = int64(itemReqId)
	userId, err := c.service.UpdateUser(ctx, data)
	if err != nil {
		return nil, errors.NewBadRequest(fmt.Errorf("data fetch error, %s", err))
	}
	return userId, nil
}

// @Summary delete user by ID
// @Tags users
// @Success		200	{object} map[string]interface{}
// @Param itemID path integer true "user ID"
// @Router /users/{itemID} [delete]
func (c *UsersController) deleteUser(r *http.Request) error {
	itemId, err := query.GetItemId(r)
	if err != nil {
		return err
	}

	err = c.service.DeleteUser(r.Context(), int64(itemId))
	if err != nil {
		return errors.NewServerError(fmt.Errorf("data fetch error, %s", err))
	}
	return nil
}

// @Summary get users suggestion
// @Tags users
// @Success		200	{object} response.List{data=[]User}
// @Param text path string true "text for suggestion"
// @Router /suggestion/users/{text} [get]
func (c *UsersController) getSuggestion(r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	if _, ok := vars["text"]; !ok {
		return nil, errors.NewInvalidInputData(fmt.Errorf("invalid path params"))
	}

	data, err := c.service.GetUserSuggestion(r.Context(), vars["text"], 10)
	if err != nil {
		return nil, errors.NewServerError(err)
	}
	return data, nil
}
