package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	// "gotrading/app/models"
	"gotrading/config"
)

var templates = template.Must(template.ParseFiles("app/views/google.html"))

// func init() {
// 	fmt.Printf("Templates loaded: %+v\n", templates)
// }
//templtesにはメモリなどが格納される
// Templates loaded: &{escapeErr:<nil> text:0xc000226480 Tree:0xc00023ec60 nameSpace:0xc000228300}

func viewChartHandler(w http.ResponseWriter, r *http.Request) {

	err := templates.ExecuteTemplate(w, "google.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func StartWebServer() error {
	http.HandleFunc("/chart/", viewChartHandler)
	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), nil)
	//dに8080を代入するという意味
}
