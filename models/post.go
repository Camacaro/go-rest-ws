package models

import "time"

type Post struct {
	Id          string    `json:"id"`
	UserId      string    `json:"user_id"`
	PostContent string    `json:"pos_content"`
	CreatedAt   time.Time `json:"created_at"`
}
