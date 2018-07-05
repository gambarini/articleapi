package model

import "time"

type (
	Article struct {
		ID       string    `json:"id" bson:"_id,omitempty"`
		Title    string    `json:"title"`
		Date     string    `json:"date"`
		DateTime time.Time `json:"-"`
		Body     string    `json:"body"`
		Tags     []string  `json:"tags"`
	}

	Tag struct {
		Tag         string   `json:"tag"`
		Count       int      `json:"count"`
		Articles    []string `json:"articles"`
		RelatedTags []string `json:"related_tags"`
	}
)
