package models

import (
	"time"
	// "log"
	// "database/sql"
)

type BlogComment struct {
	Index int
	ForBlogId int
	Owner *Account
	WriteTime time.Time
	Content string
}

type BlogEntry struct {
	Id int
	Owner *Account
	WriteTime time.Time
	Title string
}

func makeBlog(id int, owner *Account, writeTime time.Time, title string) BlogEntry {
	var blog BlogEntry
	
	blog.Id = id
	blog.Owner = owner
	blog.WriteTime = writeTime
	blog.Title = title
	
	return blog
}

func GetAllBlogEntries() []BlogEntry {
	return [] BlogEntry {
		makeBlog(1, nil, time.Now(), "this is test01"),
		makeBlog(2, nil, time.Now(), "this is test02"),
	}
}