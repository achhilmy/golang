package entities

import (
	"gopkg.in/mgo.v2/bson"
)
type Product struct{
	Id  				bson.ObjectId 			`json:"id" bson:"_id"`
	Recipe 	 			string 					`bson:"recipe_name" bson:"recipe_name"`
	Cooking  			string 					`bson:"cooking_instruction" bson:"cooking_instruction"`
	
	Chef 	 			[]ChefData 				`bson:"chef_detail" bson:"chef_detail"`
	Rating  			[]RatingData   			`bson:"recipe_rating" bson:"recipe_rating"`
	Images  			[]ImagesDate 		    `bson:"images" bson:"images"`
	Ingredients  		[]IngredData 			`bson:"ingredients" bson:"ingredients"`
}

type ChefData struct{
	Url  				string 					`json:url`
	File_Name 	 		string 					`json:file_name`
	Chef_Name 	 		string 					`json:chef_name`
}
type RatingData struct{	
	Rating				string 					`json:rating`
	Rating_Date     	string 					`json:rating_date`
	Rating_Time			string 					`json:rating_time`
	Id_User				string 					`json:id_user`
}	
type ImagesDate struct{
	Url					string 					`json:url`
	File_Name     		string 					`json:file_name`
}

type IngredData struct{
	Id_Ingred			string 				`json:id_ingredients`
	Ingred_Name     	string 				`json:ingredients_name`
	Quantity			string 				`json:quantity`
	Unit       			string 				`json:unit`
}