package entities

import (


	"gopkg.in/mgo.v2/bson"
)
type Product struct{
	Id  			bson.ObjectId 	`json:"id" bson:"_id"`
	Recipe  		string 			`bson:"recipe_name" bson:"recipe_name`	
	Images  		string 			`bson:"images" bson:"images"`
	Chef 	 		string 			`bson:"chef_detail" bson:"chef_detail"`
	Cooking  		string 			`bson:"cooking_instruction" bson:"cooking_instruction"`
	Ingredients  	string 			`bson:"ingredients" bson:"ingredients"`
	Rating 	 		string 			`bson:"recipe_rating" bson:"recipe_rating"`
}