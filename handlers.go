package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Handlers holds the router, db and helpers
type Handlers struct {
	router *mux.Router
	db     *sql.DB
	he     *Helper
}

// setRoutes creates the router and routes
func (h *Handlers) setRoutes() {
	h.router = mux.NewRouter()

	// Create coupons
	h.router.HandleFunc("/coupon/create", h.createNewCoupon).
		Queries("name", "{name}").
		Queries("brand", "{brand}").
		Queries("value", "{value}").
		Queries("expiry", "{expiry}"). // format: `2020-07-14 17:51:55`
		Methods("POST")

	// Get a coupon by id
	h.router.HandleFunc("/coupon/get/{id:[0-9]+}", h.getCoupons).Methods("GET")
	// Get a list of all coupons
	h.router.HandleFunc("/coupon/get/{id:all}", h.getCoupons).Methods("GET")
	// Get a limited number of coupon with limit and offset
	h.router.HandleFunc("/coupon/get", h.getCoupons).
		Queries("limit", "{limit}").
		Queries("offset", "{offset}").
		Methods("GET")

	//Update a coupon by id
	h.router.HandleFunc("/coupon/update/{id:[0-9]+}", h.updateCoupon).
		Queries("name", "{name}").
		Queries("brand", "{brand}").
		Queries("value", "{value}").
		Queries("expiry", "{expiry}"). // format: `2020-07-14 17:51:55`
		Methods("PUT")
}

// createNewCoupon Creates a new coupon entry in the DB
func (h *Handlers) createNewCoupon(w http.ResponseWriter, r *http.Request) {

	// Extract the vars
	vars := mux.Vars(r)

	// Parse number
	var value float64
	value, err := strconv.ParseFloat(vars["value"], 64)
	if err != nil {
		log.Println(err)
		writeError(w, http.StatusBadRequest, err)
		return
	}

	// Parse date
	expiry, err := time.Parse("2006-01-02 15:04:05", vars["expiry"])
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	// Creating the coupon in the DB
	insertedID, err := h.he.createNewCoupon(h.db, vars["name"], vars["brand"], value, expiry.Format("2006-01-02 15:04:05"))
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	writeJSON(w, http.StatusCreated, map[string]interface{}{"status": "Coupon created", "couponID": insertedID})
}

// getCoupons will get a list of coupons, one coupon by id or by limit and offset
func (h *Handlers) getCoupons(w http.ResponseWriter, r *http.Request) {
	// Extract the vars
	vars := mux.Vars(r)
	// vars holders
	var id, limit, offset int
	var err error
	if vars["id"] == "all" {
		// if we want to get the list of all the coupons
		id = 0
	} else if vars["id"] != "" {
		// if we want to get one coupon
		id, err = strconv.Atoi(vars["id"])
		if err != nil {
			log.Println(err)
			writeError(w, http.StatusBadRequest, err)
			return
		}
	}
	// Get the limit
	limit, err = strconv.Atoi(vars["limit"])
	if err != nil {
		limit = 0
	}
	// Get the offset
	offset, err = strconv.Atoi(vars["offset"])
	if err != nil {
		limit = 0
	}

	coupons, err := h.he.getCoupons(h.db, id, limit, offset)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		log.Println(err)
		return
	}
	if coupons == nil {
		writeError(w, http.StatusNotFound, errors.New("Coupon not found"))
		log.Println(err)
		return
	}
	writeJSON(w, http.StatusCreated, map[string]interface{}{"status": "Coupon(s) received from DB", "couponList": coupons})
}

// updateCoupon updates a couon in the DB
func (h *Handlers) updateCoupon(w http.ResponseWriter, r *http.Request) {

	// Extract the vars
	vars := mux.Vars(r)

	// Parse id
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		writeError(w, http.StatusBadRequest, err)
		return
	}

	// Parse number
	var value float64
	value, err = strconv.ParseFloat(vars["value"], 64)
	if err != nil {
		log.Println(err)
		writeError(w, http.StatusBadRequest, err)
		return
	}

	// Parse date
	expiry, err := time.Parse("2006-01-02 15:04:05", vars["expiry"])
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	// Creating the coupon in the DB
	rowsAffected, err := h.he.updateCoupon(h.db, id, vars["name"], vars["brand"], value, expiry.Format("2006-01-02 15:04:05"))
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	writeJSON(w, http.StatusCreated, map[string]interface{}{"status": "Coupon updated", "rowsAffected": rowsAffected})
}

// writeJSON will create a response
func writeJSON(w http.ResponseWriter, statusCode int, rawResponse interface{}) {
	response, err := json.Marshal(rawResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte(response))
}

// writeError will create an error response
func writeError(w http.ResponseWriter, statusCode int, err error) {
	writeJSON(w, statusCode, map[string]string{"error": err.Error()})
}
