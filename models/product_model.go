package models

import (
	"crudmongomux2/entities"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ProductModel struct {
	Db *mgo.Database
}

func (productModel ProductModel) FindAll() ([]entities.Product, error) {
	var products []entities.Product
	err := productModel.Db.C("cooking_studio").Find(bson.M{}).All(&products)
	if err != nil {
		return nil, err
	} else {
		return products, nil
	}
}

func (productModel ProductModel) Find(id string) (entities.Product, error) {
	var product entities.Product
	err := productModel.Db.C("cooking_studio").FindId(bson.ObjectIdHex(id)).One(&product)
	return product, err
}

func (productModel ProductModel) Create(product *entities.Product) error {
	return productModel.Db.C("cooking_studio").Insert(&product)
}

func (productModel ProductModel) Delete(id string) error {
	return productModel.Db.C("cooking_studio").RemoveId(bson.ObjectIdHex(id))
}

func (productModel ProductModel) Update(product *entities.Product) error {
	return productModel.Db.C("cooking_studio").UpdateId(product.Id, product)
}
