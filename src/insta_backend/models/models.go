package models
type User struct{
	ID interface{} `json:"id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type Post struct{
	ID interface{} `json:"id,omitempty" bson:"_id,omitempty"`
	UserId interface{} `json:"userId,omitempty" bson:"userId,omitempty"`
	Caption string `json:"caption" bson:"caption"`
	ImageUrl string `json:"imageUrl" bson:"imageUrl"`
	Posted string `json:"postTime" bson:"postTime"`
}