/*
 *
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/project-alvarium/go-simulator/api/models"
	"github.com/project-alvarium/go-simulator/handlers"

	"github.com/gorilla/mux"
)

// AddInsight adds insight to db and returns response to user
func AddInsight(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	body, err := ioutil.ReadAll(r.Body)
	var data models.Data
	json.Unmarshal([]byte(body), &data)
	if err != nil {
		json.NewEncoder(w).Encode(models.Response{Code: 400, Message: err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, err := handlers.CreateInsight(data)
	if err != nil {
		json.NewEncoder(w).Encode(models.Response{Code: 400, Message: err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(models.Response{Code: 200, Message: res})
	w.WriteHeader(http.StatusOK)
	return
}

// GetInsightByID gets insight from db and returns response to user
func GetInsightByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	res, err := handlers.GetInsight(vars["id"])
	if err != nil {
		json.NewEncoder(w).Encode(models.Response{Code: 400, Message: err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(models.Insight{InsightID: res.InsightID, DataID: res.DataID})
	w.WriteHeader(http.StatusOK)
	return
}
