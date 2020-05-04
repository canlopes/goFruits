package http

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	m "goFruits/bundles"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)
var filepath = "data.csv"
func responseOk(fruit m.Fruit) string {
	return fmt.Sprintf(`{"fruit": "%s", "calories": %d, "price": %.2f }`, fruit.FruitName, fruit.Calories, fruit.Price)
}

// check if fruit name already exists
func checkFruit(fruitName string) m.Fruit {
	for _, fruit := range m.LoadCSV(filepath) {
		if fruit.FruitName == fruitName {
			return fruit
		}
	}
	return m.Fruit{}
}

// add a new fruit
func AddFruit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var f m.Fruit
	err := json.NewDecoder(r.Body).Decode(&f)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"invalid body"}`))
		return
	}
	var fruitExists = checkFruit(f.FruitName)
	if (m.Fruit{}) == fruitExists {
		v := validator.New()
		errs := v.Struct(f)
		if errs != nil {
			var errLog = ""
			for _, e := range errs.(validator.ValidationErrors) {
				errLog = errLog + " " + e.Field()
			}
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf(`{"message":"invalid parameters%s"}`,
				errLog)))
			return
		}
		var fruits = m.LoadCSV(filepath)
		fruits = append(fruits,f)
		m.WriteCSV(filepath, fruits)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(responseOk(f)))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"fruit already exist"}`))
	}
}
//update existing fruit
func UpdateFruit(w http.ResponseWriter, r *http.Request) {

	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	fruitName := ""
	if val, ok := pathParams["fruitName"]; ok {
		fruitName = val
		if len(fruitName) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"message": "invalid fruit name"}`))
			return
		}
	}
	var fruit m.Fruit
	err := json.NewDecoder(r.Body).Decode(&fruit)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"invalid body request"}`))
		return
	}
	var fruitExists = checkFruit(fruitName)
	if (m.Fruit{}) == fruitExists {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{"message":"fruit not found %s"}`, fruitName)))
	} else if fruitExists == fruit {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"nothing to change"}`))
	} else {
		var fruits = m.LoadCSV(filepath)
		for i, record := range fruits {
			if record.FruitName == fruitName {
				v := validator.New()
				errs := v.Struct(fruit)
				if errs != nil {
					var errLog = ""
					for _, e := range errs.(validator.ValidationErrors) {
						errLog = errLog + " " + e.Field()
					}
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte(fmt.Sprintf(`{"message":"invalid parameters%s"}`,
						errLog)))
					return
				}

				fruits[i] = fruit
				w.WriteHeader(http.StatusAccepted)
				w.Write([]byte(responseOk(fruits[i])))
				m.WriteCSV(filepath, fruits)
				return
			}
		}
	}
}
//return all fruits
func GetAllFruits(w http.ResponseWriter, r *http.Request) {
	allFruits, _ := json.Marshal(m.LoadCSV(filepath))
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(string(allFruits))))
}
//get fruit by name
func GetFruit(w http.ResponseWriter, r *http.Request) {

	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	fruitName := ""
	if val, ok := pathParams["fruitName"]; ok {
		fruitName = val
		if len(fruitName) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"message": "invalid fruit name"}`))
			return
		}
	}
	var fruitExists = checkFruit(fruitName)
	if (m.Fruit{}) == fruitExists {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{"message":"fruit not found %s"}`, fruitName)))
	} else {
		w.Write([]byte(responseOk(fruitExists)))
	}
}
//delete fruit by name
func DeleteFruit(w http.ResponseWriter, r *http.Request) {
	m.LoadCSV(filepath)

	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	fruitName := ""
	if val, ok := pathParams["fruitName"]; ok {
		fruitName = val
		if len(fruitName) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"message": "invalid fruit name"}`))
			return
		}
	}
	var fruits = m.LoadCSV(filepath)
	for i, fruit := range fruits {
		if fruit.FruitName == fruitName {
			fruits = append(fruits[:i], fruits[i+1:]...)
			w.Write([]byte(responseOk(fruit)))
			m.WriteCSV(filepath, fruits)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(fmt.Sprintf(`{"message":"fruit not found %s"}`, fruitName)))
}
