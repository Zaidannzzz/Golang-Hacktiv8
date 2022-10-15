package controllers

import (
	"Assignment2/httpserver/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var db *sql.DB

type response struct {
	Order_ID int64	`json: "Order_ID"`
	Message string 	`json: "Message"`
}

type Response struct {
	Status  int           		`json:"status"`
	Message string        		`json:"message"`
	Data    []models.TheOrder	`json:"data"`
}

var order []models.TheOrder

//index
func GetTheOrder(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	order, err := models.GetTheOrder()
	if err != nil {
		log.Fatalf("Tidak bisa mengambil data. %v", err)
	}

	var response Response
	response.Status = 1
	response.Message = "Success"
	response.Data = order

	json.NewEncoder(w).Encode(response)
}

//show single order
func GetOrder(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	Order_ID, err := strconv.Atoi(params["Order_ID"])
	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", err)
	}

	order, err := models.GetOrder(int(Order_ID))
	if err != nil {
		log.Fatalf("Tidak bisa mengambil data. %v", err)
	}

	json.NewEncoder(w).Encode(order)
}

//add single order
func CreateOrder(w http.ResponseWriter, r *http.Request)  {
	var order models.TheOrder
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		log.Fatalf("Tidak bisa Request Body. %v", err)
	}
	addOrder_ID := models.CreateOrder(order)
	res := response{
		Order_ID: addOrder_ID,
		Message: "Data telah ditambahkan",
	}

	json.NewEncoder(w).Encode(res)
}

//update single order
func UpdateOrder(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["Order_ID"])
	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", err)
	}

	var order models.TheOrder
	err = json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		log.Fatalf("Tidak bisa decode request body.  %v", err)
	}

	updatedRows := models.UpdateOrder(int64(id), order)
	msg := fmt.Sprintf("Buku telah berhasil diupdate. Jumlah yang diupdate %v rows/record", updatedRows)

	res := response{
		Order_ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

//delete single order
func DeleteOrder(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)

	Order_ID, err := strconv.Atoi(params["Order_ID"])
	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", err)
	}

	deletedRows := models.DeleteOrder(int64(Order_ID))
	msg := fmt.Sprintf("Order sukses di hapus. Total data yang dihapus %v", deletedRows)
	res := response{
		Order_ID:      int64(Order_ID),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}