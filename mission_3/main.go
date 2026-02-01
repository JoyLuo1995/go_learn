package main

import (
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
  "go_learn/mission_3/query"
  "go_learn/mission_3/models"
  "fmt"
)

func main() {
  dsn := "root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  // 创建表结构
  models.Run(db)

  // 查询某个用户发布的所有文章及其对应的评论信息
  userID := uint(1) // 替换为你想查询的用户ID
  posts, comments, err := query.QueryPostsAndCommentsByUserID (db, userID)
  if err != nil {
    panic(err)
  }
  fmt.Printf("Posts by User ID %d:\n", userID)
  fmt.Println(posts)
  fmt.Println("Comments for these Posts:\n", comments)

  // 删除某篇文章的所有评论，在评论删除时检查文章的评论数量，
  // 如果评论数量为 0，则更新文章的评论状态为 "无评论"。
  var delete_comment []models.Comment
  find_err := db.Where("id = ?", 4).Find(&delete_comment).Error
  if find_err != nil {
    panic(find_err)
  }else {
    for _, comment := range delete_comment {
      db.Delete(&comment)
    }
  }



}