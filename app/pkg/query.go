package pkg

import (
	"database/sql"
	"fmt"

	"app/models"
)

func QueryBlogList(db *sql.DB) ([]models.Blog, error) {
	var blogs []models.Blog

	rows, err := db.Query("SELECT * FROM blog ORDER BY CreatedAt DESC LIMIT 2")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var blg models.Blog
		if err := rows.Scan(&blg.ID, &blg.Title, &blg.Body, &blg.CreatedAt); err != nil {
			return blogs, err
		}
		blogs = append(blogs, blg)
	}
	if err = rows.Err(); err != nil {
		return blogs, err
	}
	fmt.Println(blogs)
	return blogs, nil
}

func GetBlog(db *sql.DB, id string) (models.Blog, error) {
	var blog models.Blog

	row := db.QueryRow("SELECT * FROM blog WHERE id = ?", id)
	if err := row.Scan(&blog.ID, &blog.Title, &blog.Body, &blog.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return blog, fmt.Errorf("getBlog %s: no such blog", id)
		}
		return blog, fmt.Errorf("getBlog %s: %v", id, err)
	}
	return blog, nil
}
