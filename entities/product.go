package entities

import (
	"gopkg.in/mgo.v2/bson"
)
type Product struct{
	Id  				bson.ObjectId 			`json:"id"`
	Recipe 	 			string 					`json:"recipe_name"`
	Cooking  			string 					`json:"cooking_instruction"`
	
	Chef 	 			[]ChefData 				`json:"chef_detail"`
	Rating  			[]RatingData   			`json:"recipe_rating"`
	Images  			[]ImagesDate 		    `json:"images"`
	Ingredients  		[]IngredData 			`json:"ingredients"`
}

type ChefData struct{
	Url  				string 					`bson:"url" bson:"url"`
	File_Name 	 		string 					`bson:"file_name" bson:"file_name"`
	Chef_Name 	 		string 					`bson:"chef_name" bson:"chef_name"`
}
type RatingData struct{	
	Rating				string 					`bson:"rating" bson:"rating"`
	Rating_Date     	string 					`bson:"rating_date" bson:"rating_date"`
	Rating_Time			string 					`bson:"rating_time" bson:"rating_time"`
	Id_User				string 					`bson:"id_user" bson:"id_user"`
}	
type ImagesDate struct{
	Url					string 					`bson:"url" bson:"url"`
	File_Name     		string 					`bson:"file_name" bson:"file_name"`
}

type IngredData struct{
	Id_Ingred			string 					`bson:"id_ingredients" bson:"id_ingredients"`
	Ingred_Name     	string 					`bson:"ingredients_name" bson:"ingredients_name"`
	Quantity			string 					`bson:"quantity" bson:"quantity"`
	Unit       			string 					`bson:"unit" bson:"unit"`
}