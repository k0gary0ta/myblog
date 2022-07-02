package views

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"app/pkg"
)

func BlogPostsProcess(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}

	for k, v := range r.Form {
		fmt.Println(k, v)
	}

	token := r.Form["sub_title"][0]
	parsed, err := pkg.ValidateJWTToken(token)

	if err != nil {
		fmt.Println(err)
		return
	}

	if parsed {
		db := pkg.OpenDB()
		defer db.Close()

		blogId := pkg.GenerateULID()

		tLayout := "2006-01-02 15:04:05"
		currentTime := time.Now().Format(tLayout)

		result, err := db.Exec(
			"INSERT INTO blog (ID, Title, Body, CreatedAt) VALUES (?, ?, ?, ?)",
			blogId, r.Form["title"][0], r.Form["body"][0], currentTime,
		)
		if err != nil {
			log.Fatal("Post blog error: ", err)
			return
		}
		fmt.Println("***** result *****", result)
	}
}
