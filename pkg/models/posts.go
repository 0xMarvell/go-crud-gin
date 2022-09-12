package models

import "gorm.io/gorm"

// Post is the database model for blog posts.
type Post struct {
	gorm.Model
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author string `json:"author"`
}
