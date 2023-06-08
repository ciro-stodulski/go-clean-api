package dto

type (
	Event struct {
		Id    string `json:"id" bson:"_id,omitempty"`
		Name  string `json:"name"  bson:"name"`
		Event string `json:"event" bson:"event"`
	}
)
