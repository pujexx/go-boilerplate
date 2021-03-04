package http

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pujexx/go-boilerplate/domain"
	"github.com/pujexx/go-boilerplate/lib"
	"net/http"
	"strconv"
)

type ordersHttpHandler struct {
	Service domain.OrdersService
}

func NewOrdersHttpHandler(r *mux.Router, service domain.OrdersService) lib.BaseHandler {
	h := ordersHttpHandler{
		Service: service,
	}
	r.HandleFunc("/orders", h.Store).Methods("POST")
	r.HandleFunc("/orders/{id}", h.Delete).Methods("DELETE")
	r.HandleFunc("/orders", h.Update).Methods("PUT")
	r.HandleFunc("/orders", h.List).Methods("GET")
	r.HandleFunc("/orders/range", h.ListRange).Methods("GET")
	r.HandleFunc("/orders/{id}", h.Detail).Methods("GET")
	return &h
}

func (a ordersHttpHandler) Store(w http.ResponseWriter, r *http.Request) {
	var orders domain.Orders
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&orders); err != nil {
		fmt.Println(err)
		lib.BaseResponse(lib.ResponseError{
			Code:    "402",
			Message: "Unprocessable entity",
			Errors:  []lib.ValidateError{{Field: "json", Error: "invalid json format"}},
		}, w, r)
		return
	}

	if s, err := lib.ValidateStruct(orders); !s {
		lib.BaseResponse(lib.ResponseError{
			Code:    "402",
			Message: "Validation errors",
			Errors:  err,
		}, w, r)
		return
	}
	a.Service.Save(&orders)
	lib.BaseResponse(lib.Response{
		Code: "200",
		Data: "Success Created",
	}, w, r)
	return
}

func (a ordersHttpHandler) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	if s, err := lib.ValidateVar(id, "id", "required"); !s {
		lib.BaseResponse(lib.ResponseError{
			Code:    "402",
			Message: "Validation errors",
			Errors:  err,
		}, w, r)
		return
	}

	a.Service.Delete(id)
	lib.BaseResponse(lib.Response{
		Code: "200",
		Data: "Success Deleted",
	}, w, r)
}

func (a ordersHttpHandler) Update(w http.ResponseWriter, r *http.Request) {
	var orders domain.Orders
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&orders); err != nil {
		fmt.Println(err)
		lib.BaseResponse(lib.ResponseError{
			Code:    "402",
			Message: "Unprocessable entity",
			Errors:  []lib.ValidateError{{Field: "json", Error: "invalid json format"}},
		}, w, r)
		return
	}

	if s, err := lib.ValidateVar(orders.Id, "id", "required"); !s {
		lib.BaseResponse(lib.ResponseError{
			Code:    "402",
			Message: "Validation errors",
			Errors:  err,
		}, w, r)
		return
	}
	if s, err := lib.ValidateStruct(orders); !s {
		lib.BaseResponse(lib.ResponseError{
			Code:    "402",
			Message: "Validation errors",
			Errors:  err,
		}, w, r)
		return
	}
	a.Service.Update(&orders)
	lib.BaseResponse(lib.Response{
		Code: "200",
		Data: "Success Created",
	}, w, r)
	return
}

func (a ordersHttpHandler) List(w http.ResponseWriter, r *http.Request) {
	pageString := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		page = 1
	}
	_, b := a.Service.Find(page, []lib.FilterDomain{})

	lib.BaseResponse(lib.Response{
		Code: "200",
		Data: b.Data,
		Meta: b.Meta,
	}, w, r)
}

func (a ordersHttpHandler) ListRange(w http.ResponseWriter, r *http.Request) {
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")
	dateFrom := lib.DateFormat(from)
	dateTo := lib.DateFormat(to)

	if dateFrom == nil || dateTo == nil {
		lib.BaseResponse(lib.ResponseError{
			Code:    "402",
			Message: "Error Date Format",
		}, w, r)
		return
	}

	fmt.Println(dateFrom, dateTo, lib.DateRange(*dateFrom, *dateTo))
	err, b := a.Service.FindRange(*dateFrom, *dateTo)

	if err != nil {
		lib.BaseResponse(lib.ResponseError{
			Code:    "404",
			Message: err.Error(),
		}, w, r)
		return
	}

	lib.BaseResponse(lib.Response{
		Code: "200",
		Data: &b,
	}, w, r)
	return
}

func (a ordersHttpHandler) Detail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	if s, err := lib.ValidateVar(id, "id", "required"); !s {
		lib.BaseResponse(lib.ResponseError{
			Code:    "402",
			Message: "Validation errors",
			Errors:  err,
		}, w, r)
		return
	}
	object, err := a.Service.ByID(id)
	if err != nil {
		lib.BaseResponse(lib.ResponseError{
			Code:    "404",
			Message: "Object not found",
		}, w, r)
		return
	}
	lib.BaseResponse(lib.Response{
		Code: "200",
		Data: object,
	}, w, r)
}
