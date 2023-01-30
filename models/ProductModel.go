package models

import (
	"gorm.io/gorm"
	"log"
	"strconv"
)

type ProductModel struct {
	ID    uint    `json:"id" gorm:"primaryKey"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

//type IProductModel interface {
//	GetProduct()
//}
//
//func NewProductModel(ID uint, Name string, Price float64) *ProductModel {
//	return &ProductModel{ID: ID, Name: Name, Price: Price}
//}

func (p *ProductModel) GetProduct(db *gorm.DB) *ProductModel {
	result := db.First(&p)
	if result == nil {
		return nil
	}
	log.Println("GetProduct executed" + strconv.FormatInt(result.RowsAffected, 10))
	return p
}

func (p *ProductModel) GetProducts(db *gorm.DB, limit, start int) []ProductModel {
	var products []ProductModel
	result := db.Model(ProductModel{}).Offset(start).Limit(limit).Find(&products)
	if result == nil {
		return nil
	}
	log.Println("GetProducts executed" + strconv.FormatInt(result.RowsAffected, 10))
	return products
}

func (p *ProductModel) CreateProduct(db *gorm.DB) *ProductModel {
	result := db.Create(&p)
	if result == nil {
		return nil
	}
	log.Println("GetProducts executed" + strconv.FormatInt(result.RowsAffected, 10))
	return p
}

func (p *ProductModel) updateProduct(db *gorm.DB, newProduct *ProductModel) *ProductModel {
	result := db.Model(&p).Updates(newProduct)
	if result == nil {
		return nil
	}
	log.Println("GetProducts executed" + strconv.FormatInt(result.RowsAffected, 10))
	return p
}

func (p *ProductModel) deleteProduct(db *gorm.DB) bool {
	result := db.Delete(&p)
	if result == nil {
		return false
	}
	log.Println("GetProducts executed" + strconv.FormatInt(result.RowsAffected, 10))
	return true
}
