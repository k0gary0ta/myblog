package main

import (
	"log"
	"net/http"

	"app/views"
	// "github.com/astaxie/session"
	// _ "github.com/astaxie/session/providers/memory"
)

// var sessions *session.Manager

func main() {
	// sessions, _ = session.NewManager("momery", "gosessionid", 3600)
	// go sessions.GC()

	// views
	http.HandleFunc("/", views.TopPageHandler)
	http.HandleFunc("/blog/", views.BlogContentsPageHandler)
	http.HandleFunc("/blog/form/", views.PostPageHandler)

	// post
	http.HandleFunc("/blog/post/", views.BlogPostsProcess)

	// static
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
