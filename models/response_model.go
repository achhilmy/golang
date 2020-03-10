package models

import (
	"crudmongomux2/entities"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CodeStatusModel struct {
	Db *mgo.Database
}

func (codeStatusModel CodeStatusModel) Getdata(id string) (entities.CodeStatus, error) {
	var codestatus entities.CodeStatus
	err := codeStatusModel.Db.C("status").FindId(bson.ObjectIdHex(id)).One(&codestatus)
	return codestatus, err
}

