package pkg

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles(
	"public/templates/top_page.html",
	"public/templates/post_page.html",
	"public/templates/blog_contents.html",
))

func RenderTemplate(w http.ResponseWriter, tmpl string, i interface{}) {
	// switch i.(type) {
	// case string:
	// 	fmt.Println("The type of argument i is string.")
	// case *models.Blog:
	// 	fmt.Println("The type of argument i is *models.Blog.")
	// default:
	// 	http.Error(w, "Unable to load template", http.StatusInternalServerError)
	// }

	err := templates.ExecuteTemplate(w, tmpl+".html", i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
