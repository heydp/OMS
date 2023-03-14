package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/heydp/oms/dto"
	"github.com/heydp/oms/pkg/models"
	"gorm.io/gorm/clause"
)

func (uh Handler) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	log.Println("Calling GetOrdersHandler func")

	//var orders []models.Order
	orders := []models.Order{}
	// sortBy is expected to look like field.orderdirection i. e. id.asc
	sortBy := r.URL.Query().Get("sortBy")
	if sortBy == "" {
		// id.asc is the default sort query
		sortBy = "id.asc"
	}
	sortQuery, err := validateAndReturnSortQuery(sortBy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := uh.DB.Preload(clause.Associations).Order(sortQuery).Find(&orders).Error; err != nil {
		fmt.Println(err)
		http.Error(w, "Error on DB find for all users", http.StatusInternalServerError)
		return
	}
	var outputs []dto.Order
	for _, val := range orders {
		var output dto.Order
		dto.Convert(&output, &val)
		outputs = append(outputs, output)
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(outputs); err != nil {
		// if err := json.NewEncoder(w).Encode(orders); err != nil {

		fmt.Println(err)
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
	}
}

var userFields = getUserFields()

func getUserFields() []string {
	var field []string
	v := reflect.ValueOf(models.Order{})
	for i := 0; i < v.Type().NumField(); i++ {
		field = append(field, v.Type().Field(i).Tag.Get("json"))
	}
	return field
}
func stringInSlice(strSlice []string, s string) bool {
	for _, v := range strSlice {
		if v == s {
			return true
		}
	}
	return false
}

func validateAndReturnSortQuery(sortBy string) (string, error) {
	splits := strings.Split(sortBy, ".")
	if len(splits) != 2 {
		return "", errors.New("malformed sortBy query parameter, should be field.orderdirection")
	}
	field, order := splits[0], splits[1]
	if order != "desc" && order != "asc" {
		return "", errors.New("malformed orderdirection in sortBy query parameter, should be asc or desc")
	}
	if !stringInSlice(userFields, field) {
		return "", errors.New("unknown field in sortBy query parameter")
	}
	return fmt.Sprintf("%s %s", field, strings.ToUpper(order)), nil
}
