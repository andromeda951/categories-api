package controller

import (
	"andromeda/belajar-golang-restful-api/helper"
	"andromeda/belajar-golang-restful-api/model/web"
	"andromeda/belajar-golang-restful-api/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

// Create implements CategoryController
func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Read request body
	decoder := json.NewDecoder(request.Body)
	categoryCreateRequest := web.CategoryCreateRequest{}
	err := decoder.Decode(&categoryCreateRequest)
	helper.PanicIfError(err)

	// Create response
	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	// Send response 
	writer.Header().Add("Content-Type", "aplication/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

// Update implements CategoryController
func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Read request body
	decoder := json.NewDecoder(request.Body)
	categoryUpdateRequest := web.CategoryUpadateRequest{}
	err := decoder.Decode(&categoryUpdateRequest)
	helper.PanicIfError(err)

	// Read dynamic categoryId from parameter
	categoryId := params.ByName("categoryId")		
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)
	categoryUpdateRequest.Id = id

	// Create response
	categoryResponse := controller.CategoryService.Upadate(request.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	// Send response 
	writer.Header().Add("Content-Type", "aplication/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}


// Delete implements CategoryController
func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Read dynamic categoryId from parameter
	categoryId := params.ByName("categoryId")		
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.CategoryService.Delete(request.Context(), id)
	
	// Create response
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	// Send response 
	writer.Header().Add("Content-Type", "aplication/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

// FindById implements CategoryController
func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Read dynamic categoryId from parameter
	categoryId := params.ByName("categoryId")		
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	
	// Create response
	categoryResponse := controller.CategoryService.FindByid(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data: categoryResponse,
	}

	// Send response 
	writer.Header().Add("Content-Type", "aplication/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

// FindAll implements CategoryController
func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Create response
	categoryResponses := controller.CategoryService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data: categoryResponses,
	}

	// Send response 
	writer.Header().Add("Content-Type", "aplication/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponse)
	helper.PanicIfError(err)
}


