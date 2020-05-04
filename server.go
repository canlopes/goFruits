package main

import (
	"github.com/gorilla/mux"
	h "goFruits/transport/http"
	"log"
	"net/http"
)
/*
func init() {
	//m.Fruits = append(m.Fruits, m.Fruit{FruitName:"orange", Calories: 123, Price: 0.55})
	//m.Fruits = append(m.Fruits, m.Fruit{FruitName:"apple", Calories: 100, Price: 0.85})
	m.Fruits = m.LoadCSV()
	fmt.Println("CSV loaded", m.Fruits)
}
*/
func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/fruit/{fruitName}", h.GetFruit).Methods(http.MethodGet)
	api.HandleFunc("/fruit/{fruitName}", h.DeleteFruit).Methods(http.MethodDelete)


	api.HandleFunc("/fruit", h.AddFruit).Methods(http.MethodPost)


	api.HandleFunc("/fruit/{fruitName}", h.UpdateFruit).Methods(http.MethodPut)
	api.HandleFunc("/fruits", h.GetAllFruits).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", r))
}
