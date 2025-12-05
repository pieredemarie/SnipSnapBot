package models

type User struct {
    ID        int    `json:"id" bson:"_id"`       // Telegram ID
    Username  string `json:"username" bson:"username,omitempty"`
    FirstName string `json:"first_name" bson:"first_name,omitempty"`
    LastName  string `json:"last_name" bson:"last_name,omitempty"`
}
