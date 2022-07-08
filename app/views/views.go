package views

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"app/pkg"

	// "github.com/astaxie/session"
	"github.com/joho/godotenv"
)

func TopPageHandler(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now()
	expiration = expiration.AddDate(1, 0, 0)
	cookie := http.Cookie{Name: "username", Value: "astaxie", Expires: expiration}
	fmt.Println("############", cookie)
	http.SetCookie(w, &cookie)

	db := pkg.OpenDB()
	defer db.Close()

	blogList, err := pkg.QueryBlogList(db)
	if err != nil {
		log.Fatal(err)
	}

	// body, _ := os.ReadFile("text/top_page.txt")
	// p := &models.Blog{Title: "top_page", Body: body}
	pkg.RenderTemplate(w, "top_page", blogList)
}

func BlogContentsPageHandler(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("username")
	fmt.Println("cookie ============>", cookie)

	for _, cookie := range r.Cookies() {
		fmt.Println("cookie ------------->", cookie.Name)
	}

	fmt.Println("###################", r.Cookies())

	db := pkg.OpenDB()
	defer db.Close()

	id := strings.TrimPrefix(r.URL.Path, "/blog/")

	blog, err := pkg.GetBlog(db, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(blog)

	pkg.RenderTemplate(w, "blog_contents", blog)
}

// var globalSessions *session.Manager

func PostPageHandler(w http.ResponseWriter, r *http.Request) {
	// globalSessions, _ = session.NewManager()

	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatal("error opening env file: ", envErr)
	}

	token, _ := pkg.GenerateJWTToken(os.Getenv("ADMIN_USER_NAME"), os.Getenv("ADMIN_USER_PASS"))
	fmt.Println("token +++++++++>", token, "<++++++++")

	pkg.RenderTemplate(w, "post_page", "POST")
}

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	pkg.RenderTemplate(w, "loging_page", "POST")
}
