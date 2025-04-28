package query

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mlplabs/app-utils/pkg/http/errors"
	"net/http"
	"strconv"
)

func GetOffsetLimit(r *http.Request) (int, int) {
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
	return offset, limit
}

func GetOwner(r *http.Request) (int, string, error) {
	vars := mux.Vars(r)
	ownerId, err := strconv.Atoi(vars["ownerId"])
	if err != nil {
		return 0, "", errors.NewInvalidInputData(fmt.Errorf("invalid query params"))
	}
	ownerRef := vars["ownerRef"]
	return ownerId, ownerRef, nil
}

func GetItemId(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	if v, ok := vars["itemId"]; !ok || v == "0" {
		return 0, errors.NewInvalidInputData(fmt.Errorf("invalid path params"))
	}

	itemId, err := strconv.Atoi(vars["itemId"])
	if err != nil {
		return 0, errors.NewInvalidInputData(fmt.Errorf("invalid query params"))
	}
	return itemId, nil
}
