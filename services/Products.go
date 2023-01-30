package services

import (
	"github.com/gorilla/mux"
	"go-mux/models"
	"go-mux/utils"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

type Products struct {
	conn *gorm.DB
}

type IProducts interface {
	GetProduct(w http.ResponseWriter, r *http.Request)
	GetProducts(w http.ResponseWriter, r *http.Request)
	createProduct(w http.ResponseWriter, r *http.Request)
	updateProduct(w http.ResponseWriter, r *http.Request)
	deleteProduct(w http.ResponseWriter, r *http.Request)
}

func (p *Products) GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		log.Fatal("Invalid Product ID")
	}
	productModel := models.ProductModel{ID: uint(id)}
	result := productModel.GetProduct(p.conn)

	if result.Error == nil && result.RowsAffected == 0 {
		utils.RespondWithError(w, http.StatusNotFound, "Product Not Found")
	} else if result.Error != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, result.Error.Error())
	} else {
		utils.RespondWithJSON(w, http.StatusOK, productModel)
	}
}

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.Atoi(r.FormValue("limit"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	productModel := models.ProductModel{}
	products, result := productModel.GetProducts(p.conn, limit, start)
	if result.Error != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, result.Error.Error())
	} else {
		utils.RespondWithJSON(w, http.StatusOK, products)
	}

}

func (p *Products) createProduct(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (p *Products) updateProduct(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (p *Products) deleteProduct(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func NewProduct(conn *gorm.DB) *Products {
	return &Products{conn}
}
