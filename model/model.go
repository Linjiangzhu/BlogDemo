package model

import (
	"database/sql"

	// not get golint complain
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/xid"
)

// BlogPost is the container of blog post from sql
type BlogPost struct {
	ID      string
	Title   string
	Content string
	Time    string
}

// BlogPostPageData is for constrcting blog page
type BlogPostPageData struct {
	PageTitle string
	Posts     []BlogPost
}

var db *sql.DB = InitDB()

// RetrivePost return the posts from server
func RetriveAllPost() BlogPostPageData {
	rows, err := db.Query("SELECT * FROM posts;")
	if err != nil {
		panic(err.Error())
	}

	var postslice []BlogPost
	for rows.Next() {
		var id string
		var title string
		var content string
		var timestamp string
		rows.Scan(&id, &title, &content, &timestamp)
		postslice = append(postslice, BlogPost{ID: id, Title: title, Content: content, Time: timestamp})
	}
	data := BlogPostPageData{
		PageTitle: "博文页面",
		Posts:     postslice,
	}

	return data
}

// CreatePost insert a new post row into databases and return true if succeed
func CreatePost(postTitle, postContent string) error {
	postID := xid.New()
	_, err := db.Exec("INSERT INTO posts(id, title, content) VALUES(?, ?, ?);", postID, postTitle, postContent)
	return err
}

func RetrivePostByID(postid string) (BlogPost, error) {
	var title string
	var content string
	var time string
	row := db.QueryRow("SELECT title, content, submission_time FROM posts WHERE id=?", postid)
	err := row.Scan(&title, &content, &time)
	if err != nil {
		return BlogPost{}, err
	}
	return BlogPost{ID: postid, Title: title, Content: content, Time: time}, nil
}

func DeletePost(postid string) error {
	_, err := db.Exec("DELETE FROM posts WHERE id=?", postid)
	return err
}

func EditPost(postid, title, content string) error {
	_, err := db.Exec("UPDATE posts SET title=?, content=? WHERE id=?;", title, content, postid)
	return err
}
