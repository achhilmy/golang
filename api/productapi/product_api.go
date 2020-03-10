package productapi

import (
	"crudmongomux2/config"
	"crudmongomux2/entities"
	"crudmongomux2/models"
	"encoding/json"
	
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)
func Getdata(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()

	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
		return
	} else {
		codeStatusModel := models.CodeStatusModel{
			Db: db,
		}
		vars := mux.Vars(request)
		id := vars["id"]
		codestatus, err5 := codeStatusModel.Getdata(id)
		if err5 != nil {
			respondWithError(response, http.StatusBadRequest, err5.Error())
			return
		} else {
			respondWithJson(response, http.StatusOK, codestatus)
		}

	}
}

func FindAll(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()

	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
		return
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err2 := productModel.FindAll()
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
			return
		} else {
			respondWithJson(response, http.StatusOK, products)
		}

	}
}
func Find(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()

	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
		return
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		vars := mux.Vars(request)
		id := vars["id"]
		product, err2 := productModel.Find(id)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
			return
		} else {
			respondWithJson(response, http.StatusOK, product)
		}

	}
}

func Create(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()

	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
		return
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		var product entities.Product
		product.Id = bson.NewObjectId()
		json.NewDecoder(request.Body).Decode(&product)
		err2 := productModel.Create(&product)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, product)
		}

	}
}
func Delete(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()

	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
		return
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		vars := mux.Vars(request)
		id := vars["id"]
		err2 := productModel.Delete(id)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, entities.Product{})
		}

	}
}

func Update(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
		return
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		var product entities.Product
		json.NewDecoder(request.Body).Decode(&product)
		err2 := productModel.Update(&product)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, product)
		}
	}

}
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
