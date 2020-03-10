package entities




type CodeStatus struct{
	Status		string 		 			`bson:"status" bson:"status"`
	Message 	string  	 			`bson:"message" bson:"message"`
	ErrorCode	int32		 			`bson:"code" bson:"code"`

	Data		map[string]interface {}  `bson:"Product"`
} 

