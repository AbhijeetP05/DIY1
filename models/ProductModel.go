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

func (*ProductModel) TableName() string {
	return "products"
}

func (p *ProductModel) GetProduct(db *gorm.DB) *gorm.DB {
	result := db.First(&p)

	log.Println("GetProduct executed" + strconv.FormatInt(result.RowsAffected, 10))
	return result
}

func (p *ProductModel) GetProducts(db *gorm.DB, limit, start int) ([]ProductModel, *gorm.DB) {
	var products []ProductModel
	result := db.Model(ProductModel{}).Offset(start).Limit(limit).Find(&products)

	log.Println("GetProducts executed" + strconv.FormatInt(result.RowsAffected, 10))
	return products, result
}

func (p *ProductModel) CreateProduct(db *gorm.DB) *gorm.DB {
	result := db.Create(&p)

	log.Println("GetProducts executed" + strconv.FormatInt(result.RowsAffected, 10))
	return result
}

func (p *ProductModel) updateProduct(db *gorm.DB, newProduct *ProductModel) *gorm.DB {
	result := db.Model(&p).Updates(newProduct)

	log.Println("GetProducts executed" + strconv.FormatInt(result.RowsAffected, 10))
	return result
}

func (p *ProductModel) deleteProduct(db *gorm.DB) *gorm.DB {
	result := db.Delete(&p)

	log.Println("GetProducts executed" + strconv.FormatInt(result.RowsAffected, 10))
	return result
}
