package models

import (
    "gorm.io/gorm"
	"strconv"

)


type User struct {
    gorm.Model
    Name     string
    Post   []Post
	Post_Count int64
}

type Post struct {
    gorm.Model
    Title    string
    Content    string
	Comment   []Comment
	Comment_Count string
    UserID    uint
}

func(p *Post) AfterCreate(tx *gorm.DB) (err error) {
	var count int64
	tx.Model(&Post{}).Where("user_id = ?", p.UserID).Count(&count)
	tx.Model(&User{}).Where("id = ?", p.UserID).Update("post_count", count)
	return
}

type Comment struct {
    gorm.Model
    Content    string
    PostID    uint
}

func(c *Comment) AfterCreate(tx *gorm.DB) (err error) {
	var count int64
	tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count)
	var comment_count string
	if count == 0 {
		comment_count = "无评论"
	} else {
		comment_count = strconv.FormatInt(count, 10)
	}

	tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_count", comment_count)
	return
}

func(c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var count int64
	tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count)
	var comment_count string
	if count == 0 {
		comment_count = "无评论"
	} else {

		comment_count = string(count)
	}
	tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_count", comment_count)
	return
}

func Run (db *gorm.DB) {

    db.AutoMigrate(&User{})
    db.AutoMigrate(&Post{})
    db.AutoMigrate(&Comment{})

	users := []User{
		{Name: "Alice"},
		{Name: "Bob"},
		{Name: "Charlie"},
	}
	posts := []Post{
		{Title: "Alice's First Post", Content: "Hello, this is Alice!", UserID: 1},
		{Title: "Alice's Second Post", Content: "Hello, this is Alice again!", UserID: 1},
		{Title: "Bob's First Post", Content: "Hello, this is Bob!", UserID: 2},
		{Title: "Charlie's First Post", Content: "Hello, this is Charlie!", UserID: 3},
	}
	comments := []Comment{
		{Content: "Great post, Alice!", PostID: 1},
		{Content: "Funny, Alice!", PostID: 2},
		{Content: "Thanks for sharing, Bob!", PostID: 3},
		{Content: "Interesting read, Charlie!", PostID: 4},
	}
	db.Create(&users)
	db.Create(&posts)
	db.Create(&comments)




}

