package handler

import (
	"encoding/json"
	"gabriellfe/dto"
	"gabriellfe/helper"
	"gabriellfe/validator"
	"net/http"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("live")
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var b dto.UserSearchDto
	if err := helper.DecodeAndValidate(w, r, &b); err != nil {
		return
	}
	if errors := validator.ValidateStruct(b); len(errors.Errors) > 0 {
		helper.EncodeStatusBody(w, errors, http.StatusBadRequest)
		return
	}
	helper.EncodeStatusBody(w, b.Status, http.StatusOK)
}
