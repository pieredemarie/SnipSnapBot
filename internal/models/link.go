package models

type Link struct {
	URL      string   `json:"url" bson:"url"`
	Tags     []string `json:"tags" bson:"tags"`
	AuthorId int      `json:"author_id" bson:"author_id"`
	Created  int64    `json:"created,omitempty" bson:"created,omitempty"`
}
