package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/heydp/oms/dto"
	"github.com/heydp/oms/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (h Handler) GetOrdersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Calling GetOrdersHandler func")

	// json.NewEncoder(w).Encode(mocks.Orders)
	var orders []models.Order
	result := h.DB.Preload(clause.Associations).Find(&orders)

	if result.Error != nil {
		log.Println("Error found in querying all orders from db")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Database Error")
		return
	}
	var outputs []dto.Order
	for _, val := range orders {
		var output dto.Order
		dto.Convert(&output, &val)
		outputs = append(outputs, output)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(outputs)
}
func (h Handler) GetOrderHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Calling GetOrderHandler func")
	w.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := vars["id"]

	intId, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Give an integer as id ")
		json.NewEncoder(w).Encode("Give an integer as id")
		return
	}

	// for i, val := range mocks.Orders {
	// 	if val.Id == intId {
	// 		w.WriteHeader(http.StatusOK)
	// 		json.NewEncoder(w).Encode(mocks.Orders[i])
	// 		return
	// 	}
	// }
	var order models.Order
	result := h.DB.Preload(clause.Associations).Find(&order, &models.Order{Id: intId})
	if result.Error != nil {
		log.Println("The Order Id is not present")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("The Order Id is not present")
		return
	}
	var output dto.Order
	dto.Convert(&output, &order)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)

}
func (h Handler) UpdateOrderHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Calling UpdateOrderHandler func")

	w.Header().Add("Content-Type", "application/json")

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error in Update request, ", err)
		return
	}

	var requestOrder models.Order
	err = json.Unmarshal(body, &requestOrder)
	if err != nil {
		log.Println("Error in Update request Unmarshallaing, ", err)
		return
	}

	vars := mux.Vars(r)
	id, _ := vars["id"]

	intId, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Give an integer as id ")
		json.NewEncoder(w).Encode("Give an integer as id")
		return
	}

	// for i, val := range mocks.Orders {
	// 	if val.Id == intId {
	// 		w.WriteHeader(http.StatusOK)
	// 		mocks.Orders[i].Status = requestOrder.Status
	// 		json.NewEncoder(w).Encode(mocks.Orders[i])
	// 		return
	// 	}
	// }
	var order models.Order
	result := h.DB.Preload(clause.Associations).Find(&order, &models.Order{Id: intId})
	if result.Error != nil {
		log.Println("The Order Id is not present")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("The Order Id is not present")
		return
	}
	h.DB.Model(&order).Update("status", requestOrder.Status)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Order status updated successfully")

}

// func (h Handler) DeleteOrderHandler(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Calling DeleteOrderHandler func")

// 	w.Header().Add("Content-Type", "application/json")

// 	vars := mux.Vars(r)
// 	id, _ := vars["id"]

// 	intId, err := strconv.Atoi(id)
// 	if err != nil {
// 		log.Println("Give an integer as id ")
// 		json.NewEncoder(w).Encode("Give an integer as id")
// 		return
// 	}

// 	for i, val := range mocks.Orders {
// 		if val.Id == intId {

// 			mocks.Orders = append(mocks.Orders[:i], mocks.Orders[i+1:]...)
// 			w.WriteHeader(http.StatusOK)
// 			json.NewEncoder(w).Encode("Deleted Successfully")
// 			return
// 		}
// 	}

// 	log.Println("Record not found. Please create a new Order ")
// 	json.NewEncoder(w).Encode("Record not found. Please create a new Order")

// }

func (h Handler) PostOrderHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Calling PostOrderHandler func")

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println("Error in Post request, ", err)
		return
	}
	// log.Println(string(body))
	var requestOrder models.Order
	err = json.Unmarshal(body, &requestOrder)
	if err != nil {
		log.Println("Error in Post request Unmarshallaing, ", err)
		return
	}

	rand.Seed(time.Now().UnixNano())
	requestOrder.Id = rand.Intn(100)
	// log.Println("The order id generated is : ", requestOrder.Id)
	// mocks.Orders = append(mocks.Orders, requestOrder)
	result := h.DB.Create(&requestOrder)
	if result.Error != nil {
		json.NewEncoder(w).Encode("Error in post req")
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Order created successfully")

}

type Handler struct {
	DB *gorm.DB
}

func NewHandler(db *gorm.DB) Handler {
	return Handler{
		DB: db,
	}
}
